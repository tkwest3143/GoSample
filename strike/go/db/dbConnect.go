package db

import (
	"strike/go/common"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//GormConnect DBへの接続を行います
func GormConnect() *gorm.DB {
	appData := common.GetApplicationProperty()

	DB := appData.DbConfig.Dbname

	db, err := gorm.Open(sqlite.Open(DB), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
