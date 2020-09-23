package web
import (
    "github.com/gin-gonic/gin"
)

func Router(){
	r := gin.Default()
	r.LoadHTMLGlob("content/html/*")
	r.GET("/", Index) // 初期画面
	r.Run(":8080")
}