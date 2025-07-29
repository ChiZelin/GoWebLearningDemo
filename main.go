package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Golang Gin!")
	})

	r.GET("/resource", func(c *gin.Context) {
		c.String(http.StatusOK, "这是一个GET请求, 用于获取数据")
	})

	r.POST("/resource", func(c *gin.Context) {
		c.String(http.StatusOK, "这是一个POST请求, 用于添加数据")
	})

	r.PUT("/resource", func(c *gin.Context) {
		c.String(http.StatusOK, "这是一个PUT请求, 用于修改数据")
	})

	r.DELETE("/resource", func(c *gin.Context) {
		c.String(http.StatusOK, "这是一个DELETE请求, 用于删除数据")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "最新新闻",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
