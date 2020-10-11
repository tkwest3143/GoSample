package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

//User usersテーブルエンティティ情報
type User struct {
	gorm.Model
	UserID      string    `gorm:"primary_key"` //ユーザID
	UserName    string    //ユーザ名
	Password    string    //パスワード
	MailAddress string    //メールアドレス
	OpenRoomID  string    //最後に開いていたRoomID
	RegistDate  time.Time //登録日
	LastLogin   time.Time //最終ログイン日
}
