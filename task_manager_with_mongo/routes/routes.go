package router

import (
	"github.com/gin-gonic/gin"
	"taskmanager/controllers"
	
)

func RegisterRoutes(router *gin.Engine) {
	
	tasks := router.Group("/api")
	{
		tasks.GET("/tasks", controllers.GetTasks)
		tasks.GET("/tasks/:id", controllers.GetTasksByID)
		tasks.POST("/tasks", controllers.CreateTask)
		tasks.DELETE("/tasks/:id", controllers.DeleteTask)
		tasks.PUT("/tasks/:id", controllers.UpdateTask)

		
	}
	// router.NoRoute(controller.NotFound)

	
}

	
