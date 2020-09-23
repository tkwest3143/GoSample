package main

import (
	"github/GoSumple/gin_sumple/go/db"
	"github/GoSumple/gin_sumple/go/web"
)

func main() {
	d := db.GormConnect()
	defer d.Close()
    web.Router()
}