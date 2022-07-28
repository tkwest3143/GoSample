package web

import (
	"net/http"
	"strike/go/model"

	"github.com/gin-gonic/gin"
)

func RegisterWotkTimeOfDay(ctx *gin.Context) {

	username := ctx.Query("username")
	password := ctx.Query("password")
	userInfo := model.DoAuthentication(username, password)
	ctx.JSON(http.StatusOK, userInfo)
}

func GetWotkTimeByUser(ctx *gin.Context) {

	username := ctx.Query("username")
	password := ctx.Query("password")
	userInfo := model.DoAuthentication(username, password)
	ctx.JSON(http.StatusOK, userInfo)
}

func RegisterWorkTimeSetting(ctx *gin.Context) {

	username := ctx.Query("username")
	password := ctx.Query("password")
	userInfo := model.DoAuthentication(username, password)
	ctx.JSON(http.StatusOK, userInfo)
}
