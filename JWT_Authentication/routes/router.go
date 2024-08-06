package routes

import (
	"auth/controller"
	"auth/middleware"

	"github.com/gin-gonic/gin"
)



func RegisterRoutes(router *gin.Engine) {



	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	// protected routes
	// router.Use(middleware.AuthMiddleware())

	// grouped routes

	tasks := router.Group("/api/tasks")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.GET("/" ,  controller.GetTasks)
		tasks.GET("/:id", controller.GetTasksByID)
		tasks.POST("/" ,  controller.CreateTask)
		tasks.DELETE("/:id" , controller.DeleteTask)
		tasks.PUT("/:id", controller.UpdateTask)
		tasks.GET("/mytask" , controller.GetTasksByUserID)

	}

	user := router.Group("/api/users")
	user.Use(middleware.AuthMiddleware())

	{

		user.GET("/getAllUser" , controller.GetAllUser)
		router.DELETE("/deleteUser" ,  controller.DeleteUser)
		router.GET("/tasks"  , controller.GetTasks)
	}
}
	// router.NoRoute(controller.NotFound)

