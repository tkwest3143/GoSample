package main

import("strike/go/db")

//
func main() {
	d := db.GormConnect()
	db.ChatDBInit(d)
	db.RoomDBInit(d)
	db.RoomUserRelationDBInit(d)
	db.UserDBInit(d)
	db.UserRelationDBInit(d)

	defer d.Close()
}
