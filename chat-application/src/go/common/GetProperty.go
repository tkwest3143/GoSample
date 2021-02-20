package common

import (
	"encoding/json"
	"fmt"
	"github/GoSumple/chat-application/src/go/data"
	"io/ioutil"
	"os"
)

// GetApplicationProperty application.jsonからプロパティを取得する
func GetApplicationProperty() data.AppData {

	raw, err := ioutil.ReadFile("resource/application.json")
	if err != nil {
		fmt.Println("プロパティファイルが読み込めませんでした")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var appData data.AppData
	//読み込んだjsonファイルをデータに埋め込む
	json.Unmarshal(raw, &appData)

	return appData
}
