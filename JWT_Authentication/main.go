package main

import (
	"auth/routes"

	"github.com/gin-gonic/gin"
)




func main() {
  router := gin.Default()
  routes.RegisterRoutes(router)
  router.Run()
}