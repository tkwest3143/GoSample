// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"net/http"
	"strike/go/common"
	"strike/go/data"
	"strike/go/db"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//DoRegister POST処理（/doRegister）を実装します
func DoRegister(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("userName")
	password, _ := ctx.GetPostForm("password")
	marilAddress, _ := ctx.GetPostForm("marilAddress")

	registuser := db.UserSelectByUserName(username)
	if registuser.UserID == "" {
		userid := db.GetMaxUsersID()
		userinfo := data.User{}
		userinfo.UserID = userid
		userinfo.UserName = username
		userinfo.Password = common.DoCrypt(password)
		userinfo.MailAddress = marilAddress
		userinfo.OpenRoomID = userid
		userinfo.RegistDate = time.Now()
		userinfo.LastLogin = time.Now()

		db.UserInsert(userinfo)

		roominfo := data.Room{}
		roominfo.RoomName = username
		roominfo.RoomID = userinfo.OpenRoomID
		db.RoomInsert(roominfo)
		roomUserRelation := data.RoomUserRelation{}
		roomUserRelation.RoomID = db.GetMaxRoomID()
		roomUserRelation.UserID = userid
		roomUserRelation.AuthorityCd = '1'
		session := sessions.Default(ctx)
		session.Set("UserId", userinfo.UserID)
		session.Set("UserName", userinfo.UserName)
		session.Set("OpenRoomId", userinfo.OpenRoomID)
		session.Save()

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "login error",
		})
	}

}
