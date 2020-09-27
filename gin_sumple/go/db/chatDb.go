package db

import (
	"github/GoSumple/gin_sumple/go/data"
)

//ChatDBInit chatsテーブルの初期化を行います
func ChatDBInit() {
	d := GormConnect()
	if !d.HasTable(&data.Chat{}) {
		d.CreateTable(&data.Chat{})
	}

	defer d.Close()

}

//ChatInsert users情報を挿入します
func ChatInsert(insData data.Chat) {
	ChatDBInit()
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//ChatSelect userId,passwordをもとにusers情報を取得します。
func ChatSelect(roomId string) []data.Chat {
	ChatDBInit()
	d := GormConnect()
	selData := []data.Chat{}
	d.Find(&selData, "room_id = ?", roomId)
	defer d.Close()
	return selData
}
