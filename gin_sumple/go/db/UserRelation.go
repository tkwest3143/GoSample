package db

import (
	"github/GoSumple/gin_sumple/go/data"
)

//UserRelationInsert userRelations情報を挿入します
func UserRelationInsert(insData data.UserRelation) {
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//UserRelationExist 引数に指定されたユーザが存在するかどうかを判定します.
func UserRelationExist(userID string) bool {
	d := GormConnect()
	selData := data.UserRelation{}
	d.First(&selData, "user_id1=? or user_id2 = ?", userID, userID)
	defer d.Close()
	if selData.UserID1 == "" {
		return false
	}
	return true

}
