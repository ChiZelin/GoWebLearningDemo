package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob("templates/**/*")

	v1Router := r.Group("/v1")
	{
		v1Router.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, this is v1 API!")
		})
		v1Router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, Golang Gin!")
		})
	}

	v2Router := r.Group("/v2")
	{
		v2Router.GET("/resource", func(c *gin.Context) {
			c.String(http.StatusOK, "这是一个GET请求, 用于获取数据")
		})

		v2Router.POST("/resource", func(c *gin.Context) {
			c.String(http.StatusOK, "这是一个POST请求, 用于添加数据")
		})

		v2Router.PUT("/resource", func(c *gin.Context) {
			c.String(http.StatusOK, "这是一个PUT请求, 用于修改数据")
		})

		v2Router.DELETE("/resource", func(c *gin.Context) {
			c.String(http.StatusOK, "这是一个DELETE请求, 用于删除数据")
		})

		v2Router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		v2Router.GET("/news", func(c *gin.Context) {
			c.HTML(http.StatusOK, "default/news.html", gin.H{
				"title": "最新新闻",
			})
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
