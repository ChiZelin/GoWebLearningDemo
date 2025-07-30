package main

import (
	"GoWebLearningDemos/middlewares"
	"GoWebLearningDemos/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob("templates/**/*")
	r.Use(middlewares.GlobalInitMiddlewares)

	routers.V1Router(r)
	routers.V2Router(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
