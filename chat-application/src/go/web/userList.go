// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"github/GoSumple/chat-application/src/go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUserList /getUserListのGET処理を実装します
func GetUserList(ctx *gin.Context) {

    username := ctx.Query("username")
	password := ctx.Query("password")
	userInfo := model.DoAuthentication(username,password)
ctx.Writer.Header().Set("Content-Type", "application/json")
        ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
        ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
        ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.JSON(http.StatusOK, userInfo)
}

