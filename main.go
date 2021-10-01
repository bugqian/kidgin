package main

import (
	"github.com/gin-gonic/gin"
	"kidgin/Routers"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("Views/*")
	Routers.LoadHello(r)
	r.Run(":8084")
}
