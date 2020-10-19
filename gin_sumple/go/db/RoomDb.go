package db

import (
	"fmt"
	"github/GoSumple/gin_sumple/go/data"
)

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
	fmt.Println(roomID)
	d.Find(&selData, "room_id=?", roomID)
	fmt.Println(selData)
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
