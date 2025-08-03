package main

import (
	"GoWebLearningDemos/middlewares"
	"GoWebLearningDemos/models"
	"GoWebLearningDemos/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB()
	r := gin.Default()
	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob("templates/**/*")
	r.Use(middlewares.GlobalInitMiddlewares)

	//配置session中间件
	//store := cookie.NewStore([]byte("secret1111"))

	//分布式场景 session 存在 redis中
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", "", []byte("secret1111"))
	r.Use(sessions.Sessions("mysession", store))

	routers.V1Router(r)
	routers.V2Router(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
