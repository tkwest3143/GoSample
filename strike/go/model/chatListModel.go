package model

import (
	"strike/go/data"
	"strike/go/db"
)

//GetChatList RoomIDをもとに画面表示用チャット情報を取得します。
func GetChatList(roomID string) []data.ChatList {
	d := db.GormConnect()

	//チャット情報の取得
	selData := []data.ChatList{}
	d.Raw("select "+
		"t1.chat_text,t2.user_id,t2.user_name as contributer,t1.bote_date,"+
		"t1.room_id from chats t1 "+
		"inner join users t2 on t1.contributer = t2.user_id "+
		"where t1.room_id = ? order by t1.bote_date", roomID).Scan(&selData)
	return selData
}
