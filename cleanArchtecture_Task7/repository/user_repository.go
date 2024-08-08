package repository

import (
	"errors"
	"task7/domain"
)

type UserRepositoryImpl struct {
	users map[string]domain.User
}

func NewUserRepositoryImpl() domain.UserRepository {
	return &UserRepositoryImpl{users: make(map[string]domain.User)}
}

func (repo *UserRepositoryImpl) CreateUser(user domain.User) (domain.User, error) {
	if _, exists := repo.users[user.Email]; exists {
		return domain.User{}, errors.New("user already exists")
	}
	repo.users[user.Email] = user
	return user, nil
}
