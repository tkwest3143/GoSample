/*
 */
package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github/GoSumple/gin_sumple/go/db"
	"net/http"
)

//Index index.htmlのGET処理を実装します
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

//Login index.htmlのPOST処理（/login）を実装します
func Login(ctx *gin.Context) {
	userid, _ := ctx.GetPostForm("userId")
	password, _ := ctx.GetPostForm("password")

	userinfo := db.UserSelect(userid, password)

	if userinfo.UserId != "" {
		session := sessions.Default(ctx)
		session.Set("UserId", userinfo.UserId)
		session.Save()

		ctx.Redirect(http.StatusSeeOther, "/chatList")
	} else {
		ctx.HTML(http.StatusOK, "loginError.html", gin.H{})
	}

}
