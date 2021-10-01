package Apps

import (
	"github.com/gin-gonic/gin"
	"kidgin/Models"
	"net/http"
	"strconv"
)

/*
	添加文章
*/
func ArticleAdd(c *gin.Context) {
	var article Models.Article
	var err error
	if err = c.ShouldBind(&article); err == nil {
		err = article.Add(article)
		if err == nil {
			c.JSON(200, gin.H{"code": 201, "msg": "添加成功"})
			return
		}
	}
	c.JSON(200, gin.H{"code": 201, "msg": err.Error()})
	return
}

func ArticleUpdate(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	var article Models.Article
	var err error
	if err = c.ShouldBind(&article); err == nil {
		a, er := article.Info(idInt)
		if er != nil || a == (Models.Article{}) {
			c.JSON(http.StatusNotFound, gin.H{"code": 201, "msg": "数据找不到"})
			return
		}
		err = article.Update(article, idInt)
		if err == nil {
			c.JSON(200, gin.H{"code": 200, "msg": "修改成功"})
			return
		}
	}
	c.JSON(202, gin.H{"code": 201, "msg": err.Error()})
	return
}

func ArticleDelete(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	var article Models.Article
	err := article.Delete(idInt)
	if err != nil {
		c.JSON(202, gin.H{"code": 202, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "删除成功"})
	return
}
