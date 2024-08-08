package routers

import (
    "task7/Delivery/controllers"
    "task7/usecase"
    "task7/repository"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
    userRepo := repository.NewUserRepositoryImpl()
    userUsecase := usecase.NewUserUsecase(userRepo)
    userController := controllers.NewUserController(userUsecase)

    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/register", userController.RegisterUser)
        // Add more user-related endpoints here
    }
}
