package Routers

import (
	"github.com/gin-gonic/gin"
	"kidgin/Apps"
)

func LoadHello(e *gin.Engine) {
	e.GET("/hello", Apps.HelloHandler)
}
