package main

import (
	"GoWebLearningDemos/routers"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func initMiddlewares(c *gin.Context) {
	start := time.Now().UnixNano()

	fmt.Println("Initializing Global middlewares...")

	c.Next() // Call the next handler in the chain

	//c.Abort() // Abort the chain to prevent further handlers from executing

	fmt.Println("Middlewares Global initialized.")

	end := time.Now().UnixNano()

	duration := end - start

	fmt.Printf("Execution time: %d nanoseconds\n", duration)
}

func main() {
	r := gin.Default()
	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob("templates/**/*")
	r.Use(initMiddlewares)

	routers.V1Router(r)
	routers.V2Router(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
