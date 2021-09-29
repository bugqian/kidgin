package Apps

import (
	"github.com/gin-gonic/gin"
	"kidgin/Validate"
	"net/http"
)

type user struct {
	Name string `json:"name" form:"name" binding:"required,NameNotNullAndAdmin" reg_error_info:"名字不能为空"`
	Age  int    `json:"age" form:"age" binding:"required,gt=18" reg_error_info:"年龄不能为空并且必须大于18岁"`
}

func MyWorldHandler(c *gin.Context) {
	var user user
	if e := c.ShouldBind(&user); e == nil {
		c.JSON(http.StatusOK, user)
	} else {
		errInfo := Validate.ProcessErr(user, e)
		if len(errInfo) != 0 {
			c.String(201, errInfo)
			return
		}
		c.String(200, "success")
	}
}
