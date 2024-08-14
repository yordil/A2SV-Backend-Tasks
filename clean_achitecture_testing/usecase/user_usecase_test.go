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

// UserUsecaseTestSuite defines the test suite structure
type UserUsecaseTestSuite struct {
	suite.Suite
	mockRepo    *mocks.UserRepository
	userUsecase domain.UserUsecase
}

// SetupTest sets up the test environment before each test
func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.UserRepository)
	suite.userUsecase = usecase.NewUserUsecase(suite.mockRepo)
}

// TestRegisterUser tests the RegisterUser function
func (suite *UserUsecaseTestSuite) TestRegisterUser() {
	suite.Run("success - valid user", func() {
		user := domain.User{
			Email:    "test@example.com",
			Password: "password123",
			Role:     "user",
		}

		suite.mockRepo.On("CreateUser", mock.Anything).Return(user, nil).Once()
		user.Password = "hashedPassword"

		result := suite.userUsecase.RegisterUser(user)

		assert.IsType(suite.T(), &domain.SuccessResponse{}, result)
		successResponse := result.(*domain.SuccessResponse)
		assert.Equal(suite.T(), "User created successfully", successResponse.Message)
		assert.Equal(suite.T(), 200, successResponse.Status)

		suite.mockRepo.AssertExpectations(suite.T())
	})

	suite.Run("failure - missing required fields", func() {
		user := domain.User{
			Email:    "",
			Password: "password123",
			Role:     "user",
		}

		result := suite.userUsecase.RegisterUser(user)

		assert.IsType(suite.T(), &domain.ErrorResponse{}, result)
		errorResponse := result.(*domain.ErrorResponse)
		assert.Equal(suite.T(), "All fields are required", errorResponse.Message)
		assert.Equal(suite.T(), 400, errorResponse.Status)
	})

	suite.Run("failure - user already exists", func() {
		user := domain.User{
			Email:    "test@example.com",
			Password: "password123",
			Role:     "user",
		}

		suite.mockRepo.On("CreateUser", mock.Anything).Return(domain.User{}, assert.AnError).Once()

		result := suite.userUsecase.RegisterUser(user)

		assert.IsType(suite.T(), &domain.ErrorResponse{}, result)
		errorResponse := result.(*domain.ErrorResponse)
		assert.Equal(suite.T(), "User with this email Already Exists", errorResponse.Message)
		assert.Equal(suite.T(), 400, errorResponse.Status)

		suite.mockRepo.AssertExpectations(suite.T())
	})
}

// TestDeleteUser tests the DeleteUser function
func (suite *UserUsecaseTestSuite) TestDeleteUser() {
	suite.Run("success - user deleted", func() {
		userID := primitive.NewObjectID().Hex()

		suite.mockRepo.On("DeleteUser", userID).Once()

		result := suite.userUsecase.DeleteUser(userID)

		assert.IsType(suite.T(), &domain.SuccessResponse{}, result)
		successResponse := result.(*domain.SuccessResponse)
		assert.Equal(suite.T(), "User deleted successfully", successResponse.Message)

		suite.mockRepo.AssertExpectations(suite.T())
	})
}

// TestGetAllUsers tests the GetAllUsers function
func (suite *UserUsecaseTestSuite) TestGetAllUsers() {
	suite.Run("success - users retrieved", func() {
		users := []domain.User{
			{ID: primitive.NewObjectID(), Email: "user1@example.com", Password: "hashedPassword1", Role: "user"},
			{ID: primitive.NewObjectID(), Email: "user2@example.com", Password: "hashedPassword2", Role: "admin"},
		}

		suite.mockRepo.On("GetAllUsers").Return(users, nil).Once()

		result := suite.userUsecase.GetAllUsers()

		assert.IsType(suite.T(), &domain.AllUserResponse{}, result)
		allUserResponse := result.(*domain.AllUserResponse)
		assert.Len(suite.T(), allUserResponse.All_User, 2)
		assert.Equal(suite.T(), users, allUserResponse.All_User)

		suite.mockRepo.AssertExpectations(suite.T())
	})

	suite.Run("error - unable to retrieve users", func() {
		suite.mockRepo.On("GetAllUsers").Return(nil, assert.AnError).Once()

		result := suite.userUsecase.GetAllUsers()

		assert.IsType(suite.T(), &domain.ErrorResponse{}, result)
		errorResponse := result.(*domain.ErrorResponse)
		assert.Equal(suite.T(), "Error getting users", errorResponse.Message)

		suite.mockRepo.AssertExpectations(suite.T())
	})
}

// TestUpdateUser tests the UpdateUser function
func (suite *UserUsecaseTestSuite) TestUpdateUser() {
	suite.Run("success - user updated", func() {
		userID := primitive.NewObjectID().Hex()
		updatedUser := domain.User{
			ID:       primitive.NewObjectID(),
			Email:    "updated@example.com",
			Password: "newHashedPassword",
			Role:     "user",
		}

		suite.mockRepo.On("UpdateUser", userID, &updatedUser).Return(updatedUser, nil).Once()

		result := suite.userUsecase.UpdateUser(userID, &updatedUser, "superAdmin")

		assert.IsType(suite.T(), &domain.SingleUserResponse{}, result)
		singleUserResponse := result.(*domain.SingleUserResponse)
		assert.Equal(suite.T(), updatedUser, singleUserResponse.Single_User)

		suite.mockRepo.AssertExpectations(suite.T())
	})

	suite.Run("error - unauthorized role change", func() {
		userID := primitive.NewObjectID().Hex()
		user := &domain.User{
			Email: "test@example.com",
			Role:  "user",
		}

		result := suite.userUsecase.UpdateUser(userID, user, "admin")

		assert.IsType(suite.T(), &domain.ErrorResponse{}, result)
		errorResponse := result.(*domain.ErrorResponse)
		assert.Equal(suite.T(), "Only Super Admin Can Change a Role", errorResponse.Message)
	})

	suite.Run("error - email not editable", func() {
		userID := primitive.NewObjectID().Hex()
		user := &domain.User{
			Email: "",
		}

		result := suite.userUsecase.UpdateUser(userID, user, "superAdmin")

		assert.IsType(suite.T(), &domain.ErrorResponse{}, result)
		errorResponse := result.(*domain.ErrorResponse)
		assert.Equal(suite.T(), "Email is Not editable", errorResponse.Message)
	})

	suite.Run("error - cannot update user", func() {
		userID := primitive.NewObjectID().Hex()
		user := &domain.User{
			Email: "test@example.com",
			Role:  "user",
		}

		suite.mockRepo.On("UpdateUser", userID, user).Return(domain.User{}, assert.AnError).Once()

		result := suite.userUsecase.UpdateUser(userID, user, "superAdmin")

		assert.IsType(suite.T(), &domain.ErrorResponse{}, result)
		errorResponse := result.(*domain.ErrorResponse)
		assert.Equal(suite.T(), "Cannot Update The user", errorResponse.Message)

		suite.mockRepo.AssertExpectations(suite.T())
	})
}

// Run the test suite
func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
