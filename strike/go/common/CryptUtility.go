package common

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
)

//DoCrypt 文字列をハッシュ化します
func DoCrypt(str string) string {

	hash, _ := bcrypt.GenerateFromPassword([]byte(str), 12)

	return string(hash)
}

// ランダム文字列作成
func GetRandom() (string, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return string(b), err
}
