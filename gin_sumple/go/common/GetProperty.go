package common

import (
	"encoding/json"
	"fmt"
	"github/GoSumple/gin_sumple/go/data"
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

	json.Unmarshal(raw, &appData)

	return appData
}
