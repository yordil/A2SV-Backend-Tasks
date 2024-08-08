package infrastructure

import (
	"task7/domain"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user domain.User) (string, error) {
	
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"user_id": user.ID,
	"email":   user.Email,
	"role":   user.Role,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err

	}

	return tokenString, nil

}	