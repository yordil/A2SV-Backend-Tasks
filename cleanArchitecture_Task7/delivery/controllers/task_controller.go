package controllers

import (
	"net/http"
	"task7/domain"
	"task7/usecase"

	"github.com/gin-gonic/gin"
)



type taskController struct{
	taskUsecase *usecase.TaskUseCase
}


func NewTaskController(taskUsecase *usecase.TaskUseCase) *taskController {
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
	task, err := tc.taskUsecase.CreateTask(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task, "message": "task created successfully"})
	

}

func (tc * taskController) GetTasksByUserID(c *gin.Context) {
	id := c.GetString("user_id")
	tasks, err := tc.taskUsecase.GetTasksByUserID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error this": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}


func (tc * taskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskUsecase.GetAllTask()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

func (tc * taskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	
	role := c.GetString("role")
	userid := c.GetString("user_id")
	
	err := tc.taskUsecase.DeleteTask(userid , id , role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

func (tc * taskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	role := c.GetString("role")
	task, err := tc.taskUsecase.GetTaskByID(id , role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
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

	task , err :=tc.taskUsecase.UpdateTask(userid , taskid , tasks)

	if err != nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"err" : err })
	}
	
	
	c.IndentedJSON(http.StatusAccepted , gin.H{"message" :  "Task Updates Successfully" , "task" : task})
	
}

