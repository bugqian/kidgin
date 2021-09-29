package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Middlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Printf("中间件执行了%d秒\n", t2)
	}
}
