package db

import (
	"github/GoSumple/gin_sumple/go/data"
)

//RoomUserRelationInsert RoomUserRelations情報を挿入します
func RoomUserRelationInsert(insData data.RoomUserRelation) {
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//RoomUserRelationSelectByUserID ユーザIDをもとにRoomUserRelations情報を取得します
func RoomUserRelationSelectByUserID(userID string) data.RoomUserRelation {
	d := GormConnect()
	selData := data.RoomUserRelation{}
	d.First(&selData, "user_id=?", userID)
	defer d.Close()
	return selData
}

//RoomUserRelationSelectByRoomID ルームIDをもとにRoomUserRelations情報を取得します
func RoomUserRelationSelectByRoomID(roomID string) []data.RoomUserRelation {
	d := GormConnect()
	selData := []data.RoomUserRelation{}
	d.First(&selData, "room_id=?", roomID)
	defer d.Close()
	return selData
}
