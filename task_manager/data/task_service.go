package data

import (
	"errors"
	"taskmanager/model"
	"time"
)



var tasks = []model.Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}


func GetTasks() *[]model.Task {
	return &tasks
}	

func GetTaskByID(id string) (*model.Task){
	
	for _ , t := range tasks{
		if t.ID == id {
			return &t
		}
	}

	return nil
}

func PostTasks(task *model.Task) *model.Task {

	if len(tasks) == 0 { 
		
	}
	tasks = append(tasks, *task)

	return task
	
}

func DeleteTask(id string) (*model.Task , error) {

	for i, task := range tasks {
		if task.ID == id{
			deletedData := task
			tasks = append(tasks[:i], tasks[i+1:]...)
			return &deletedData , nil
		}
	}

	return &model.Task{} , errors.New("not found")

}

func UpdateTask (id string , modifiedTask *model.Task) *model.Task  {


	existing := GetTaskByID(id)
	if existing == nil {
		return existing
	}
	
	if modifiedTask.Title != ""{
		existing.Title = modifiedTask.Title
	}
	if modifiedTask.Description != "" {
		existing.Description = modifiedTask.Description
	}
	if !modifiedTask.DueDate.IsZero(){
		existing.DueDate = modifiedTask.DueDate
	}
	if modifiedTask.Status != ""{ 
		existing.Status = modifiedTask.Status
	}

	return existing

}