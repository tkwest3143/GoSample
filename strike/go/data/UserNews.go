package data

import (
	"gorm.io/gorm"
)

//User usersテーブルエンティティ情報
type UserNews struct {
	gorm.Model
	NewsID    string `gorm:"primary_key"` //ニュースID
	UserID    string //ユーザ名
	TargetUrl string //対象のURL
}
