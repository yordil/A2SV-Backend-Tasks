// Delivery/routers/router.go
package routers

import (
	"github.com/gin-gonic/gin"
    "task7/bootstrap"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    // db connection
    bootstrap.ConnectDB()
    
    SetupUserRoutes(router)
    // SetupTaskRoutes(router)


    return router
}

