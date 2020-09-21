package main

import (
    "github.com/gin-gonic/gin"
	"net/http"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("content/html/*.html")

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(http.StatusOK, "index.html", gin.H{})
    })

    router.Run()
}