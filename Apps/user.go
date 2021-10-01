package Apps

import (
	"github.com/gin-gonic/gin"
	"kidgin/Models"
)

type LoginInfo struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginInfo LoginInfo
	if err := c.ShouldBind(&loginInfo); err != nil {
		c.JSON(201, err.Error())
		return
	}
	var user *Models.User
	token, err := user.Login(loginInfo.UserName, loginInfo.PassWord)
	if err != nil {
		c.JSON(201, gin.H{"msg": err.Error(), "token": token})
		return
	}
	c.JSON(200, gin.H{"msg": "success", "token": token})
}
