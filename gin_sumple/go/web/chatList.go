// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github/GoSumple/gin_sumple/go/data"
	"github/GoSumple/gin_sumple/go/db"
	"github/GoSumple/gin_sumple/go/model"
	"net/http"
	"strings"
	"time"
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
	friendIDList := db.GetFriendList(userID)
	friendList := []data.User{}
	for _, friend := range friendIDList {
		friendList = append(friendList, db.UserSelect(friend))
	}
	ctx.HTML(http.StatusOK, "chatList.html", gin.H{
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

	ctx.Redirect(http.StatusSeeOther, "/chatList")
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
	ctx.Redirect(http.StatusSeeOther, "/chatList")

}
