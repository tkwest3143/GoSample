// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"net/http"
	"strike/go/data"
	"strike/go/db"
	"strike/go/model"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//ChatList chatList.htmlのGET処理を実装します
func ChatList(ctx *gin.Context) {

	session := sessions.Default(ctx)
	roomID := session.Get("OpenRoomId").(string)
	chatList := model.GetChatList(roomID)

	username := session.Get("UserName").(string)
	userID := session.Get("UserId").(string)

	//ルームIDよりルーム情報を取得
	roomInfo := db.RoomSelect(roomID)
	friendList := model.GetUserFriendList(userID)
	ctx.JSON(http.StatusOK, gin.H{
		"chatList":   chatList,
		"userName":   username,
		"userId":     userID,
		"roomName":   roomInfo.RoomName,
		"friendList": friendList,
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

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

//DoRoomCreate ルームを作成する処理を実装します。
func DoRoomCreate(ctx *gin.Context) {
	//ルーム作成ボタンを押下した際にルームを作成する処理
	roomID := db.GetMaxRoomID()
	room := data.Room{}
	room.RoomName, _ = ctx.GetPostForm("roomName")
	room.Discription, _ = ctx.GetPostForm("discription")
	room.RoomID = roomID
	session := sessions.Default(ctx)
	Administrator := session.Get("UserId").(string)
	roomUserRelation := data.RoomUserRelation{}
	roomUserRelation.UserID = Administrator
	roomUserRelation.RoomID = roomID
	roomUserRelation.AuthorityCd = '1'
	db.RoomInsert(room)
	db.RoomUserRelationInsert(roomUserRelation)
	//ルームに所属するユーザの登録
	selectFriendList, _ := ctx.GetPostForm("selectFriendList")
	//選択されたフレンドをカンマで区切り配列に格納
	sliceFriendList := strings.Split(selectFriendList, ",")

	for _, friend := range sliceFriendList {
		friendRoom := data.RoomUserRelation{}
		friendRoom.UserID = friend
		friendRoom.RoomID = roomID
		friendRoom.AuthorityCd = '0'
		db.RoomUserRelationInsert(friendRoom)
	}
	//ルームを開くようセッションに登録
	session.Set("OpenRoomId", roomID)
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

}
