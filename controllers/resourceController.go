package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceController struct{}

func (con ResourceController) Add(c *gin.Context) {
	c.String(http.StatusOK, "这是一个POST请求, 用于添加数据")
}

func (con ResourceController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "这是一个PUT请求, 用于修改数据")
}

func (con ResourceController) Index(c *gin.Context) {
	c.String(http.StatusOK, "这是一个GET请求, 用于获取数据")
}

func (con ResourceController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "这是一个DELETE请求, 用于删除数据")
}

func (con ResourceController) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (con ResourceController) News(c *gin.Context) {
	c.HTML(http.StatusOK, "default/news.html", gin.H{
		"title": "最新新闻",
	})
}
