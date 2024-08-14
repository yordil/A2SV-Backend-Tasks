package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID       primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Title       string `json:"title"`
	USERID      string `json:"user_id"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetTaskByID(id string , role string) (Task, error)
	GetAllTask() ([]Task, error)
	DeleteTask(id string) error
	UpdateTask(id string , task Task) (Task, error)
	GetTaskByUserID(userID string) ([]Task, error)
	UserIdGetter(id string , taskid string) bool
}

type TaskUsecase interface {
	CreateTask(task Task) (interface{})
	GetTaskByID(id string , role string) (interface{})
	GetAllTask() (interface{})
	DeleteTask(userID string , taskID string , role string) interface{}
	UpdateTask(userID string , taskid string , task Task) (interface{})
	GetTasksByUserID(id string) (interface{})

}