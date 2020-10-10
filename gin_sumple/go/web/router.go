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

	//index.htmlのGET、POST処理
	r.GET("/", Index)

	//login.html
	r.GET("/login", Login)
	r.POST("/doLogin", DoLogin)

	//regist.html
	r.GET("/regist", Regist)
	r.POST("/doRegist", DoRegist)

	//loginError.html
	r.POST("/returnLogin", ReturnLogin)

	//chatList.html
	r.GET("/chatList", ChatList)
	r.POST("/chatPost", ChatPost)
}
