package usecase

import (
	"errors"
	"task7/domain"
)

type UserUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (uc *UserUsecase) RegisterUser(user domain.User) error {
	if user.Email == "" || user.Password == "" || user.Role == "" {
		return errors.New("missing required fields")
	}
	_, err := uc.UserRepo.CreateUser(user)
	return err
}
