// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"net/http"
	"strike/go/model"

	"github.com/gin-gonic/gin"
)

//GetUserList /getUserListのGET処理を実装します
func GetUserList(ctx *gin.Context) {

	username := ctx.Query("username")
	password := ctx.Query("password")
	userInfo := model.DoAuthentication(username, password)
	ctx.JSON(http.StatusOK, userInfo)
}
