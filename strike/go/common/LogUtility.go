package common

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// ErrLvl エラーレベル
type ErrLvl int

const (
	_ ErrLvl = iota
	// TRACE トレース
	TRACE
	// INFO 情報
	INFO
	// WARM 警告
	WARM
	// ERROR 異常
	ERROR
)

var errLvlStrings = [5]string{"未定義", "TRACE", "INFO", "WARM", "ERROR"}

func (s ErrLvl) String() string {
	return errLvlStrings[s]
}

// LogCreate ログを作成する
func LogCreate() io.Writer {
	//アプリケーションプロパティからログファイルのパスを取得し、出力先に設定
	appData := GetApplicationProperty()
	if _, err := os.Stat(appData.LogPath.Directory); os.IsNotExist(err) {
		fmt.Println("ディレクトリを新たに作成します。")
		err = os.Mkdir(appData.LogPath.Directory, 0777)
		fmt.Println(err)
	}
	infoPath := appData.LogPath.Directory + "/" + appData.LogPath.Info
	infoFile, err := os.OpenFile(infoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		fmt.Println("ログファイルがありませんでした。新たに作成します。")
		infoFile, err = os.Create(infoPath)
		if err != nil {
			fmt.Println("ログファイルが作成できませんでした")
			fmt.Println(err)
			return nil
		}
	}

	multiLogFile := io.MultiWriter(os.Stdout, infoFile)
	conf := gin.LoggerConfig{}
	conf.Output = infoFile
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.SetOutput(multiLogFile)
	return multiLogFile
}
