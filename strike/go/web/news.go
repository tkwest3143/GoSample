package web

import (
	"fmt"
	"net/http"
	"strike/go/data"
	"strike/go/model"

	"github.com/gin-gonic/gin"
)

func GetNewsByUser(ctx *gin.Context) {
	var data data.GetNewsDataByUser
	if ctx.ShouldBind(&data) == nil {
		fmt.Println(data)
	}
	xml_data, err := model.GetNewsListByUser("NM00000001")
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
func GetNewsByCategory(ctx *gin.Context) {
	var data data.GetNewsDataByCategory

	if ctx.ShouldBind(&data) == nil {
		if data.CategoryId == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "not request parameter category_id"})
			return
		}
		xml_data, err := model.GetNewsListByCategory(data.CategoryId)
		fmt.Println(err)
		if err == nil {

			ctx.JSON(http.StatusOK, gin.H{
				"news_list": xml_data,
			})
			return
		} else {
			panic(err)
		}
	} else {
		panic("エラーが発生")
	}

}
func GetNewsCategory(ctx *gin.Context) {
	data, err := model.GetAllNewsCategoryData()
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"news_category": data,
		})
		return
	} else {
		panic(err)
	}

}
