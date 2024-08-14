package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"task7/Delivery/controllers"
	"task7/domain"
	"task7/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserController_SignUp(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	userController := controllers.NewUserController(mockUserUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUser := domain.User{
			Email:    "test@gmail.com",
			Password: "password",
		}

		mockUserUsecase.On("RegisterUser", mock.AnythingOfType("domain.User")).Return(domain.Response{Status: http.StatusOK}).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		jsonInput := `{"email": "test@gmail.com", "password": "password"}`
		c.Request, _ = http.NewRequest(http.MethodPost, "/signup", strings.NewReader(jsonInput))
		c.Request.Header.Set("Content-Type", "application/json")

		userController.SignUp(c)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"email": "test"}`))
		c.Request.Header.Set("Content-Type", "application/json")

		userController.SignUp(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUserController_Login(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	userController := controllers.NewUserController(mockUserUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUsecase.On("Login", "test@gmail.com", "password").Return(domain.Response{Status: http.StatusOK}).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		jsonInput := `{"email": "test@gmail.com", "password": "password"}`
		c.Request, _ = http.NewRequest(http.MethodPost, "/login", strings.NewReader(jsonInput))
		c.Request.Header.Set("Content-Type", "application/json")

		userController.Login(c)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email": "test"}`))
		c.Request.Header.Set("Content-Type", "application/json")

		userController.Login(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUserController_DeleteUser(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	userController := controllers.NewUserController(mockUserUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUsecase.On("DeleteUser", "123").Return(domain.Response{Status: http.StatusOK}).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: "123"}}
		c.Set("user_id", "123")
		c.Set("role", "admin")

		c.Request, _ = http.NewRequest(http.MethodDelete, "/users/123", nil)

		userController.DeleteUser(c)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("forbidden", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: "123"}}
		c.Set("user_id", "456")
		c.Set("role", "user")

		c.Request, _ = http.NewRequest(http.MethodDelete, "/users/123", nil)

		userController.DeleteUser(c)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}

func TestUserController_GetAllUsers(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	userController := controllers.NewUserController(mockUserUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUsecase.On("GetAllUsers").Return(domain.Response{Status: http.StatusOK}).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("role", "superAdmin")

		c.Request, _ = http.NewRequest(http.MethodGet, "/users", nil)

		userController.GetAllUsers(c)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("forbidden", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("role", "user")

		c.Request, _ = http.NewRequest(http.MethodGet, "/users", nil)

		userController.GetAllUsers(c)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}

func TestUserController_UpdateUser(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	userController := controllers.NewUserController(mockUserUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUser := domain.User{
			Email:    "test@gmail.com",
			Password: "password",
		}

		mockUserUsecase.On("UpdateUser", "123", mock.AnythingOfType("*domain.User"), "admin").Return(domain.Response{Status: http.StatusOK}).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("User_id", "123")
		c.Set("Role", "admin")

		jsonInput := `{"email": "test@gmail.com", "password": "password"}`
		c.Request, _ = http.NewRequest(http.MethodPut, "/users/123", strings.NewReader(jsonInput))
		c.Request.Header.Set("Content-Type", "application/json")

		userController.UpdateUser(c)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodPut, "/users/123", strings.NewReader(`{"email": "test"}`))
		c.Request.Header.Set("Content-Type", "application/json")

		userController.UpdateUser(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
