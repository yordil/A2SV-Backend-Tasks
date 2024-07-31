package main

import (
	"taskmanager/router"

	"github.com/gin-gonic/gin"
)

func main(){
	route := gin.Default()
	router.RegisterRoutes(route)
	route.Run("localhost:8080")

}