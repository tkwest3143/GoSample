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

	//login.htmlのGET、POST処理
	r.POST("/doLogin", DoLogin)

	//regist.htmlのGET、POST処理
	r.POST("/doRegist", DoRegist)

	//loginError.htmlのGET、POST処理
	r.POST("/returnLogin", ReturnLogin)

	//chatList.htmlのGET、POST処理
	r.GET("/chatList", ChatList)
	r.POST("/chatPost", ChatPost)
	r.POST("/doRoomCreate", DoRoomCreate)

	/**
	* Json形式のデータのみを連携するRest処理
	*/
	r.POST("/getLogin",GetLogin)
	r.GET("/getLogin",GetLogin)
	r.GET("/getUserList",GetUserList)

}
