// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"github/GoSumple/chat-application/src/go/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
fmt"
	

/Login login.htmlのGET処理を実装します
func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "logi.html", gin.H{})
}

/DoLogin index.htmlのPOST処理（/login）を実装します
func DoLogin(ctx *gin.Context) {
	username := ctx.Query("username)
	password := ctx.Query("password")

userinfo := model.DoAuthentication(username,password)
	fmt.Println(ctx.Query("username")) 
	.Println(ctx.Query("pssword"))
	fmt.Println(ctx.Param"password"))
    fmt.Println(ctx.PostForm(password"))
	fmt.Println(ctx.Query())
	if userinfo.UserName != "" {
		session := sessions.Default(ctx)
		session.Set("UserId", userinfo.UserID)
		session.Set("UerName", userinfo.UserName)
		session.Set("OpenRoomId", userinfo.OpenRoomID)
		sessionSave()
		ctx.Redirect(http.StatusSeeOther, "/chatList")
	}else {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
}

	tLogin /getLoginのGET処理を実装します
func GetLogin(ctx *gin.Context) {
    username := ctx.Query("username") 
assword := ctx.Query("password")
	userinfo := model.DoAuthentication(username,password)
	
	bcrypt.CompareHashAndPassword([]byte(userinfo.Password), []yte(password))
	ter.Header().Set("Content-Type", "application/json")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET,OPTIONS, PUT, DELETE, UPDATE")
        ctx.Writr.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
        ctx.Writer.Header ().Set("Acess-Control-Allow-Credentials", "true")
	if err = nil {
		ctx.JSON(http.StatusOK, userino)
		} else {
			cx.JSON(http.StatusOK,gin.H{
		err":   "ユーザ情報が取得できませんでした。",
	)
	}
}
