package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func EmateraiRoute(router *gin.Engine) {
	router.POST("/ematerai/login", controllers.LoginHandler())
}
