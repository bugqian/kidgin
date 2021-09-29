package Routers

import (
	"github.com/gin-gonic/gin"
	"kidgin/Apps"
	"kidgin/Middlewares"
)

func LoadHello(e *gin.Engine) {
	e.GET("/hello", Apps.HelloHandler)
	e.GET("/world", Apps.WorldHandler)
	e.GET("/async", Apps.AsyncHandler)
	e.GET("/sync", Apps.SyncHandler)
	e.GET("/mid", Middlewares.Middlewares(), Apps.Middlerware) //局部中间件
	e.GET("/myWorld", Apps.MyWorldHandler)
}
