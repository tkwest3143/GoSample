package main

import (
	"strike/go/common"
	"strike/go/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//出力するログの設定
	writer := common.LogCreate()
	gin.DefaultWriter = writer
	appData := common.GetApplicationProperty()
	port := appData.ApplicationPort
	if len(port) == 0 {
		port = "8080"
	}

	r.LoadHTMLGlob("content/html/*")

	//htmlで参照するcss、jsのパスの設定
	r.Static("/css", "./content/css/")
	r.Static("/js", "./content/js/")

	web.Router(r)

	r.Run(":" + port)

}
