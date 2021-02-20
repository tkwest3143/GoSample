package main

import (
	"github/GoSumple/chat-application/src/go/common"
	"github/GoSumple/chat-application/src/go/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//出力するログの設定
	writer := common.LogCreate()
	gin.DefaultWriter = writer

	r.LoadHTMLGlob("src/html/*")

	//htmlで参照するcss、jsのパスの設定
	r.Static("/css", "./src/css/")
	r.Static("/js", "./src/js/")

	web.Router(r)

	r.Run(":8081")

}
