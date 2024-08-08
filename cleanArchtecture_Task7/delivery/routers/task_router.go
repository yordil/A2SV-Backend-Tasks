package routers

import (
	"task7/bootstrap"
	"task7/Delivery/controllers"
	"task7/infrastructure"
	"task7/repository"
	"task7/usecase"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine) {

    collection := bootstrap.GetCollection("taskss")
    
    taskRepo := repository.NewTaskRepositoryImpl(collection)
    taskUseCase := usecase.NewTaskUseCase(taskRepo)
    taskController := controllers.NewTaskController(taskUseCase)

        tasks := router.Group("/tasks")
        tasks.Use(infrastructure.AuthMiddleware())
    {
        tasks.GET("/" , infrastructure.AdminMiddleware() ,  taskController.GetTasks) 
		tasks.GET("/:id", taskController.GetTaskByID) 
		tasks.POST("/" ,  taskController.CreateTask)  
		tasks.DELETE("/:id" , taskController.DeleteTask)
		tasks.PUT("/:id", taskController.UpdateTask)  
		tasks.GET("/mytask" , taskController.GetTasksByUserID)

    }
}
