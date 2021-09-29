package main

import (
	"github.com/gin-gonic/gin"
	"kidgin/Routers"
	"kidgin/Validate"
)

func main() {

	r := gin.Default()
	Validate.ValidateRegister()
	r.LoadHTMLGlob("Views/*")
	Routers.LoadHello(r)
	r.Run(":8084")
}
