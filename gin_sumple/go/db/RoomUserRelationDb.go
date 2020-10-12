package db

import (
	"github/GoSumple/gin_sumple/go/data"
)

//RoomUserRelationInsert users情報を挿入します
func RoomUserRelationInsert(insData data.RoomUserRelation) {
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}
