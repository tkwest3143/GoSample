/*
 */
package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github/GoSumple/gin_sumple/go/data"
	"github/GoSumple/gin_sumple/go/db"
	"net/http"
	"time"
)

//ChatList chatList.htmlのGET処理を実装します
func ChatList(ctx *gin.Context) {

	session := sessions.Default(ctx)
	userid := session.Get("UserId").(string)
	chatList := db.ChatSelect(userid)

	ctx.HTML(http.StatusOK, "chatList.html", gin.H{
		"chatList": chatList,
		"userId":   userid,
	})
}

//ChatPost chatList.htmlのPOST処理（/chatPost）を実装します
func ChatPost(ctx *gin.Context) {
	chattext, _ := ctx.GetPostForm("chatText")
	chat := data.Chat{}
	chat.ChatText = chattext
	chat.Contributer = "111"
	chat.BoteDate = time.Now()
	chat.RoomId = "111"
	db.ChatInsert(chat)

	ctx.Redirect(http.StatusSeeOther, "/chatList")
}
