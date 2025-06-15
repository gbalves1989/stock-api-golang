package routes

import (
	"stock-api-golang/app/http/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/hello", controllers.Hello)
	router.POST("/signup", controllers.SignUp)
	router.POST("/signin", controllers.SignIn)
}