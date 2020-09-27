/*
 */
package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//ReturnLogin lotinError.htmlのGET処理を実装します
func ReturnLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
