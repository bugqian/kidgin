package Apps

import (
	"github.com/gin-gonic/gin"
	"kidgin/Models"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	count := Models.Hello()
	c.JSON(http.StatusOK, gin.H{
		"num": count,
	})
}
