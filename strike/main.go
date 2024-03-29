package main

import (
	"strike/go/common"
	"strike/go/web"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	appData := common.GetApplicationProperty()
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: appData.ALLOW_IP,
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
	//出力するログの設定
	writer := common.LogCreate()
	gin.DefaultWriter = writer
	port := common.GetApplicationProperty().ApplicationPort
	if len(port) == 0 {
		port = "8080"
	}

	web.Router(r)

	r.Run(":" + port)

}
