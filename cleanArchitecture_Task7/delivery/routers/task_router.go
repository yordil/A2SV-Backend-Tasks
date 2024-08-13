package routers

import (
	"task7/bootstrap"
	"task7/Delivery/controllers"
	"task7/repository"
	"task7/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Client, router *gin.RouterGroup) {


    
    taskRepo := repository.NewTaskRepositoryImpl(db ,  env.DBName, env.TaskCollection)
    taskUseCase := usecase.NewTaskUseCase(taskRepo)
    taskController := controllers.NewTaskController(taskUseCase)

	

			taskroute := router.Group("/task")
      {

		  taskroute.GET("/" ,  taskController.GetTasks) 
		  taskroute.GET("/:id", taskController.GetTaskByID) 
		  taskroute.POST("/" ,  taskController.CreateTask)  
		  taskroute.DELETE("/:id" , taskController.DeleteTask)
		  taskroute.PUT("/:id", taskController.UpdateTask)  
		  taskroute.GET("/mytask" , taskController.GetTasksByUserID)
	  }

    }

