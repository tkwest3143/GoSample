package db

import (
	"github.com/jinzhu/gorm"
	"github/GoSumple/gin_sumple/go/data"
)

//ChatDBInit chatsテーブルの初期化を行います
func ChatDBInit(d *gorm.DB) {
	if !d.HasTable(&data.Chat{}) {
		d.CreateTable(&data.Chat{})
	}
}

//ChatInsert users情報を挿入します
func ChatInsert(insData data.Chat) {
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//ChatSelect userId,passwordをもとにusers情報を取得します。
func ChatSelect(roomID string) []data.Chat {
	d := GormConnect()
	selData := []data.Chat{}
	d.Find(&selData, "room_id = ?", roomID)
	defer d.Close()
	return selData
}
