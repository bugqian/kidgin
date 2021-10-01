package Routers

import (
	"github.com/gin-gonic/gin"
	"kidgin/Apps"
	"kidgin/Middlewares"
	"kidgin/Validate"
)

func LoadHello(e *gin.Engine) {
	Validate.ValidateRegister()
	//Middlewares.Middlewares()
	e.POST("/login", Apps.Login)
	e.Use(Middlewares.AuthMiddleware())
	{
		e.GET("/hello", Apps.HelloHandler)
		e.GET("/world", Apps.WorldHandler)
		e.GET("/async", Apps.AsyncHandler)
		e.GET("/sync", Apps.SyncHandler)
		e.GET("/mid", Middlewares.Middlewares(), Apps.Middlerware) //局部中间件
		e.GET("/myWorld", Apps.MyWorldHandler)
		article := e.Group("/article")
		{
			article.POST("/add", Apps.ArticleAdd)
			article.PUT("/update/:id", Apps.ArticleUpdate)
			article.DELETE("/delete/:id", Apps.ArticleDelete)
		}
	}
}
