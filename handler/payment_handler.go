package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syahyudi09/BankingAPI/auth"
	"github.com/syahyudi09/BankingAPI/model"
	"github.com/syahyudi09/BankingAPI/usecase"
)

type PaymentHanlder interface {
}

type paymentHandlerImpl struct {
	serve *gin.Engine
	paymentUsecase usecase.PaymentUsecase
}

func(ph *paymentHandlerImpl) Payment(ctx *gin.Context){
	paymentRequst := &model.PaymentAccount{}
	err := ctx.ShouldBindJSON(&paymentRequst)
	if err != nil {
        fmt.Println("error on", err)
        ctx.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Invalid data JSON",
        })
        return
    }

	err = ph.paymentUsecase.Payment(*paymentRequst)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Terjadi Kesalahan Saat menyimpan Data Customer",
		})
		return
	}
}

func NewPaymantHandler(serve *gin.Engine,paymentUsecase usecase.PaymentUsecase) PaymentHanlder {
	 paymentHandler := paymentHandlerImpl{
		serve: serve,
		paymentUsecase: paymentUsecase,
	}

	authService := auth.NewServiceJWT()
	middleware := auth.NewMiddleware(authService)

	authenticated := serve.Group("/")
	authenticated.Use(middleware.AuthMiddleware())

	authenticated.POST("/payment", paymentHandler.Payment)
	return paymentHandler
}