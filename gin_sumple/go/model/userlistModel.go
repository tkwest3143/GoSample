package model

import (
	"github/GoSumple/gin_sumple/go/data"
	"github/GoSumple/gin_sumple/go/db"
)

//GetUserFriendList userIDをもとに画面表示用フレンド情報を取得します。
func GetUserFriendList(userID string) []data.UserListData {
	d := db.GormConnect()

	//チャット情報の取得
	selData := []data.UserListData{}
	d.Raw("select "+
		"t2.user_id,t2.user_name, t1.relation_cd,t2.last_login, t2.user_name_search, t2.login_flg "+
		"from user_relations t1 "+
		"inner join users t2 on t1.user_id2 = t2.user_id "+
		"where t1.user_id1 = ? order by t2.user_name", userID).Scan(&selData)
	defer d.Close()
	return selData
}
