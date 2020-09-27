package data

import (
	"time"
)

//Chat Chatsテーブル情報
type ChatList struct {
	ChatText    string
	UserId      string
	Contributer string    //投稿者
	BoteDate    time.Time //投稿日
	RoomId      string    //部屋ID
}
