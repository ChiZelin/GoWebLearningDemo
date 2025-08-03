package routers

import (
	api "GoWebLearningDemos/controllers"

	"github.com/gin-gonic/gin"
)

func V2Router(r *gin.Engine) {
	v2Router := r.Group("/v2")
	{
		v2Router.GET("/resource", api.UserController{}.Index)

		v2Router.POST("/resource", api.UserController{}.Add)

		v2Router.PUT("/resource", api.UserController{}.Edit)

		v2Router.DELETE("/resource", api.UserController{}.Delete)

		v2Router.GET("/ping", api.UserController{}.Pong)

		v2Router.GET("/news", api.UserController{}.News)

		v2Router.GET("/file/form", api.UserController{}.Form)

		v2Router.POST("/user/doUpload", api.UserController{}.DoUpload)
	}
}
