package routes

import (
	"auth/controller"
	"auth/middleware"
	"os"

	"github.com/gin-gonic/gin"
)


var admincode = os.Getenv("admincode")

func RegisterRoutes(router *gin.Engine) {

	tasks := router.Group("/api/tasks")
	{
		tasks.GET("/", controller.GetTasks)
		tasks.GET("/:id", controller.GetTasksByID)
		tasks.POST("/", controller.CreateTask)
		tasks.DELETE("/:id", controller.DeleteTask)
		tasks.PUT("/:id", controller.UpdateTask)

	}
	user := router.Group("/api/users")
	{
		
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
		user.GET("/getAllUser" ,middleware.AuthMiddleware(admincode) ,  controller.GetAllUser)
		router.DELETE("/deleteUser" , middleware.AuthMiddleware(admincode) , controller.DeleteUser)
		router.GET("/tasks" , middleware.AuthMiddleware(admincode) , controller.GetTasks)
	}
	// router.NoRoute(controller.NotFound)

}