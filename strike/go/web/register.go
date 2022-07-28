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
	var registerData data.DoRegisterData
	if ctx.ShouldBind(&registerData) == nil {
		registuser := db.UserSelectByUserName(registerData.Username)
		if registuser.ID == "" {
			userid := db.GetMaxUsersID()
			userinfo := data.User{}
			userinfo.ID = userid
			userinfo.UserName = registerData.Username
			userinfo.Password = common.DoCrypt(registerData.Password)
			userinfo.MailAddress = registerData.MarilAddress
			userinfo.OpenRoomID = userid
			userinfo.RegistDate = time.Now()
			userinfo.LastLogin = time.Now()

			db.UserInsert(userinfo)

			roominfo := data.Room{}
			roominfo.RoomName = registerData.Username
			roominfo.RoomID = userinfo.OpenRoomID
			db.RoomInsert(roominfo)
			roomUserRelation := data.RoomUserRelation{}
			roomUserRelation.RoomID = db.GetMaxRoomID()
			roomUserRelation.UserID = userid
			roomUserRelation.AuthorityCd = '1'
			session := sessions.Default(ctx)
			session.Set("UserId", userinfo.ID)
			session.Set("UserName", userinfo.UserName)
			session.Set("OpenRoomId", userinfo.OpenRoomID)
			session.Save()

			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})

		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "すでに登録されています",
			})
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "予期せぬエラー",
		})
	}

}
