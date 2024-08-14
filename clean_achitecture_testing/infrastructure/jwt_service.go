package infrastructure

import (
	"task7/domain"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user domain.User) (string, error) {
	
	claims := &domain.JwtCustomClaims{
		User_id: user.ID,
		Email:   user.Email,
		Role:    user.Role,
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"claims": claims})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}	