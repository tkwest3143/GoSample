package data

import (
	"github.com/jinzhu/gorm"
)

//Room Roomsテーブル情報
type Room struct {
	gorm.Model
	RoomName    string //ルーム名
	RoomID      string //ルームID
	Discription string //概要
}
