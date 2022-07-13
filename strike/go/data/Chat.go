package data

import (
	"time"

	"gorm.io/gorm"
)

//Chat Chatsテーブルエンティティ情報
type Chat struct {
	gorm.Model
	ChatText    string
	Contributer string    //投稿者
	BoteDate    time.Time //投稿日
	RoomID      string    //部屋ID
}
