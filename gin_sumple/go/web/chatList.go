// Package web チャット画面からのGET、POST処理を実装します
package web

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github/GoSumple/gin_sumple/go/data"
	"github/GoSumple/gin_sumple/go/db"
	"github/GoSumple/gin_sumple/go/model"
	"net/http"
	"time"
)

//ChatList chatList.htmlのGET処理を実装します
func ChatList(ctx *gin.Context) {

	session := sessions.Default(ctx)
	roomid := session.Get("OpenRoomId").(string)
	chatList := model.GetChatList(roomid)

	username := session.Get("UserName").(string)
	userId := session.Get("UserId").(string)
	fmt.Println(userId)
	ctx.HTML(http.StatusOK, "chatList.html", gin.H{
		"chatList": chatList,
		"userName": username,
		"userId":   userId,
	})
}

//ChatPost chatList.htmlのPOST処理（/chatPost）を実装します
func ChatPost(ctx *gin.Context) {
	chattext, _ := ctx.GetPostForm("chatText")
	session := sessions.Default(ctx)
	userid := session.Get("UserId").(string)
	roomid := session.Get("OpenRoomId").(string)
	chat := data.Chat{}
	chat.ChatText = chattext
	chat.Contributer = userid
	chat.BoteDate = time.Now()
	chat.RoomID = roomid
	db.ChatInsert(chat)

	ctx.Redirect(http.StatusSeeOther, "/chatList")
}
