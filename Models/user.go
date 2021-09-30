package Models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"kidgin/Utils"
)

type User struct {
	Id       int    `gorm:"primaryKey" json:"id" form:"id"`
	UserName string `json:"username" form:"username" column:"username"`
	PassWord string `json:"username" form:"username" column:"password"`
	Gender   int    `json:"gender" form:"gender" column:"gender"`
	Age      int    `json:"age" form:"age" column:"age"`
}

//登录
func (u *User) Login(UserName, PassWord string) (string, error) {
	var user User
	if record := db.Model(&user).Where("username = ?", UserName).First(&user).RowsAffected; record == 1 {
		data := []byte(PassWord)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)
		if md5str == PassWord {
			token, err := Utils.CreateJwtToken(user.Id, UserName)
			return token, err
		}
	}
	return "", errors.New("用户名或者密码不正确")
}
