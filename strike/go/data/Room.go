package data

import (
	"gorm.io/gorm"
)

//Room Roomsテーブル情報
type Room struct {
	gorm.Model
	RoomName    string //ルーム名
	RoomID      string //ルームID
	Discription string //概要
}
