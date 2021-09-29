package Apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func HelloHandler(c *gin.Context) {
	//count := Models.Hello()
	//1.json返回
	//c.JSON(http.StatusOK, gin.H{"message": "ok", "code": 200})
	//2.struct返回
	//var msg struct {
	//	Name string `json:"name"`
	//	Age  int    `json:"age"`
	//	Msg  string `json:"msg"`
	//}
	//msg.Age = 29
	//msg.Name = "bugQian"
	//msg.Msg = "Success"
	//c.JSON(http.StatusOK, msg)
	//3.XML返回
	//c.XML(200, gin.H{"msg": "Success"})
	//4.Yaml返回
	//c.YAML(200, gin.H{"message": "error"})
	//5.重定向
	c.Redirect(302, "http://www.baidu.com")
}

func WorldHandler(c *gin.Context) {
	c.HTML(200, "world.html", gin.H{"title": "我的世界", "name": "bugqian"})
}

//异步
func AsyncHandler(c *gin.Context) {
	copyContext := c.Copy()
	//异步处理
	fmt.Println("test1")
	go func() {
		//time.Sleep(3 * time.Second)
		log.Println("异步执行：" + copyContext.Request.URL.Path)
	}()
	fmt.Println("test2")
}

//同步
func SyncHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行：" + c.Request.URL.Path)
	fmt.Println("test3")
}

func Middlerware(c *gin.Context) {
	fmt.Println(c.Get("request"))
	req, _ := c.Get("request")
	c.String(200, req.(string))
}
