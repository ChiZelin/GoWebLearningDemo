package api

import "github.com/gin-gonic/gin"

type HelloController struct{}

func (con HelloController) SayHello(c *gin.Context) {
	c.String(200, "Hello, this is v1 API!")
}

func (con HelloController) HelloGin(c *gin.Context) {
	c.String(200, "Hello, Golang Gin!")
}
