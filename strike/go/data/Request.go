package data

type DoLoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DoRegisterData struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	MarilAddress string `json:"mail_address"`
}

type GetNewsData struct {
	Filename string `json:"filename"`
}

type GetNewsDataByUser struct {
	UserId string `json:"user_id"`
}

type GetNewsDataByCategory struct {
	CategoryId string `json:"category_id"`
}
