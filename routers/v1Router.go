package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func V1Router(r *gin.Engine) { //Uppercase first character of function name to export it
	v1Router := r.Group("/v1")
	{
		v1Router.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, this is v1 API!")
		})
		v1Router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, Golang Gin!")
		})
	}
}
