package api

import (
	"GoWebLearningDemos/models"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (con UserController) Add(c *gin.Context) {
	c.String(http.StatusOK, "这是一个GET请求, 用于获取数据")
}

func (con UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "这是一个PUT请求, 用于修改数据")
}

func (con UserController) Index(c *gin.Context) {

	userList := []models.Usertable{}

	models.DB.Find(&userList)

	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})

	c.String(http.StatusOK, "这是一个GET请求, 用于获取数据")
}

func (con UserController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "这是一个DELETE请求, 用于删除数据")
}

func (con UserController) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (con UserController) News(c *gin.Context) {
	c.HTML(http.StatusOK, "default/news.html", gin.H{
		"title": "最新新闻",
	})
}

func (con UserController) Form(c *gin.Context) {
	c.HTML(http.StatusOK, "default/useradd.html", gin.H{})
}

func (con UserController) DoUpload(c *gin.Context) {

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
