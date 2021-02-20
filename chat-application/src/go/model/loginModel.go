package model

import (
	"github/GoSumple/chat-application/src/go/data"
	"github/GoSumple/chat-application/src/go/db"

	"golang.org/x/crypto/bcrypt"
)


func DoAuthentication(username string, password string) data.User{

	userinfo := db.UserSelectByUserName(username)
	err := bcrypt.CompareHashAndPassword([]byte(userinfo.Password), []byte(password))
	if err != nil {
		return data.User{}
	} 
	return userinfo
}