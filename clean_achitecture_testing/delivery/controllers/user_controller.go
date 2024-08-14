package controllers

import (
	"fmt"
	"net/http"
	"task7/domain"

	"github.com/gin-gonic/gin"
)


type UserController struct {
    UserUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
    return &UserController{UserUsecase: userUsecase}
}


func (uc *UserController) SignUp(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response := uc.UserUsecase.RegisterUser(user); 

    HandleResponse(c , response)
}



func (uc *UserController) Login(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response := uc.UserUsecase.Login(user.Email, user.Password)
    
    HandleResponse(c , response)
}



func (uc *UserController) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    claimasID := c.GetString("user_id")
    claimsRole := c.GetString("role")
    var response interface{}
    fmt.Println(claimasID)
    if claimsRole == "admin" ||  id == claimasID || claimsRole == "superAdmin"   {
        response = uc.UserUsecase.DeleteUser(id)
        HandleResponse(c , response)
        return 
    }
     
    c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this user"})
    
}



func (uc *UserController) GetAllUsers(c *gin.Context) {

    if c.GetString("role") != "superAdmin"  || c.GetString("role") != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to view this page"})
        return
    }
    response := uc.UserUsecase.GetAllUsers()
    
    HandleResponse(c , response)
}


func (uc *UserController) UpdateUser(c *gin.Context) {
    
    var user domain.User

    if err := c.ShouldBindJSON(&user); err != nil{
         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    role := c.GetString("Role")
    id := c.GetString("User_id")

    response := uc.UserUsecase.UpdateUser(id , &user , role)

   
    HandleResponse(c , response)

}