package common

import (
	"golang.org/x/crypto/bcrypt"
)

//DoCrypt 文字列をハッシュ化します
func DoCrypt(str string) string {

	hash, _ := bcrypt.GenerateFromPassword([]byte(str), 12)

	return string(hash)
}
