// Delivery/routers/router.go
package routers

import (
	"task7/bootstrap"
	"task7/infrastructure"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(env *bootstrap.Env , timeout time.Duration , db *mongo.Client , gin *gin.Engine)  {
    publicRouter := gin.Group("")

	NewSignUPRouter(env , timeout , db , publicRouter)
	NewLoginRouter(env , timeout , db , publicRouter)

	privateRouter := gin.Group("")	
	privateRouter.Use(infrastructure.AuthMiddleware())

	// NewTaskRouter(env , timeout , db , privateRouter)
	NewUserRouter(env , timeout , db , privateRouter)
	NewTaskRouter(env , timeout , db , privateRouter)

	
	
}

