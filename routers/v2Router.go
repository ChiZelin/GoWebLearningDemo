package routers

import (
	api "GoWebLearningDemos/controllers"

	"github.com/gin-gonic/gin"
)

func V2Router(r *gin.Engine) {
	v2Router := r.Group("/v2")
	{
		v2Router.GET("/resource", api.ResourceController{}.Index)

		v2Router.POST("/resource", api.ResourceController{}.Add)

		v2Router.PUT("/resource", api.ResourceController{}.Edit)

		v2Router.DELETE("/resource", api.ResourceController{}.Delete)

		v2Router.GET("/ping", api.ResourceController{}.Pong)

		v2Router.GET("/news", api.ResourceController{}.News)

		v2Router.GET("/file/form", api.ResourceController{}.Form)

		v2Router.POST("/user/doUpload", api.ResourceController{}.DoUpload)
	}
}
