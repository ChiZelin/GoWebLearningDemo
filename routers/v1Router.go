package routers

import (
	api "GoWebLearningDemos/controllers"

	"github.com/gin-gonic/gin"
)

func V1Router(r *gin.Engine) { //Uppercase first character of function name to export it
	v1Router := r.Group("/v1")
	{
		v1Router.GET("/hello", api.HelloController{}.SayHello)
		v1Router.GET("/", api.HelloController{}.HelloGin)
	}
}
