package router

import (
	"github.com/gin-gonic/gin"
	"taskmanager/controller"
	
)

func RegisterRoutes(router *gin.Engine) {
	
	tasks := router.Group("/api")
	{
		tasks.GET("/tasks", controller.GetTasks)
		tasks.GET("/tasks/:id", controller.GetTasksByID)
		tasks.POST("/tasks", controller.PostTasks)
		tasks.DELETE("/tasks/:id", controller.DeleteTask)
		tasks.PUT("/tasks/:id", controller.UpdateTask)

		
	}
	router.NoRoute(controller.NotFound)

	
}

	
