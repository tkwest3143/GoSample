package web
import (
    "github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context){
        ctx.HTML(http.StatusOK, "index.html", gin.H{
		"text": "hello",
		})
    }