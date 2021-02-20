// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ReturnLogin lotinError.htmlのGET処理を実装します
func ReturnLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
