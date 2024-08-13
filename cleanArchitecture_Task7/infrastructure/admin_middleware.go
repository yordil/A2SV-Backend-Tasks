package infrastructure

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		fmt.Println(role , "this is admin middleware")
		if role == "admin" || role == "superAdmin" {
			c.Next()
			
			return
		}
		c.JSON(403, gin.H{"error": "You are not authorized to view this page"})
		c.Abort()
		
	}
}