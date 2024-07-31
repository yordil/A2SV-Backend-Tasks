package controller

import (
	"net/http"
	"taskmanager/data"
	"taskmanager/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTasks(c *gin.Context) {
	tasks := data.GetTasks()
	
	if len(*tasks) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No tasks found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTasksByID(c *gin.Context) {
	id := c.Param("id")

	task := data.GetTaskByID(id)

	if task == nil {
		c.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task Not Found"})
		return 
	}
	c.IndentedJSON(http.StatusOK, data.GetTaskByID(id))
}
func PostTasks(c *gin.Context) {
	
	var newTask model.Task
	randomId := uuid.New().String()
	newTask.ID = randomId

	err := c.BindJSON(&newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusCreated, data.PostTasks(&newTask))
	

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	deleted_task , err := data.DeleteTask(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "the task Not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message" : "task deleted successfully", 
		"task" : deleted_task ,
	})
}

func UpdateTask(c *gin.Context) {

	id := c.Param("id")

	var newtask model.Task
	
	if err := c.ShouldBindJSON(&newtask); err != nil {
		c.IndentedJSON(http.StatusBadRequest , gin.H{"messgae":  err.Error()})
	}

	res :=  data.UpdateTask(id , &newtask)
	if res == nil {
		c.IndentedJSON(http.StatusNotFound , gin.H{
			"message" : "task not found" ,
		})
		return 
	}

	c.IndentedJSON(http.StatusOK , res)

}

func NotFound(c * gin.Context){
	c.IndentedJSON(http.StatusNotFound , gin.H{
		"message" : "The route Not Found",
	})
}