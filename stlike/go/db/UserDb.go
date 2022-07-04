package db

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"strike/go/data"
)

//UserDBInit Usersテーブルの初期化を行います
func UserDBInit(d *gorm.DB) {
	if !d.HasTable(&data.User{}) {
		d.CreateTable(&data.User{})
	}
}
//UserInsert users情報を挿入します
func UserInsert(insData data.User) {
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//UserSelect userIDをもとにusers情報を取得します。
func UserSelect(userID string) data.User {
	d := GormConnect()
	selData := data.User{}
	d.First(&selData, "user_id=?", userID)
	defer d.Close()
	return selData
}

//UserSelectByUserName userIdをもとにusers情報を取得します。
func UserSelectByUserName(userName string) data.User {
	d := GormConnect()
	selData := data.User{}
	d.First(&selData, "user_name=?", userName)
	defer d.Close()
	return selData
}

//GetMaxUsersID IDの最大値を取得します。
func GetMaxUsersID() string {
	d := GormConnect()
	maxNo := 0
	d.Table("users").Count(&maxNo)
	newUserID := "UI" + fmt.Sprintf("%08d", maxNo+1)
	defer d.Close()
	return newUserID
}
