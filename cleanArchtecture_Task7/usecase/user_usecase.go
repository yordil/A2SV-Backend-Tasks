package usecase

import (
	"errors"
	"task7/domain"
	"task7/infrastructure"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// hashing the passwordmf
	hashedpassword , err := infrastructure.HashPassword(user.Password)
	
	if err != nil { 
		return  err
	}

	user.Password = hashedpassword
	user.ID = primitive.NewObjectID()
	_, err = uc.UserRepo.CreateUser(user)

	return err
}


func (uc *UserUsecase) Login(email string, password string) (string, error) {


	user, err := uc.UserRepo.Login(email, password)
	if err != nil {
		return "", err
	}

	// comparing hashed and current
	match := infrastructure.CheckPassword(user.Password, password)
	if !match {
		return "", errors.New("invalid email or password")
	}

	user.Password = ""	
	// tokenization of user
	// _, err = infrastructure.GenerateToken()
	token , err := infrastructure.GenerateToken(user)
	if err != nil { 
		return "", err
	}
	
	return token, nil
}

func (uc *UserUsecase) DeleteUser(id string) error {
	uc.UserRepo.DeleteUser(id)
	return nil
}

func (uc *UserUsecase) GetAllUsers() ([]domain.User, error) {
	users, err := uc.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}