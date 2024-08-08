// Delivery/routers/router.go
package routers

import (
	"log"
	"task7/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    // db connection
    err :=  godotenv.Load(".env")
    if err != nil { 
        log.Fatal("Error loading .env file")
    }
    bootstrap.ConnectDB()
    
    SetupUserRoutes(router)
    SetupTaskRoutes(router)


    return router
}

