package db

import (
	"strike/go/data"

	"gorm.io/gorm"
)

//UserRelationInit UserRelationsテーブルの初期化を行います
func UserRelationInit(d *gorm.DB) {
	d.AutoMigrate(&data.UserRelation{})
}

//UserRelationInsert userRelations情報を挿入します
func UserRelationInsert(insData data.UserRelation) {
	d := GormConnect()
	d.Create(&insData)
}

//GetFriendList 引数に指定されたユーザがをもとにフレンドリストを取得します。
func GetFriendList(userID string) []string {
	d := GormConnect()
	selData := []data.UserRelation{}
	d.Find(&selData, "user_id1=?", userID)
	friendList := []string{}
	for _, b := range selData {
		if userID != b.UserID1 {
			friendList = append(friendList, b.UserID1)
		} else {
			friendList = append(friendList, b.UserID2)
		}
	}
	return friendList
}
