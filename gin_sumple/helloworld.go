package main

import (
    "github.com/gin-gonic/gin"
	"net/http"
	"go/db"
)

func main() {
	Text:=""
    router := gin.Default()
    router.LoadHTMLGlob("content/html/*.html")

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(http.StatusOK, "index.html", gin.H{
		"text": "hello",
		})
    })
	
	 router.POST("/new", func(c *gin.Context) {
      text := c.PostForm("text")
      Text=text
      c.Redirect(302, "/")
   })

    router.Run()
}