package main
import (
  "github.com/jinzhu/gorm"

    _ "github.com/lib/pq"
)

func gormConnect() *gorm.DB {
  DBMS     := "postgres"
  USER     := "postgres"
  PASS     := "password"
  PROTOCOL := "localhost:5432"
  DBNAME   := "GODB"

  CONNECT := DBMS+"://"+USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?sslmode=disable"
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  return db
}

func main(){
  db := gormConnect()
  
  defer db.Close()
}