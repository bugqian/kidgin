package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kidgin/Utils"
	"net/http"
)

//中间件解析Token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		tokeString := c.GetHeader("token")
		fmt.Println(tokeString, "当前token")
		if tokeString == "" {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"code":    1005,
				"message": "必须传递token",
			})
			c.Abort()
			return
		}
		claims, err := Utils.ParseJwtToken(tokeString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1005,
				"message": "token解析错误",
			})
			c.Abort()
			return
		}
		// 从token中解析出来的数据挂载到上下文上,方便后面的控制器使用
		c.Set("Uid", claims.Uid)
		c.Set("UserName", claims.UserName)
		c.Next()
	}
}
