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
