// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"strike/go/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//Login login.htmlのGET処理を実装します
func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

//DoLogin index.htmlのPOST処理（/login）を実装します
func DoLogin(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("userName")
	password, _ := ctx.GetPostForm("password")

	userinfo := model.DoAuthentication(username,password)

	if userinfo.UserName != "" {
		session := sessions.Default(ctx)
		session.Set("UserId", userinfo.UserID)
		session.Set("UserName", userinfo.UserName)
		session.Set("OpenRoomId", userinfo.OpenRoomID)
		session.Save()
		ctx.Redirect(http.StatusSeeOther, "/chatList")
	} else {
		ctx.HTML(http.StatusOK, "error.html", gin.H{})
	}
}

//GetLogin /getLoginのGET処理を実装します
func GetLogin(ctx *gin.Context) {
    username := ctx.Query("username")
	password := ctx.Query("password")
	userinfo := model.DoAuthentication(username,password)
	
	err := bcrypt.CompareHashAndPassword([]byte(userinfo.Password), []byte(password))
	ctx.Writer.Header().Set("Content-Type", "application/json")
        ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
        ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
        ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if err == nil {
		ctx.JSON(http.StatusOK,userinfo)
	} else {
		ctx.JSON(http.StatusOK,gin.H{
		"err":   "ユーザ情報が取得できませんでした。",
	})
	}
}
