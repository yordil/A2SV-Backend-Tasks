package infrastructure

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("secret")

func AuthMiddleware() gin.HandlerFunc {

  return func(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

if authHeader == "" {
  c.JSON(401, gin.H{"error": "Authorization header is required"})
  c.Abort()
  return

}

authParts := strings.Split(authHeader, " ")

if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
  c.JSON(401, gin.H{"error": "Invalid authorization header"})
  c.Abort()
  return

}

token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
  if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
    return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
  }
  
  return jwtSecret, nil
})


if err != nil || !token.Valid {
  c.JSON(401, gin.H{"error": "Invalid JWT" , "err" : err.Error()})
  c.Abort()
  return
}


if claims, ok := token.Claims.(jwt.MapClaims); ok {
	c.Set("user_id" , claims["user_id"])  
    c.Set("role" , claims["role"])		
   
}
    c.Next()
  }
}

