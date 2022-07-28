package db

import (
	"strike/go/data"

	"gorm.io/gorm"
)

// WorkTimeテーブルの初期化を行います
func WorkTimeInit(d *gorm.DB) {
	d.AutoMigrate(&data.WorkTime{})

}

// WorkTime情報を挿入します
func WorkTimeInsert(insData data.WorkTime) {
	d := GormConnect()
	d.Create(&insData)
}

//userId,passwordをもとにusers情報を取得します。
func WorkTimeSelect(id string) []data.WorkTime {
	d := GormConnect()
	selData := []data.WorkTime{}
	d.Order("bote_date desc").Find(&selData, "id = ?", id)
	return selData
}
