package main

import (
	"stock-api-golang/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	routes.SetupRoutes(router)

	router.Run()
}