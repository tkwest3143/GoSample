package db

import (
	"github.com/jinzhu/gorm"
	// PostgreSQL driver
	_ "github.com/lib/pq"
	"strike/go/common"
)

//GormConnect DBへの接続を行います
func GormConnect() *gorm.DB {
	appData := common.GetApplicationProperty()

	DBMS := appData.DbConfig.Dbms
	USER := appData.DbConfig.User
	PASS := appData.DbConfig.Password
	PROTOCOL := appData.DbConfig.Protocol
	DBNAME := appData.DbConfig.Dbname

	CONNECT := DBMS + "://" + USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?sslmode=disable"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
