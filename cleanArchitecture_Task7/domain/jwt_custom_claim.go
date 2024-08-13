package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtCustomClaims struct {
	User_id primitive.ObjectID `json:"user_id"`
	Email   string `json:"email"`
	Role string `json:"role"`
	// jwt.StandardClaims
}
