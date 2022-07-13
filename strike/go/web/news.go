package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strike/go/model"
)

func GetNewsList(ctx *gin.Context) {
	xml_data, err := model.ReadXmlFromHttp("https://news.yahoo.co.jp/rss/categories/it.xml")
	fmt.Println(err)
	if err == nil {

		ctx.JSON(http.StatusOK, gin.H{
			"news_list": xml_data,
		})
		return
	} else {
		panic(err)
	}

}
