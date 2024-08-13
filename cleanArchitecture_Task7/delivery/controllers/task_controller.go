package controllers

import (
	"net/http"
	"task7/domain"

	"github.com/gin-gonic/gin"
)



type taskController struct{
	taskUsecase domain.TaskUsecase
}


func NewTaskController(taskUsecase domain.TaskUsecase ) *taskController {
	return &taskController{taskUsecase: taskUsecase}
}

func (tc *taskController) CreateTask(c *gin.Context){
	userid := c.GetString("user_id")
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.USERID = userid
	response := tc.taskUsecase.CreateTask(task)
	

	HandleResponse(c , response)
}

func (tc * taskController) GetTasksByUserID(c *gin.Context) {
	id := c.GetString("user_id")
	response := tc.taskUsecase.GetTasksByUserID(id)
	
	HandleResponse(c , response)

	
}


func (tc * taskController) GetTasks(c *gin.Context) {
	response := tc.taskUsecase.GetAllTask()
	
	HandleResponse(c , response)

	

}

func (tc * taskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	
	role := c.GetString("role")
	userid := c.GetString("user_id")
	
	response := tc.taskUsecase.DeleteTask(userid , id , role)

	HandleResponse(c , response)

}

func (tc * taskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	role := c.GetString("role")
	response := tc.taskUsecase.GetTaskByID(id , role)
	
	HandleResponse(c , response)

}


func (tc *taskController) UpdateTask(c *gin.Context){
	
	userid := c.GetString("user_id")
	taskid := c.Param("id")
	var tasks domain.Task
	if err := c.ShouldBindJSON(&tasks) ; err != nil {
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err})
		return 
	}

	// calling services

	response :=tc.taskUsecase.UpdateTask(userid , taskid , tasks)

	HandleResponse(c , response)

	
}

