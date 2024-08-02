package controllers

import (
	"net/http"
	"taskmanager/models"
	"taskmanager/services"

	"github.com/gin-gonic/gin"
)	

func GetTasks(c *gin.Context) {
	
    tasks, err := services.GetTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch tasks"})
        return
    }

    c.JSON(http.StatusOK, tasks)
}


func CreateTask(c *gin.Context) {

	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error  *********": err.Error()})
		return
	}

	err := services.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error #########": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetTasksByID(c *gin.Context) {
	id := c.Param("id")

	task , err := services.GetTaskByID(id)

	if task == nil {
		c.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task Not Found"})
		return 
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "Unable to fetch task"})
		return 
	}

	
	c.IndentedJSON(http.StatusOK, gin.H{"task" : task})
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	updatedTask , err := services.UpdateTask(id , &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": updatedTask})
}


func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	res , err := services.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": res})
}


