package controllers

import (
	"net/http"
	"task7/domain"
	"task7/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    UserUsecase *usecase.UserUsecase
}

func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
    return &UserController{UserUsecase: userUsecase}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := uc.UserUsecase.RegisterUser(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token , err := uc.UserUsecase.Login(user.Email, user.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token , "message": "Login successfully"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    claimasID := c.GetString("user_id")
    claimsRole := c.GetString("role")
    
    if id != claimasID  && claimsRole != "admin"  {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this user"})
        return
    }
        
    if err := uc.UserUsecase.DeleteUser(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}


func (uc *UserController) GetAllUsers(c* gin.Context) {
    
    role := c.GetString("role")
    
    if role != "admin" { 
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to view this page"})
        return
    }
    users , err := uc.UserUsecase.GetAllUsers()
    
    if err != nil { 
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"users": users})
}