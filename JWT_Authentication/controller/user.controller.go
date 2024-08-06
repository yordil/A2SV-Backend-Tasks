package controller

import (
	"auth/data"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)




func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := data.Register(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest , gin.H{"err" : err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK , gin.H {"message" : "User created successfully"})
	
}

func Login(c *gin.Context) {
	var user models.User
	var existingUser *models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	existingUser = data.Login(&user)
	if existingUser == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}

	// Verify the password
	err := utils.PasswordVerify([]byte(existingUser.Password), user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Incorrect password"})
		return
	}

	// Generate JWT token
	jwtToken, err := utils.TokenGenerator(existingUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "cannot create a token"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": jwtToken})
}



func GetAllUser(c * gin.Context) {

	users , _ := data.GetAllUser()

	if users != nil {
		if len(users) == 0 {
			c.IndentedJSON(http.StatusOK , gin.H{
				"message" : "No Uses found",
				"user" : users,
			})

			return 
		}

		c.IndentedJSON(http.StatusOK , gin.H{"users" : users})
	}

}

func UpdateUser(c * gin.Context) {
	var updatedUser *models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest , gin.H{"message" : err.Error()})
	}

	user , err := data.UpdateUser(id , updatedUser)

	if err != nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{
			"message"  :  "cannot update The user data",
		})
	}

	c.IndentedJSON(http.StatusBadRequest , gin.H{
		"message" :  "updated Successfully", 
		"user" : user,
	})

}
 

func DeleteUser(c *gin.Context){
	id := c.Param("id")

	err := data.DeleteUser(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest , gin.H{"message" : "User not found"})
	}

	c.IndentedJSON(http.StatusOK , gin.H{"message" : "User deleted successfully"})
	
}


