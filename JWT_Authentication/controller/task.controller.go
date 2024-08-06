package controller

import (
	"auth/data"
	"auth/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)	

func GetTasks(c *gin.Context) {
	
    tasks, err := data.GetTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch tasks"})
        return
    }
    c.JSON(http.StatusOK, tasks)
}


func CreateTask(c *gin.Context) {

	var task models.Task
	id := c.GetString("user_id")
	
	task.USERID = id

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := data.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task": task,
	})
}

func GetTasksByID(c *gin.Context) {
	id := c.Param("id")

	personID := c.GetString("user_id")
	role := c.GetString("role")
	if role != "thisisadmin" {
		owner  := data.UserIdGetter(personID , id)
		fmt.Println(owner , "******** This is owner ")
		if !owner { 
			c.JSON(http.StatusNotFound, gin.H{"error": "Not authorized to view this task"})
			return
		}
	}
	task , err := data.GetTaskByID(id)

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

func GetTasksByUserID(c *gin.Context) {
		
		id  := c.GetString("user_id")
		// fmt.Println(id , fmt.Sprintf("%T", id) , "********")
		tasks, err := data.GetTasksByUserID(id)

		if err != nil { 
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "Unable to fetch tasks"})
			return 
		}
		c.JSON(http.StatusOK, gin.H{"tasks" : tasks})

}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskid := c.Param("id")
	userid := c.GetString("user_id")

	// role := c.GetString("role")	
	
	flag := data.UserIdGetter(userid , taskid)
	fmt.Println(flag , "************************************")

	if !flag {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not authorized to update this task"})
		return
	}

	updatedTask , err := data.UpdateTask(taskid , &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": updatedTask})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	userid := c.GetString("user_id")

	flag := data.UserIdGetter(userid , id)

	if !flag {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not authorized to delete this task"})
		return
	}

	res , err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": res})
}

