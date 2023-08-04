package handler

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
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
	var marchantInput model.MarchantModel
	err := ctx.ShouldBindJSON(&marchantInput)
	if err != nil {
        fmt.Println("error on", err)
        ctx.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Invalid data JSON",
        })
        return
    }

	err = mh.marchantUsecase.Create(&marchantInput)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Terjadi Kesalahan Saat menyimpan Data Marchant",
		})
		return
	}
}

func NewMarchantHandler(serve *gin.Engine, marchantUsecase usecase.MarchantUsecase) MarchantHandler {
	marchantHandler := &marchantHandler{
		serve: serve,
		marchantUsecase: marchantUsecase,
	}

	serve.POST("/marchant", marchantHandler.Create)

	return marchantHandler
}