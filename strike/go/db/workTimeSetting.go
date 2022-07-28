package db

import (
	"strike/go/data"

	"gorm.io/gorm"
)

// WorkTimeSettingテーブルの初期化を行います
func WorkTimeSettingInit(d *gorm.DB) {
	d.AutoMigrate(&data.WorkTime{})

}

// WorkTimeSetting情報を挿入します
func WorkTimeSettingInsert(insData data.WorkTime) {
	d := GormConnect()
	d.Create(&insData)
}

//userId,passwordをもとにusers情報を取得します。
func WorkTimeSettingSelect(id string) []data.WorkTime {
	d := GormConnect()
	selData := []data.WorkTime{}
	d.Order("bote_date desc").Find(&selData, "id = ?", id)
	return selData
}
