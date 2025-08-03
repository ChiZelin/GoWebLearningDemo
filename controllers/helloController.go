package api

import (
	"GoWebLearningDemos/models"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

type HelloController struct {
	BaseController // Embedding BaseController to inherit its methods
}

func (con HelloController) SayHello(c *gin.Context) {

	c.SetCookie("username", "cookietest", 3600, "/", "localhost", false, false)

	//设置session
	session := sessions.Default(c)
	session.Set("username", "sessiontest")
	session.Save()

	con.Success(c)
	username, exists := c.Get("username")
	if !exists {
		username = "unknown"
	}
	c.String(200, "Hello, this is v1 API! Username: %s", username)

	fmt.Println(models.UnixToTime(1629788576)) // Example usage of a model function
	fmt.Println("SayHello method called")
}

func (con HelloController) HelloGin(c *gin.Context) {

	username, _ := c.Cookie("username")

	c.String(200, "Hello, Golang Gin!\n")
	c.String(200, "cookie="+username)

	//获取session
	session := sessions.Default(c)
	sessionUsername := session.Get("username")
	c.String(200, " session="+sessionUsername.(string))
}
