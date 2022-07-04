package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

//UserListData UserListData
type UserListData struct {
	gorm.Model
	UserID    string //ユーザID
	UserName    string //ユーザ名
	RelationCd rune   //関係コード
	LastLogin      time.Time //最終ログイン日
	UserNameSearch string    //ユーザ名（検索用）
	LoginFlg       string    //ログインフラグ
}
