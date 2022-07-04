package main

import (
	"strike/go/db"
)

//
func main() {
	d := db.GormConnect()
	db.ChatInit(d)
	db.RoomInit(d)
	db.RoomUserRelationInit(d)
	db.UserInit(d)
	db.UserRelationInit(d)

}
