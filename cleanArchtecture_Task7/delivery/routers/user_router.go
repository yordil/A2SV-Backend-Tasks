package routers

import (
	"task7/Delivery/controllers"
	"task7/bootstrap"
	"task7/infrastructure"
	"task7/repository"
	"task7/usecase"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {

    collection := bootstrap.GetCollection("users")
    
    userRepo := repository.NewUserRepositoryImpl(collection)
    userUsecase := usecase.NewUserUsecase(userRepo)
    userController := controllers.NewUserController(userUsecase)

    userRoutes := router.Group("/users")
    userRoutes.POST("/register", userController.RegisterUser)
    userRoutes.POST("/login", userController.Login)

    userRoutes.Use(infrastructure.AuthMiddleware())
    {
        userRoutes.DELETE("/:id" , userController.DeleteUser)
        userRoutes.GET("/" ,  infrastructure.AdminMiddleware() , userController.GetAllUsers)

    //     userRoutes.GET("/" , userController.GetAllUsers)

    }
  
}


