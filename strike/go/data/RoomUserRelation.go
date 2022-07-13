package data

import (
	"gorm.io/gorm"
)

//RoomUserRelation RoomUserRelationsテーブル情報
type RoomUserRelation struct {
	gorm.Model
	RoomID      string //部屋ID
	UserID      string //参加者
	AuthorityCd rune   //権限コード
}
