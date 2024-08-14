package usecase_test

import (
	"task7/domain"
	"task7/mocks"
	"task7/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCaseTestSuite struct {
    suite.Suite
    mockRepo    *mocks.TaskRepository
    taskUsecase domain.TaskUsecase
}

func (suite *TaskUseCaseTestSuite) SetupTest() {
    suite.mockRepo = new(mocks.TaskRepository)
    suite.taskUsecase = usecase.NewTaskUseCase(suite.mockRepo)
}

func (suite *TaskUseCaseTestSuite) TestCreateTask_Success() {
    task := domain.Task{
        Description: "Task description",
        DueDate:     "2024-08-15",
        Status:      "pending",
        Title:       "Task title",
        USERID:      "user123",
    }

    suite.mockRepo.On("CreateTask", mock.Anything).Return(task, nil).Once()

    result := suite.taskUsecase.CreateTask(task)

    suite.IsType(&domain.TaskSuccessResponse{}, result)
    successResponse := result.(*domain.TaskSuccessResponse)
    suite.Equal("task created successfully", successResponse.Message)
    suite.Equal(200, successResponse.Status)

    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestCreateTask_Failure_MissingFields() {
    task := domain.Task{
        Description: "",
        DueDate:     "2024-08-15",
        Status:      "pending",
        Title:       "Task title",
        USERID:      "user123",
    }

    result := suite.taskUsecase.CreateTask(task)

    suite.IsType(&domain.ErrorResponse{}, result)
    errorResponse := result.(*domain.ErrorResponse)
    suite.Equal("missing required fields", errorResponse.Message)
    suite.Equal(400, errorResponse.Status)
}

func (suite *TaskUseCaseTestSuite) TestGetTasksByUserID_Success() {
    userID := "user123"
    tasks := []domain.Task{
        {ID: primitive.NewObjectID(), Description: "Task 1", DueDate: "2024-08-15", Status: "pending", Title: "Task 1", USERID: userID},
        {ID: primitive.NewObjectID(), Description: "Task 2", DueDate: "2024-08-16", Status: "completed", Title: "Task 2", USERID: userID},
    }

    suite.mockRepo.On("GetTaskByUserID", userID).Return(tasks, nil).Once()

    result := suite.taskUsecase.GetTasksByUserID(userID)

    suite.IsType(&domain.TaskResponse{}, result)
    taskResponse := result.(*domain.TaskResponse)
    suite.Len(taskResponse.All_Task, 2)
    suite.Equal(tasks, taskResponse.All_Task)
}

func (suite *TaskUseCaseTestSuite) TestGetTasksByUserID_Failure() {
    userID := "user123"

    suite.mockRepo.On("GetTaskByUserID", userID).Return(nil, assert.AnError).Once()

    result := suite.taskUsecase.GetTasksByUserID(userID)

    suite.IsType(&domain.ErrorResponse{}, result)
    errorResponse := result.(*domain.ErrorResponse)
    suite.Equal("cannot get tasks", errorResponse.Message)
    suite.Equal(500, errorResponse.Status)
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask_Success() {
    userID := "user123"
    taskID := primitive.NewObjectID().Hex()
    role := "admin"

    suite.mockRepo.On("DeleteTask", taskID).Return(nil).Once()
    suite.mockRepo.On("UserIdGetter", userID, taskID).Return(true).Once()

    result := suite.taskUsecase.DeleteTask(userID, taskID, role)

    suite.IsType(&domain.TaskSuccessResponse{}, result)
    successResponse := result.(*domain.TaskSuccessResponse)
    suite.Equal("task deleted successfully", successResponse.Message)
    suite.Equal(200, successResponse.Status)
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask_Unauthorized() {
    userID := "user123"
    taskID := primitive.NewObjectID().Hex()
    role := "user"

    suite.mockRepo.On("UserIdGetter", userID, taskID).Return(false).Once()

    result := suite.taskUsecase.DeleteTask(userID, taskID, role)

    suite.IsType(&domain.ErrorResponse{}, result)
    errorResponse := result.(*domain.ErrorResponse)
    suite.Equal("you are not authorized to delete this task", errorResponse.Message)
    suite.Equal(401, errorResponse.Status)
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask_Success() {
    userID := "user123"
    taskID := primitive.NewObjectID().Hex()
    updatedTask := domain.Task{
        Description: "Updated description",
        DueDate:     "2024-08-20",
        Status:      "completed",
        Title:       "Updated title",
    }

    suite.mockRepo.On("UserIdGetter", userID, taskID).Return(false).Once()
    suite.mockRepo.On("UpdateTask", taskID, updatedTask).Return(updatedTask, nil).Once()

    result := suite.taskUsecase.UpdateTask(userID, taskID, updatedTask)

    suite.IsType(&domain.TaskSuccessResponse{}, result)
    successResponse := result.(*domain.TaskSuccessResponse)
    suite.Equal("task updated successfully", successResponse.Message)
    suite.Equal(200, successResponse.Status)
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask_Unauthorized() {
    userID := "user123"
    taskID := primitive.NewObjectID().Hex()
    task := domain.Task{
        Description: "Task description",
        DueDate:     "2024-08-15",
        Status:      "pending",
        Title:       "Task title",
    }

    suite.mockRepo.On("UserIdGetter", userID, taskID).Return(true).Once()

    result := suite.taskUsecase.UpdateTask(userID, taskID, task)

    suite.IsType(&domain.ErrorResponse{}, result)
    errorResponse := result.(*domain.ErrorResponse)
    suite.Equal("you are not authorized to update this task", errorResponse.Message)
    suite.Equal(401, errorResponse.Status)
}

func (suite *TaskUseCaseTestSuite) TestGetTaskByID_Success() {
    userID := "user123"
    taskID := primitive.NewObjectID().Hex()
    role := "admin"
    task := domain.Task{
        ID:          primitive.NewObjectID(),
        Description: "Task description",
        DueDate:     "2024-08-15",
        Status:      "pending",
        Title:       "Task title",
        USERID:      userID,
    }

    suite.mockRepo.On("GetTaskByID", taskID, role).Return(task, nil).Once()

    result := suite.taskUsecase.GetTaskByID(taskID, role)

    suite.IsType(&domain.SingleTaskResponse{}, result)
    singleTaskResponse := result.(*domain.SingleTaskResponse)
    suite.Equal(task, singleTaskResponse.Single_Task)
}

func (suite *TaskUseCaseTestSuite) TestGetTaskByID_Unauthorized() {
    // userID := "user123"
    taskID := primitive.NewObjectID().Hex()
    role := "user"
    task := domain.Task{
        ID:          primitive.NewObjectID(),
        Description: "Task description",
        DueDate:     "2024-08-15",
        Status:      "pending",
        Title:       "Task title",
        USERID:      "anotherUser",
    }

    suite.mockRepo.On("GetTaskByID", taskID, role).Return(task, nil).Once()

    result := suite.taskUsecase.GetTaskByID(taskID, role)

    suite.IsType(&domain.ErrorResponse{}, result)
    errorResponse := result.(*domain.ErrorResponse)
    suite.Equal("you are not authorized to view this task", errorResponse.Message)
    suite.Equal(401, errorResponse.Status)
}

func TestTaskUseCaseTestSuite(t *testing.T) {
    suite.Run(t, new(TaskUseCaseTestSuite))
}
