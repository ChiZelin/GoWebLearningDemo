package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func V2Router(r *gin.Engine) {
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
}
