package db

import (
	"fmt"
	"strike/go/data"

	"gorm.io/gorm"
)

//NewsMasterInit NewsMasterテーブルの初期化を行います
func NewsMasterInit(d *gorm.DB) {
	d.AutoMigrate(&data.NewsMaster{})
}

//NewsMasterInsert NewsMaster情報を挿入します
func NewsMasterInsert(insData data.NewsMaster) {
	d := GormConnect()
	id := GetMaxCategoryID()
	insData.ID = id
	d.Create(&insData)
}

/*
全情報を取得します。
*/
func NewsMasterSelectAll() []data.NewsMaster {
	d := GormConnect()
	selData := []data.NewsMaster{}
	d.Where("deleted_at IS NULL").Find(&selData)
	return selData
}

//NewsMasterSelectById IDをもとにusers情報を取得します。
func NewsMasterSelectById(categoryID string) data.NewsMaster {
	d := GormConnect()
	selData := data.NewsMaster{}
	d.First(&selData, "id=?", categoryID)
	return selData
}

//UserSelectByUserName userIdをもとにusers情報を取得します。
func NewsMasterSelectByTitle(title string) data.NewsMaster {
	d := GormConnect()
	selData := data.NewsMaster{}
	d.First(&selData, "title=?", title)
	return selData
}

//GetMaxCategoryID category IDの最大値を取得します。
func GetMaxCategoryID() string {
	d := GormConnect()
	var maxNo int64
	d.Table("news_masters").Count(&maxNo)
	newCategoryID := "NM" + fmt.Sprintf("%08d", maxNo+1)
	return newCategoryID
}
