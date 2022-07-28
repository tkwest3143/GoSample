package data

import (
	"time"

	"gorm.io/gorm"
)

//User usersテーブルエンティティ情報
type User struct {
	gorm.Model
	ID               string    `gorm:"primary_key"` //ユーザID
	UserName         string    //ユーザ名
	Password         string    //パスワード
	MailAddress      string    //メールアドレス
	OpenRoomID       string    //最後に開いていたRoomID
	RegistDate       time.Time //登録日
	LastLogin        time.Time //最終ログイン日
	UserNameSearch   string    //ユーザ名（検索用）
	LoginFlg         string    //ログインフラグ
	WorkTime         []WorkTime
	WorkTimeSettings WorkTimeSetting
}
