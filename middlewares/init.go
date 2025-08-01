package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddlewares(c *gin.Context) {
	start := time.Now().UnixNano()

	fmt.Println("Initializing middlewares...")

	c.Next() // Call the next handler in the chain

	//c.Abort() // Abort the chain to prevent further handlers from executing

	fmt.Println("Middlewares finished.")

	end := time.Now().UnixNano()

	duration := end - start

	fmt.Printf("Execution time: %d nanoseconds\n", duration)
}

func GlobalInitMiddlewares(c *gin.Context) {
	start := time.Now().UnixNano()

	fmt.Println("Initializing Global middlewares...")

	c.Next() // Call the next handler in the chain

	//c.Abort() // Abort the chain to prevent further handlers from executing

	fmt.Println("Middlewares Global finished.")

	end := time.Now().UnixNano()

	duration := end - start

	fmt.Printf("Execution time: %d nanoseconds\n", duration)
}

func GrouplInitMiddlewares(c *gin.Context) {
	fmt.Println("Initializing Group middlewares...")

	c.Set("username", "testuser") // Example of setting a value in the context

	cCp := c.Copy() // Create a copy of the context for use in the middleware
	go func() {
		time.Sleep(2 * time.Second) // Simulate some processing delay
		fmt.Println("Done! in path:", cCp.Request.URL.Path)
	}()

	c.Next() // Call the next handler in the chain

	//c.Abort() // Abort the chain to prevent further handlers from executing

	fmt.Println("Middlewares Group finished.")

}
