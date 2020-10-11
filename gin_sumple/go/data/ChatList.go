package data

import (
	"time"
)

//ChatList Chatsの画面表示用情報
type ChatList struct {
	ChatText    string    //投稿文
	UserID      string    //ユーザID
	Contributer string    //投稿者
	BoteDate    time.Time //投稿日
	RoomID      string    //部屋ID
	RoomName    string    //部屋名
}
