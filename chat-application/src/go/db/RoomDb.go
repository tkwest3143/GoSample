package db

import (
	"fmt"
	"github/GoSumple/chat-application/src/go/data"

	"github.com/jinzhu/gorm"
)

//RoomDBInit Roomsテーブルの初期化を行います
func RoomDBInit(d *gorm.DB) {
	if !d.HasTable(&data.Room{}) {
		d.CreateTable(&data.Room{})
	}
}
//RoomInsert users情報を挿入します
func RoomInsert(insData data.Room) {
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//RoomSelect roomIDをもとにrooms情報を取得します。
func RoomSelect(roomID string) data.Room {
	d := GormConnect()
	selData := data.Room{}
	d.Find(&selData, "room_id=?", roomID)
	defer d.Close()
	return selData
}

//GetMaxRoomID IDの最大値を取得します。
func GetMaxRoomID() string {
	d := GormConnect()
	maxNo := 0
	d.Table("rooms").Count(&maxNo)
	newUserID := "RI" + fmt.Sprintf("%08d", maxNo+1)
	defer d.Close()
	return newUserID
}
