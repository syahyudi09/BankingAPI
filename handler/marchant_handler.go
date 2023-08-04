package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syahyudi09/BankingAPI/auth"
	"github.com/syahyudi09/BankingAPI/model"
	"github.com/syahyudi09/BankingAPI/usecase"
)

type MarchantHandler interface {
}

type marchantHandler struct {
	serve *gin.Engine
	marchantUsecase usecase.MarchantUsecase
}

func (mh *marchantHandler) Create(ctx *gin.Context) {
	marchant := &model.MarchantModel{}
	err := ctx.ShouldBindJSON(&marchant)
	if err != nil {
        fmt.Println("error on", err)
        ctx.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Invalid data JSON",
        })
        return
    }

	err = mh.marchantUsecase.CreateMarchant(marchant)
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Terjadi Kesalahan Saat menyimpan Data Customer",
        })
        return
    }
	ctx.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Marchant Berhasil dibuat",
    })
}

func NewMarchantHandler(serve *gin.Engine, marchantUsecase usecase.MarchantUsecase) MarchantHandler {
    marchantHandler := &marchantHandler{
        serve: serve,
        marchantUsecase: marchantUsecase,
    }

	authService := auth.NewServiceJWT()
	middleware := auth.NewMiddleware(authService)

	authenticated := serve.Group("/")
	authenticated.Use(middleware.AuthMiddleware())


    authenticated.POST("/marchant", marchantHandler.Create)

    return marchantHandler
}
