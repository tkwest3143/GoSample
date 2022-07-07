// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//Upload index.htmlのPOST処理（/login）を実装します
func Upload(ctx *gin.Context) {
	// マルチパートフォーム
	form, _ := ctx.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		filename := filepath.Base(file.Filename)

		// 特定のディレクトリにファイルをアップロードする
		ctx.SaveUploadedFile(file, filename)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
