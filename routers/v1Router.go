package routers

import (
	api "GoWebLearningDemos/controllers"
	"GoWebLearningDemos/middlewares"

	"github.com/gin-gonic/gin"
)

func V1Router(r *gin.Engine) { //Uppercase first character of function name to export it
	v1Router := r.Group("/v1", middlewares.GrouplInitMiddlewares)
	{
		v1Router.GET("/hello", middlewares.InitMiddlewares, api.HelloController{}.SayHello)
		v1Router.GET("/", middlewares.InitMiddlewares, api.HelloController{}.HelloGin)
	}
}
