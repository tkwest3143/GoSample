package db

import (
	"fmt"
	"strike/go/data"

	"gorm.io/gorm"
)

//RoomInit Roomsテーブルの初期化を行います
func RoomInit(d *gorm.DB) {
	d.Migrator().CreateTable(&data.Room{})
}

//RoomInsert users情報を挿入します
func RoomInsert(insData data.Room) {
	d := GormConnect()
	d.Create(&insData)
}

//RoomSelect roomIDをもとにrooms情報を取得します。
func RoomSelect(roomID string) data.Room {
	d := GormConnect()
	selData := data.Room{}
	d.Find(&selData, "room_id=?", roomID)
	return selData
}

//GetMaxRoomID IDの最大値を取得します。
func GetMaxRoomID() string {
	d := GormConnect()
	var maxNo int64
	d.Table("rooms").Count(&maxNo)
	newUserID := "RI" + fmt.Sprintf("%08d", maxNo+1)
	return newUserID
}
