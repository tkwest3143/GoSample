package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

//Chat Chatsテーブルエンティティ情報
type Chat struct {
	gorm.Model
	ChatText    string
	Contributer string    //投稿者
	BoteDate    time.Time //投稿日
	RoomID      string    //部屋ID
}
