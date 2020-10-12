package main

import (
	"github.com/gin-gonic/gin"
	"github/GoSumple/gin_sumple/go/web"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("content/html/*")

	r.Static("/css", "./content/css/")
	r.Static("/js", "./content/js/")

	web.Router(r)

	r.Run(":8080")

}
