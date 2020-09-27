package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

//Chat Chatsテーブル情報
type Chat struct {
	gorm.Model
	ChatText    string
	Contributer string    //投稿者
	BoteDate    time.Time //投稿日
	RoomId      string    //部屋ID
}
