package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task7/delivery/controllers"
	"task7/domain"
	"task7/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask(t *testing.T) {
	// Set up Gin and the mock usecase
	gin.SetMode(gin.TestMode)
	mockTaskUsecase := new(mocks.TaskUsecase)
	router := gin.Default()
	tc := controllers.NewTaskController(mockTaskUsecase)
	router.POST("/tasks", tc.CreateTask)

	// Define input and expected output
	task := domain.Task{
		Title: "Sample Task",
	}
	taskJSON, _ := json.Marshal(task)

	mockTaskUsecase.On("CreateTask", mock.AnythingOfType("domain.Task")).Return(&domain.TaskSuccessResponse{Message: "Task created successfully"}, nil)

	// Test case 1: Successful creation of a task
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user_id", "123")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskUsecase.AssertExpectations(t)

	// Test case 2: Invalid JSON input
	invalidTaskJSON := []byte(`{"invalid":}`)
	req, err = http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(invalidTaskJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user_id", "123")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTasksByUserID(t *testing.T) {
	// Set up Gin and the mock usecase
	gin.SetMode(gin.TestMode)
	mockTaskUsecase := new(mocks.TaskUsecase)
	router := gin.Default()
	tc := controllers.NewTaskController(mockTaskUsecase)
	router.GET("/tasks/user", tc.GetTasksByUserID)
	id := primitive.NewObjectID()

	strid := id.Hex()
	// Define expected output
	tasks := []domain.Task{
		{ Title: "Task 1", USERID: strid},
		{Title: "Task 2", USERID: strid},
	}
	mockTaskUsecase.On("GetTasksByUserID", strid).Return(&domain.TaskResponse{All_Task: tasks}, nil)

	// Test case: Get tasks by user ID
	req, err := http.NewRequest(http.MethodGet, "/tasks/user", nil)
	assert.NoError(t, err)
	req.Header.Set("User_id", strid)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	mockTaskUsecase.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	// Set up Gin and the mock usecase
	gin.SetMode(gin.TestMode)
	mockTaskUsecase := new(mocks.TaskUsecase)
	router := gin.Default()
	tc := controllers.NewTaskController(mockTaskUsecase)
	router.DELETE("/tasks/:id", tc.DeleteTask)

	// Define input and expected output
	mockTaskUsecase.On("DeleteTask", "123", "1", "user").Return(&domain.SuccessResponse{Message: "Task deleted successfully"}, nil)

	// Test case: Successful deletion of a task
	req, err := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	assert.NoError(t, err)
	req.Header.Set("user_id", "123")
	req.Header.Set("role", "user")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskUsecase.AssertExpectations(t)
}

