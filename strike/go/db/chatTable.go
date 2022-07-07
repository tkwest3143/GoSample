package db

import (
	"strike/go/data"

	"gorm.io/gorm"
)

//ChatInit chatsテーブルの初期化を行います
func ChatInit(d *gorm.DB) {
	d.AutoMigrate(&data.Chat{})

}

//ChatInsert users情報を挿入します
func ChatInsert(insData data.Chat) {
	d := GormConnect()
	d.Create(&insData)
}

//ChatSelect userId,passwordをもとにusers情報を取得します。
func ChatSelect(roomID string) []data.Chat {
	d := GormConnect()
	selData := []data.Chat{}
	d.Order("bote_date desc").Find(&selData, "room_id = ?", roomID)
	return selData
}
