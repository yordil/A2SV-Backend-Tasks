package repository_test

import (
	"task7/domain"
	"task7/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskRepositoryTestSuite defines the suite struct
type TaskRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.TaskRepository
}

// SetupTest initializes the test suite
func (suite *TaskRepositoryTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.TaskRepository)
}

// TestCreateTask tests the CreateTask function
func (suite *TaskRepositoryTestSuite) TestCreateTask() {
	task := domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}

	suite.mockRepo.On("CreateTask", task).Return(task, nil)

	result, err := suite.mockRepo.CreateTask(task)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), task.Title, result.Title)
	assert.Equal(suite.T(), task.Description, result.Description)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestGetAllTask tests the GetAllTask function
func (suite *TaskRepositoryTestSuite) TestGetAllTask() {
	tasks := []domain.Task{
		{Title: "Task 1", Description: "Description 1"},
		{Title: "Task 2", Description: "Description 2"},
	}

	suite.mockRepo.On("GetAllTask").Return(tasks, nil)

	result, err := suite.mockRepo.GetAllTask()
	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), tasks, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestGetTaskByUserID tests the GetTaskByUserID function
func (suite *TaskRepositoryTestSuite) TestGetTaskByUserID() {
	userID := primitive.NewObjectID().Hex()

	tasks := []domain.Task{
		{Title: "Task 1", Description: "Description 1", USERID: userID},
		{Title: "Task 2", Description: "Description 2", USERID: userID},
	}

	suite.mockRepo.On("GetTaskByUserID", userID).Return(tasks, nil)

	result, err := suite.mockRepo.GetTaskByUserID(userID)
	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), tasks, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestDeleteTask tests the DeleteTask function
func (suite *TaskRepositoryTestSuite) TestDeleteTask() {
	taskID := primitive.NewObjectID().Hex()

	suite.mockRepo.On("DeleteTask", taskID).Return(nil)

	err := suite.mockRepo.DeleteTask(taskID)
	assert.NoError(suite.T(), err)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestGetTaskByID tests the GetTaskByID function
func (suite *TaskRepositoryTestSuite) TestGetTaskByID() {
	taskID := primitive.NewObjectID().Hex()

	task := domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}

	suite.mockRepo.On("GetTaskByID", taskID, "user").Return(task, nil)

	result, err := suite.mockRepo.GetTaskByID(taskID, "user")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), task, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestUpdateTask tests the UpdateTask function
func (suite *TaskRepositoryTestSuite) TestUpdateTask() {
	taskID := primitive.NewObjectID().Hex()

	task := domain.Task{
		Title:       "Updated Task",
		Description: "This is an updated task description",
	}

	updatedTask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Updated Task",
		Description: "This is an updated task description",
	}

	suite.mockRepo.On("UpdateTask", taskID, task).Return(updatedTask, nil)

	result, err := suite.mockRepo.UpdateTask(taskID, task)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedTask, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestUserIdGetter tests the UserIdGetter function
func (suite *TaskRepositoryTestSuite) TestUserIdGetter() {
	userID := primitive.NewObjectID().Hex()
	taskID := primitive.NewObjectID().Hex()

	suite.mockRepo.On("UserIdGetter", userID, taskID).Return(true)

	result := suite.mockRepo.UserIdGetter(userID, taskID)
	assert.True(suite.T(), result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// Run the test suite
func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
