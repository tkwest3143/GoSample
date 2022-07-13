package data

import (
	"gorm.io/gorm"
)

//UserRelation UserRelationsテーブル情報
type UserRelation struct {
	gorm.Model
	UserID1    string //部屋ID
	UserID2    string //参加者
	RelationCd rune   //関係コード
}
