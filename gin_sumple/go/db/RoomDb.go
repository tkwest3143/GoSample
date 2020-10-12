package db

import (
	"fmt"
	"github/GoSumple/gin_sumple/go/data"
	"strconv"
)

//RoomInsert users情報を挿入します
func RoomInsert(insData data.Room) {
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//GetMaxRoomID IDの最大値を取得します。
func GetMaxRoomID() string {
	d := GormConnect()
	maxNo := 0
	d.Table("rooms").Count(&maxNo)
	newUserID := "RI" + fmt.Sprintf("%08d", strconv.Itoa(maxNo+1))
	defer d.Close()
	return newUserID
}
