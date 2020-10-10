package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

//Room Roomsテーブル情報
type Room struct {
	gorm.Model
	RoomName      string
	Administrator string    //管理者
	CreateDate    time.Time //作成日
	RoomID        string    //部屋ID
}
