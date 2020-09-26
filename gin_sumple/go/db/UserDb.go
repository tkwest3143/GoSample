package db
import (
	"github/GoSumple/gin_sumple/go/data"
)

func UserInsert() {
	insData := data.User{}
	insData.UserId = "111"
	insData.Password = "password"
	insData.RegistDate = "2018-04-01"
	insData.LastLogin = "2018-04-01"
	d := GormConnect()
	d.CreateTable(&data.User{})
	d.Create(&insData)
	defer d.Close()
}

func  UserSelect(userId string,password string) data.User{
	d := GormConnect()
	selData := data.User{}
	//d.Where("user_id = ? ",userId).First(&selData)
	d.First(&selData,"user_id=? and password = ?",userId,password)
	defer d.Close()
	return selData
}