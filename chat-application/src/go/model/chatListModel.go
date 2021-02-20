package model

import (
	"github/GoSumple/chat-application/src/go/data"
	"github/GoSumple/chat-application/src/go/db"
)

//GetChatList RoomIDをもとに画面表示用チャット情報を取得します。
func GetChatList(roomID string) []data.ChatList {
	d := db.GormConnect()

	//チャット情報の取得
	selData := []data.ChatList{}
	d.Raw("select "+
		"t1.chat_text,t2.user_id,t2.user_name as contributer,t1.bote_date,to_char(t1.bote_date,'hh24:mi') as bote_date_disp,"+
		"t1.room_id from chats t1 "+
		"inner join users t2 on t1.contributer = t2.user_id "+
		"where t1.room_id = ? order by t1.bote_date", roomID).Scan(&selData)
	defer d.Close()
	return selData
}
