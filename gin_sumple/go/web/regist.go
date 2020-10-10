/*
 */
package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github/GoSumple/gin_sumple/go/common"
	"github/GoSumple/gin_sumple/go/data"
	"github/GoSumple/gin_sumple/go/db"
	"net/http"
	"time"
)

//Regist regist.htmlのGET処理を実装します
func Regist(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "regist.html", gin.H{})
}

//DoRegist regist.htmlのPOST処理（/regist）を実装します
func DoRegist(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("userName")
	password, _ := ctx.GetPostForm("password")

	registuser := db.UserSelectByUserName(username)
	if registuser.UserID == "" {
		userid := db.GetMaxUsersID()
		userinfo := data.User{}
		userinfo.UserID = userid
		userinfo.UserName = username
		userinfo.Password = common.DoCrypt(password)
		userinfo.OpenRoomID = userid
		userinfo.RegistDate = time.Now()
		userinfo.LastLogin = time.Now()

		db.UserInsert(userinfo)

		roominfo := data.Room{}
		roominfo.RoomName = username
		roominfo.RoomID = userinfo.OpenRoomID
		roominfo.Administrator = userid
		roominfo.CreateDate = time.Now()
		db.RoomInsert(roominfo)
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
