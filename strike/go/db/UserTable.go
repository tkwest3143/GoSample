package db

import (
	"fmt"
	"strike/go/data"

	"gorm.io/gorm"
)

//UserInit Usersテーブルの初期化を行います
func UserInit(d *gorm.DB) {
	d.AutoMigrate(&data.User{})
}

//UserInsert users情報を挿入します
func UserInsert(insData data.User) {
	d := GormConnect()
	d.Create(&insData)
}

//UserSelect userIDをもとにusers情報を取得します。
func UserSelect(userID string) data.User {
	d := GormConnect()
	selData := data.User{}
	d.First(&selData, "user_id=?", userID)
	return selData
}

//UserSelectByUserName userIdをもとにusers情報を取得します。
func UserSelectByUserName(userName string) data.User {
	d := GormConnect()
	selData := data.User{}
	d.First(&selData, "user_name=?", userName)
	return selData
}

//GetMaxUsersID IDの最大値を取得します。
func GetMaxUsersID() string {
	d := GormConnect()
	var maxNo int64
	d.Table("users").Count(&maxNo)
	newUserID := "UI" + fmt.Sprintf("%08d", maxNo+1)
	return newUserID
}
