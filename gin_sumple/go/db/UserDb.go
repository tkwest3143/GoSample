package db

import (
	"github/GoSumple/gin_sumple/go/data"
	"strconv"
)

//UserDBInit usersテーブルの初期化を行います
func UserDBInit() {
	d := GormConnect()
	if !d.HasTable(&data.User{}) {
		d.CreateTable(&data.User{})
	}

	defer d.Close()

}

//UserInsert users情報を挿入します
func UserInsert(insData data.User) {
	UserDBInit()
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//UserSelect userId,passwordをもとにusers情報を取得します。
func UserSelect(userName string, password string) data.User {
	UserDBInit()
	d := GormConnect()
	selData := data.User{}
	d.First(&selData, "user_name=? and password = ?", userName, password)
	defer d.Close()
	return selData
}

//UserSelectByUserName userIdをもとにusers情報を取得します。
func UserSelectByUserName(userName string) data.User {
	UserDBInit()
	d := GormConnect()
	selData := data.User{}
	d.First(&selData, "user_name=?", userName)
	defer d.Close()
	return selData
}

//GetMaxUsersID IDの最大値を取得します。
func GetMaxUsersID() string {
	UserDBInit()
	d := GormConnect()
	maxNo := 0
	d.Table("users").Count(&maxNo)
	newUserID := "UI" + strconv.Itoa(maxNo+1)
	defer d.Close()
	return newUserID
}
