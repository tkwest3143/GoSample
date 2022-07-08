// Package web WebアプリのGet、Post処理をまとめたパッケージ
package web

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strike/go/common"

	"github.com/gin-gonic/gin"
)

//Upload index.htmlのPOST処理（/login）を実装します
func Upload(ctx *gin.Context) {
	// マルチパートフォーム
	form, _ := ctx.MultipartForm()
	log.Println(form)
	files := form.File["upload"]
	log.Println(files)
	for _, file := range files {
		appData := common.GetApplicationProperty()
		log.Println(file.Filename)
		filename := filepath.Base(file.Filename)

		// 特定のディレクトリにファイルをアップロードする
		ctx.SaveUploadedFile(file, appData.UPLOAD_DIRECTORY+"/"+filename)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
