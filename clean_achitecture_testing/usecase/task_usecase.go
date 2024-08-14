package usecase

import (
	"task7/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) domain.TaskUsecase {
	return &TaskUseCase{taskRepo: taskRepo}
}

func (tu *TaskUseCase) CreateTask(task domain.Task) interface{} {
	if task.Description == "" || task.DueDate == "" || task.Status == "" || task.Title == "" || task.USERID == "" {
		return &domain.ErrorResponse{Message: "missing required fields",  Status: 400}
	}

	id := primitive.NewObjectID()
	task.ID = id
	task, err := tu.taskRepo.CreateTask(task)
	if err != nil { 
		return &domain.ErrorResponse{Message: "cannot create task", Status: 500}
	}

	return &domain.TaskSuccessResponse{Message: "task created successfully", Status: 200, Data: task}
}


func (tu *TaskUseCase) GetTasksByUserID(id string) (interface{}) {
	
	tasks, err := tu.taskRepo.GetTaskByUserID(id)

	if err != nil { 
		return &domain.ErrorResponse{Message: "cannot get tasks", Status: 500}
	}

	return &domain.TaskResponse{All_Task: tasks}
}

func (tu *TaskUseCase) GetAllTask() (interface{}) {
	tasks, err := tu.taskRepo.GetAllTask()

	if err != nil {
		return &domain.ErrorResponse{Message: "cannot get tasks", Status: 500}
	}

	return &domain.TaskResponse{All_Task: tasks}
}


func (tu *TaskUseCase) DeleteTask(userID string , taskID string , role string) (interface{}) {
	
	
	if role == "admin" || tu.taskRepo.UserIdGetter(userID , taskID) || role == "superAdmin" {
		err := tu.taskRepo.DeleteTask(taskID)
		
		if err != nil {
			return &domain.ErrorResponse{Message: "cannot delete task", Status: 500}
		}

		return &domain.TaskSuccessResponse{Message: "task deleted successfully", Status: 200}
		
	}

	return &domain.ErrorResponse{Message: "you are not authorized to delete this task", Status: 401}
}	


func (tu *TaskUseCase) GetTaskByID(id string , role string) (interface{}) {
	task, err := tu.taskRepo.GetTaskByID(id , role)

	if err != nil {
		return &domain.ErrorResponse{Message: "cannot get task", Status: 500}
	}

	if role == "admin" || task.USERID == id || role == "superAdmin" {
		return &domain.SingleTaskResponse{Single_Task: task}
	}
	return &domain.ErrorResponse{Message: "you are not authorized to view this task", Status: 401}

}

func (tu *TaskUseCase) UpdateTask(userID string , taskid string , task domain.Task) (interface{}) {	
	if task.Description == "" || task.DueDate == "" || task.Status == "" || task.Title == ""  {
		return &domain.ErrorResponse{Message: "missing required fields",  Status: 400}
	}

	userIDCheker := tu.taskRepo.UserIdGetter(userID , taskid)

	if !userIDCheker {
		updatedTask , err := tu.taskRepo.UpdateTask(taskid , task)
		
		if err != nil {
			return &domain.ErrorResponse{Message: "cannot update task", Status: 500}
		}
		return &domain.TaskSuccessResponse{Message: "task updated successfully", Status: 200, Data: updatedTask}
	}

	return &domain.ErrorResponse{Message: "you are not authorized to update this task", Status: 401}
}