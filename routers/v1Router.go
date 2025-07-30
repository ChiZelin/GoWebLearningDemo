package routers

import (
	api "GoWebLearningDemos/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func initMiddlewares(c *gin.Context) {
	fmt.Println("Initializing middlewares...")
}

func V1Router(r *gin.Engine) { //Uppercase first character of function name to export it
	v1Router := r.Group("/v1")
	{
		v1Router.GET("/hello", initMiddlewares, api.HelloController{}.SayHello)
		v1Router.GET("/", initMiddlewares, api.HelloController{}.HelloGin)
	}
}
