package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//Router 各画面のGET、POST処理を一覧で設定しています
func Router(r *gin.Engine) {

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//loginのPOST処理
	r.POST("/doLogin", DoLogin)

	//registerのPOST処理
	r.POST("/doRegister", DoRegister)

	//chatList.htmlのGET、POST処理
	r.GET("/chatList", ChatList)
	r.POST("/chatPost", ChatPost)
	r.POST("/doRoomCreate", DoRoomCreate)
	r.POST("/upload", Upload)

	r.GET("/getUserList", GetUserList)

	r.GET("/getNewsList", GetNewsList)

}
