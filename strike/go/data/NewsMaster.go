package data

import (
	"gorm.io/gorm"
)

//User usersテーブルエンティティ情報
type NewsMaster struct {
	gorm.Model
	ID       string //タイトル
	Title    string //タイトル
	FileName string //ファイル名
}
