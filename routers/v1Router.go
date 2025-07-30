package routers

import (
	api "GoWebLearningDemos/controllers"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func initMiddlewares(c *gin.Context) {
	start := time.Now().UnixNano()

	fmt.Println("Initializing middlewares...")

	c.Next() // Call the next handler in the chain

	//c.Abort() // Abort the chain to prevent further handlers from executing

	fmt.Println("Middlewares initialized.")

	end := time.Now().UnixNano()

	duration := end - start

	fmt.Printf("Execution time: %d nanoseconds\n", duration)
}

func V1Router(r *gin.Engine) { //Uppercase first character of function name to export it
	v1Router := r.Group("/v1")
	{
		v1Router.GET("/hello", initMiddlewares, api.HelloController{}.SayHello)
		v1Router.GET("/", initMiddlewares, api.HelloController{}.HelloGin)
	}
}
