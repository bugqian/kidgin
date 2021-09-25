package main

import (
	"github.com/gin-gonic/gin"
	"kidgin/Routers"
)

func main() {
	r := gin.Default()
	Routers.LoadHello(r)
	r.Run(":8084")
}
