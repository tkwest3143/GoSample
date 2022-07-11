package data

// Dbconfig DB接続情報
type Dbconfig struct {
	Dbms     string `json:"DBMS"`
	User     string `json:"USER"`
	Password string `json:"PASSWORD"`
	Protocol string `json:"PROTOCOL"`
	Dbname   string `json:"DBNAME"`
}

// AppData アプリケーション情報
type AppData struct {
	ApplicationMode  int      `json:"application_mode"`
	ALLOW_IP         []string `json:"allow_ip"`
	UPLOAD_DIRECTORY string   `json:"upload_directory"`
	ApplicationPort  string   `json:"application_port"`
	DbConfig         Dbconfig `json:"dbconfig"`
	LogPath          LogPath  `json:"logPath"`
}

// LogPath ログファイルパス
type LogPath struct {
	Directory string `json:"directory"`
	Info      string `json:"info"`
	Error     string `json:"error"`
}
