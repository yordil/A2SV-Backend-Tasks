package main

import (
	"taskmanager/config"
	router "taskmanager/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()
	route := gin.Default()
	router.RegisterRoutes(route)
	route.Run("localhost:8080")
	// routes.Router.Run(config.EnvConfigs.LocalServerPort)
	 
} 