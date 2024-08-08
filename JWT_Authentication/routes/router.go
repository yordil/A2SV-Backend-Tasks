package routes

import (
	"auth/controller"
	"auth/middleware"

	"github.com/gin-gonic/gin"
)



func RegisterRoutes(router *gin.Engine) {



	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	tasks := router.Group("/api/tasks")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.GET("/" , middleware.AdminMiddleware() ,  controller.GetTasks) 
		tasks.GET("/:id", controller.GetTasksByID) 
		tasks.POST("/" ,  controller.CreateTask)  
		tasks.DELETE("/:id" , controller.DeleteTask)  
		tasks.PUT("/:id", controller.UpdateTask)  
		tasks.GET("/mytask" , controller.GetTasksByUserID)

	}

	user := router.Group("/api/users")
	user.Use(middleware.AuthMiddleware())

	{

		user.GET("/getAllUser" , middleware.AdminMiddleware() ,  controller.GetAllUser)
		user.DELETE("/deleteUser/:id" ,middleware.AdminMiddleware() ,   controller.DeleteUser)
	}
}
	// router.NoRoute(controller.NotFound)

