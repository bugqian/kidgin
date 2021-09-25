package Apps

import (
	"github.com/gin-gonic/gin"
	"kidgin/Utils"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	Utils.Config()
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}
