package main

import (
	"GoSosmed/config"
	"GoSosmed/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()
	// config.AutoMigrateModels()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "jnck",
		})
	})

	router.AuthRouter(api)
	router.PostRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
