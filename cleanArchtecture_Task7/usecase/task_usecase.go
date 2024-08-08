package usecase

import (
	"errors"
	"task7/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{taskRepo: taskRepo}
}

func (tu *TaskUseCase) CreateTask(task domain.Task) (domain.Task, error) {
	if task.Description == "" || task.DueDate == "" || task.Status == "" || task.Title == "" || task.USERID == "" {
		return domain.Task{}, errors.New("missing required fields")
	}

	id := primitive.NewObjectID()
	task.ID = id
	task, err := tu.taskRepo.CreateTask(task)

	return task, err

}

func (tu *TaskUseCase) GetTasksByUserID(id string) ([]domain.Task, error) {
	
	tasks, err := tu.taskRepo.GetTaskByUserID(id)

	return tasks, err
}

func (tu *TaskUseCase) GetAllTask() ([]domain.Task, error) {
	tasks, err := tu.taskRepo.GetAllTask()
	return tasks, err
}


func (tu *TaskUseCase) DeleteTask(userID string , taskID string , role string) error {
	
	
	if role == "admin" || tu.taskRepo.UserIdGetter(userID , taskID){
		err := tu.taskRepo.DeleteTask(taskID)
		
		return err
		
	}

	return errors.New("not authorized to delete this task")
}	


func (tu *TaskUseCase) GetTaskByID(id string , role string) (domain.Task, error) {
	task, err := tu.taskRepo.GetTaskByID(id , role)

	if err != nil {
		return domain.Task{}, err
	}

	if role != "admin" && task.USERID != id {
		return domain.Task{}, errors.New("you are not authorized to view this task")
	
		
	}
	return task, nil

}

func (tu *TaskUseCase) UpdateTask(userID string , taskid string , task domain.Task) (domain.Task, error) {
	if task.Description == "" || task.DueDate == "" || task.Status == "" || task.Title == ""  {
		return domain.Task{}, errors.New("missing required fields")
	}

	userIDCheker := tu.taskRepo.UserIdGetter(userID , taskid)

	if !userIDCheker {
		updatedTask , err := tu.taskRepo.UpdateTask(taskid , task)
		
		if err != nil {
			return domain.Task{} , errors.New("cannot Update the task")
		}
		return updatedTask , nil
	}

	return domain.Task{} , errors.New("you are not authorized to edit this Page")
}