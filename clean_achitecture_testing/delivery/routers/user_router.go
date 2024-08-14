package routers

import (
	"task7/Delivery/controllers"
	"task7/bootstrap"
	"task7/infrastructure"
	"task7/repository"
	"task7/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)



func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Client, router *gin.RouterGroup) {

   
 
    userRepo := repository.NewUserRepositoryImpl(db ,  env.DBName, env.UserCollection)
    userUsecase := usecase.NewUserUsecase(userRepo)
    userController := controllers.NewUserController(userUsecase)


	router.DELETE("/:id", userController.DeleteUser)
	router.GET("/" , infrastructure.AdminMiddleware() ,userController.GetAllUsers)
	// router.PUT("/:id", userController.UpdateUser)

  
}


