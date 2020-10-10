/*
 */
package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github/GoSumple/gin_sumple/go/common"
	"github/GoSumple/gin_sumple/go/db"
	"net/http"
)

//Login login.htmlのGET処理を実装します
func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

//DoLogin index.htmlのPOST処理（/login）を実装します
func DoLogin(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("userName")
	password, _ := ctx.GetPostForm("password")

	userinfo := db.UserSelect(username, common.DoCrypt(password))

	if userinfo.UserID != "" {
		session := sessions.Default(ctx)
		session.Set("UserId", userinfo.UserID)
		session.Set("UserName", userinfo.UserName)
		session.Save()

		ctx.Redirect(http.StatusSeeOther, "/chatList")
	} else {
		ctx.HTML(http.StatusOK, "error.html", gin.H{})
	}

}
