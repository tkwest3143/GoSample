package db

import (
	"github/GoSumple/gin_sumple/go/data"
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
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//UserSelect userId,passwordをもとにusers情報を取得します。
func UserSelect(userId string, password string) data.User {
	d := GormConnect()
	selData := data.User{}
	d.First(&selData, "user_id=? and password = ?", userId, password)
	defer d.Close()
	return selData
}
