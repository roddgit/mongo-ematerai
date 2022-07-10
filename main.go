package main

import (
	"gin-mongo-api/configs"

	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// run database
	configs.ConnectDb()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"data": "hello from gin-gonic & mongob",
	// 	})

	// })

	routes.UserRoute(router)
	routes.EmateraiRoute(router)

	router.Run("localhost:8080")

}
