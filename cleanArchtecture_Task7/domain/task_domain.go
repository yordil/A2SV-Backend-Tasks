package domain

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	USERID      string `json:"user_id"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetTaskByID(id int) (Task, error)
	GetAllTask() ([]Task, error)
	DeleteTask(id int) error
	UpdateTask(task Task) (Task, error)
	GetTaskByUserID(userID int) ([]Task, error)
}

type TaskUsecase interface {
	CreateTask(task Task) (Task, error)
	GetTaskByID(id int) (Task, error)
	GetAllTask() ([]Task, error)
	DeleteTask(id int) error
	UpdateTask(task Task) (Task, error)
	GetTaskByUserID(userID int) ([]Task, error)
	
}