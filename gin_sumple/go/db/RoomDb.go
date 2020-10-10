package db

import (
	"github/GoSumple/gin_sumple/go/data"
)

//RoomDBInit Roomsテーブルの初期化を行います
func RoomDBInit() {
	d := GormConnect()
	if !d.HasTable(&data.Room{}) {
		d.CreateTable(&data.Room{})
	}

	defer d.Close()

}

//RoomInsert users情報を挿入します
func RoomInsert(insData data.Room) {
	RoomDBInit()
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}
