package controller

import (
	"auth/data"
	"auth/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("authsecret")

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
		c.JSON(http.StatusBadRequest , gin.H{"message" : err.Error(),
		})
	}

	existingUser = data.Login(&user)	
	if existingUser == nil {	 
		c.IndentedJSON(http.StatusBadRequest , gin.H{"message" : "User not found",})
		return 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"user_id": existingUser.ID,
	"email":   existingUser.Email,
	"role":   existingUser.Role,
	})

	jwtToken, err := token.SignedString(jwtSecret)

	if err != nil { 
		c.IndentedJSON(http.StatusInternalServerError , gin.H{"message" :	"cannot create a token",})
	}

	c.IndentedJSON(http.StatusOK , gin.H{"message" : "User logged in successfully ", "token": jwtToken})

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

	userid := c.GetString("user_id")

	if id != userid { 
		c.IndentedJSON(http.StatusBadRequest , gin.H{"message" : "Not authorized to update this user"})
	}

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




