package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
	BaseController // Embedding BaseController to inherit its methods
}

func (con HelloController) SayHello(c *gin.Context) {
	con.Success(c)
	c.String(200, "Hello, this is v1 API!")
	fmt.Println("SayHello method called")
}

func (con HelloController) HelloGin(c *gin.Context) {
	c.String(200, "Hello, Golang Gin!")
}
