package main

import (
	"auth/routes"

	"github.com/gin-gonic/gin"
)




func main() {
  // loading envs
  // err := godotenv.Load()
  router := gin.Default()
  routes.RegisterRoutes(router)
  router.Run()
}