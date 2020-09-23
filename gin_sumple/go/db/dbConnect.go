package main
import (
  "github.com/jinzhu/gorm"

    _ "github.com/lib/pq"
)

func gormConnect() *gorm.DB {
  DBMS     := "postgres"
  USER     := "postgres"
  PASS     := "Aa13572468"
  PROTOCOL := "localhost:5432"
  DBNAME   := "TKDB"

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