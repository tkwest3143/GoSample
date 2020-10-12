package data

import (
	"github.com/jinzhu/gorm"
)

//RoomUserRelation RoomUserRelationsテーブル情報
type RoomUserRelation struct {
	gorm.Model
	RoomID string //部屋ID
	UserID string //参加者
}
