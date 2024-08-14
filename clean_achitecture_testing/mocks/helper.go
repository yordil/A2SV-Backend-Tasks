package mocks

import (
	"task7/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetNewUser returns a mock instance of domain.User.
func GetNewUser() *domain.User {
	return &domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "mockemail@example.com",
		Password: "mockPassword",
		Role:     "mockRole",
	}
}

// GetManyUsers returns a slice of mock instances of domain.User.
func GetManyUsers() []domain.User {
	return []domain.User{
		{
			ID:       primitive.NewObjectID(),
			Email:    "mockemail1@example.com",
			Password: "mockPassword1",
			Role:     "mockRole1",
		},
		{
			ID:       primitive.NewObjectID(),
			Email:    "mockemail2@example.com",
			Password: "mockPassword2",
			Role:     "mockRole2",
		},
	}
}
