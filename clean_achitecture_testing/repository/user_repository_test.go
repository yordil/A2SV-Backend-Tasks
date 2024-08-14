package repository_test

import (
	"task7/domain"
	"task7/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define the test suite struct
type UserRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.UserRepository
}

// SetupTest initializes the test suite
func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.UserRepository)
}

// TestCreateUser tests the CreateUser function
func (suite *UserRepositoryTestSuite) TestCreateUser() {
	user := domain.User{
		Email:    "test@example.com",
		Password: "password123",
		Role:     "Admin",
	}

	suite.mockRepo.On("CreateUser", user).Return(user, nil)

	result, err := suite.mockRepo.CreateUser(user)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestDeleteUser tests the DeleteUser function
func (suite *UserRepositoryTestSuite) TestDeleteUser() {
	userID := primitive.NewObjectID().Hex()

	suite.mockRepo.On("DeleteUser", userID).Return(nil)

	suite.mockRepo.DeleteUser(userID)

	suite.mockRepo.AssertCalled(suite.T(), "DeleteUser", userID)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestGetAllUsers tests the GetAllUsers function
func (suite *UserRepositoryTestSuite) TestGetAllUsers() {
	users := []domain.User{
		{Email: "user1@example.com", Password: "password1", Role: "Admin"},
		{Email: "user2@example.com", Password: "password2", Role: "Admin"},
	}

	suite.mockRepo.On("GetAllUsers").Return(users, nil)

	result, err := suite.mockRepo.GetAllUsers()
	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), users, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestLogin tests the Login function
func (suite *UserRepositoryTestSuite) TestLogin() {
	email := "test@example.com"
	password := "password123"

	user := domain.User{
		Email:    email,
		Password: password,
	}

	suite.mockRepo.On("Login", email, password).Return(user, nil)

	result, err := suite.mockRepo.Login(email, password)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestUpdateUser tests the UpdateUser function
func (suite *UserRepositoryTestSuite) TestUpdateUser() {
	userID := primitive.NewObjectID().Hex()

	updatedUser := domain.User{
		Email:    "updated@example.com",
		Password: "newpassword123",
	}

	suite.mockRepo.On("UpdateUser", userID, &updatedUser).Return(updatedUser, nil)

	result, err := suite.mockRepo.UpdateUser(userID, &updatedUser)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedUser, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// TestUserEmailGetter tests the UserEmailGetter function
func (suite *UserRepositoryTestSuite) TestUserEmailGetter() {
	email := "test@example.com"

	user := domain.User{
		Email:    email,
		Password: "password123",
	}

	suite.mockRepo.On("UserEmailGetter", email, &user).Return(user, nil)

	result, err := suite.mockRepo.UserEmailGetter(email, &user)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

// Run the test suite
func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
