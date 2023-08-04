package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/syahyudi09/BankingAPI/auth"
	"github.com/syahyudi09/BankingAPI/model"
	"github.com/syahyudi09/BankingAPI/usecase"
)

type CustomerHandler interface {
}

type customerHandlerImpl struct {
	server *gin.Engine
	customerUsecase usecase.CustomerUsecase
	accountUsecae usecase.AccountUsecase
}

func (customerHandler *customerHandlerImpl) CreateCustomer(ctx *gin.Context) {
    customer := &model.RegistrasiCustomerModel{}
    err := ctx.ShouldBindJSON(&customer)
    if err != nil {
        fmt.Println("error on", err)
        ctx.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Invalid data JSON",
        })
        return
    }

    validate := validator.New()
    if err := validate.Struct(customer); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": err.Error(),
        })
        return
    }

    // Register the customer
    err = customerHandler.customerUsecase.RegisterCustomer(customer)
    if err != nil {
        log.Println("err", err) // Gunakan log.Println atau log.Printf untuk mencatat kesalahan, bukan log.Fatal.
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Terjadi Kesalahan Saat menyimpan Data Customer",
        })
        return
    }

    account := &model.AccountModel{
		CustomerID: customer.ID,
	}
    err = customerHandler.accountUsecae.CreateAccount(account)
    if err != nil {
        log.Println("err", err) // Gunakan log.Println atau log.Printf untuk mencatat kesalahan, bukan log.Fatal.
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Terjadi Kesalahan Saat menyimpan Data Customer",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Registrasi Berhasil",
    })
}

func (customerHandler *customerHandlerImpl) LoginCustomer(ctx *gin.Context) {
	var inputLogin model.CustomerLogin
	if err := ctx.ShouldBindJSON(&inputLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid data JSON",
		})
		return
	}

	token, err := customerHandler.customerUsecase.LoginCustomer(inputLogin)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid email or password",
		})
		return
	}

	// handler
	session := sessions.Default(ctx)
	session.Set("email", inputLogin.Email)
	session.Set("password", inputLogin.Password)
	
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to save session",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successful",
		"token": token,
	})
}

func (customerHandler *customerHandlerImpl) GetAllCustomer(ctx *gin.Context){
	session := sessions.Default(ctx)

	// Cek apakah ada sesi "email" yang ada
	if session.Get("email") == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "You need to login first",
		})
		return
	}

	customers, err := customerHandler.customerUsecase.GetAllCustomer()
	if err != nil{
		log.Fatal("err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Terjadi Kesalahan Saat menyimpan Data Customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": customers,
	})	
}	

func(customerHandler *customerHandlerImpl) LogoutHandler(ctx *gin.Context) {
	// Dapatkan instance toko sesi default dari objek context
	session := sessions.Default(ctx)

	// Hapus data "email" dari sesi
	session.Delete("email")
	session.Delete("password")

	// Simpan sesi setelah menghapus data "email"
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to clear session",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout success",
	})

}

func NewCustomerHandler(serve *gin.Engine, customerUsecase usecase.CustomerUsecase, accountUsecase usecase.AccountUsecase) CustomerHandler {
	customerHandler := &customerHandlerImpl{
		server: serve,
		customerUsecase: customerUsecase,
		accountUsecae: accountUsecase,
	}

	authService := auth.NewServiceJWT()
	middleware := auth.NewMiddleware(authService)

	authenticated := serve.Group("/")
	authenticated.Use(middleware.AuthMiddleware())

	serve.POST("/register",  customerHandler.CreateCustomer)
	serve.POST("/login",customerHandler.LoginCustomer)
	serve.POST("/logout", customerHandler.LogoutHandler)
	authenticated.GET("/GetAllCustomer" ,customerHandler.GetAllCustomer)
	return customerHandler
}
