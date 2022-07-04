package db

import (
	"strike/go/data"

	"gorm.io/gorm"
)

//RoomUserRelationDBInit roomUserRelationsテーブルの初期化を行います
func RoomUserRelationInit(d *gorm.DB) {
	d.AutoMigrate(&data.RoomUserRelation{})
}

//RoomUserRelationInsert RoomUserRelations情報を挿入します
func RoomUserRelationInsert(insData data.RoomUserRelation) {
	d := GormConnect()
	d.Create(&insData)
}

//RoomUserRelationSelectByUserID ユーザIDをもとにRoomUserRelations情報を取得します
func RoomUserRelationSelectByUserID(userID string) data.RoomUserRelation {
	d := GormConnect()
	selData := data.RoomUserRelation{}
	d.First(&selData, "user_id=?", userID)
	return selData
}

//RoomUserRelationSelectByRoomID ルームIDをもとにRoomUserRelations情報を取得します
func RoomUserRelationSelectByRoomID(roomID string) []data.RoomUserRelation {
	d := GormConnect()
	selData := []data.RoomUserRelation{}
	d.First(&selData, "room_id=?", roomID)
	return selData
}
