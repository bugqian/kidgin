package Apps

import (
	"github.com/gin-gonic/gin"
	"kidgin/Models"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	Models.Hello()
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}
