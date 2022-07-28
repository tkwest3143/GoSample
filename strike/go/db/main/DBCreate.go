package main

import (
	"fmt"
	"strike/go/db"
)

//
func main() {
	fmt.Printf("\x1b[32m%s\x1b[0m\n", "starting db create")
	d := db.GormConnect()
	db.WorkTimeInit(d)
	db.WorkTimeSettingInit(d)
	db.ChatInit(d)
	db.RoomInit(d)
	db.RoomUserRelationInit(d)
	db.UserInit(d)
	db.UserRelationInit(d)
	db.NewsMasterInit(d)
	fmt.Printf("\x1b[32m%s\x1b[0m\n", "finished db create!")
}
