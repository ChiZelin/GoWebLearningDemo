package api

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type ResourceController struct{}

func (con ResourceController) Add(c *gin.Context) {
	c.String(http.StatusOK, "这是一个GET请求, 用于获取数据")
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

func (con ResourceController) Form(c *gin.Context) {
	c.HTML(http.StatusOK, "default/useradd.html", gin.H{})
}

func (con ResourceController) DoUpload(c *gin.Context) {

	username := c.PostForm("username")
	file, err := c.FormFile("face")
	dst := path.Join("./static/upload", file.Filename)

	if err == nil {
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "上传成功",
		"username": username,
		"file":     file.Filename,
		"dst":      dst,
	})
}
