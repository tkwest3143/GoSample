package main

import (
	"github.com/gin-gonic/gin"
	"github/GoSumple/gin_sumple/go/common"
	"github/GoSumple/gin_sumple/go/web"
)

func main() {
	r := gin.Default()
	//出力するログの設定
	writer := common.LogCreate()
	gin.DefaultWriter = writer

	r.LoadHTMLGlob("content/html/*")

	//htmlで参照するcss、jsのパスの設定
	r.Static("/css", "./content/css/")
	r.Static("/js", "./content/js/")

	web.Router(r)

	r.Run(":8080")

}
