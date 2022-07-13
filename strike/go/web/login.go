// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"fmt"
	"net/http"

	"strike/go/data"
	"strike/go/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//DoLogin POST処理（/doLogin）を実装します
func DoLogin(ctx *gin.Context) {
	xml_data, err := model.ReadXmlFromHttp("https://news.yahoo.co.jp/rss/topics/it.xml")
	fmt.Println(xml_data, err)
	var loginData data.DoLoginData
	if ctx.ShouldBind(&loginData) == nil {
		userinfo := model.DoAuthentication(loginData.Username, loginData.Password)

		if userinfo.UserName != "" {
			session := sessions.Default(ctx)
			session.Set("UserId", userinfo.UserID)
			session.Set("UserName", userinfo.UserName)
			session.Set("OpenRoomId", userinfo.OpenRoomID)
			session.Save()
			ctx.JSON(http.StatusOK, gin.H{
				"user_id":  userinfo.UserID,
				"username": userinfo.UserName,
				"email":    userinfo.MailAddress,
				"password": userinfo.Password,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "login error",
			})
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "not expected request parameter",
		})
	}

}
