package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//GormConnect DBへの接続を行います
func GormConnect() *gorm.DB {
	DBMS := "postgres"
	USER := "USER"
	PASS := "password"
	PROTOCOL := "localhost:5432"
	DBNAME := "GODB"

	CONNECT := DBMS + "://" + USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?sslmode=disable"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
