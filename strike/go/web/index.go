// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Index index.htmlのGET処理を実装します
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
