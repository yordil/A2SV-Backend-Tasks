package usecase

import (
	"task7/domain"
	"task7/infrastructure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (uc *UserUsecase) RegisterUser(user domain.User) (interface{} ) {
	if user.Email == "" || user.Password == "" || user.Role == "" {
		return &domain.ErrorResponse{Message: "All fields are required", Status: 400}
	}

	hashedpassword , err := infrastructure.HashPassword(user.Password)
	if err != nil { 
		return  &domain.ErrorResponse{Message: "Error hashing password", Status: 500}
	}
	
	user.Password = hashedpassword
	user.ID = primitive.NewObjectID()
	_, err = uc.UserRepo.CreateUser(user)

	if err != nil {
		return &domain.ErrorResponse{Message: "User with this email Already Exists" , Status: 400}
	}


	return &domain.SuccessResponse{Message: "User created successfully", Status: 200}
}


func (uc *UserUsecase) Login(email string, password string)(interface{}) {


	user, err := uc.UserRepo.Login(email, password)
	if err != nil {
		return  &domain.ErrorResponse{Message: "invalid email or password", Status: 400}
	}

	// comparing hashed and current
	match := infrastructure.CheckPassword(user.Password, password)
	if !match {
		return &domain.ErrorResponse{Message: "invalid email or password", Status: 400}
	}
	user.Password = ""	
	// tokenization of user
	// _, err = infrastructure.GenerateToken()
	token , err := infrastructure.GenerateToken(user)
	if err != nil { 
		return &domain.ErrorResponse{Message: "Error generating token", Status: 500}
	}


	
	return &domain.LoginResponse{ Message: "Login successfully", Token: token}
}

func (uc *UserUsecase) DeleteUser(id string) (interface{} ) {
	uc.UserRepo.DeleteUser(id)
	return &domain.SuccessResponse{Message: "User deleted successfully", Status: 200}
}

func (uc *UserUsecase) GetAllUsers() (interface{} ) {
	users, err := uc.UserRepo.GetAllUsers()
	if err != nil {
		return &domain.ErrorResponse{Message: "Error getting users", Status: 500}
	}
	return &domain.AllUserResponse{All_User: users}
}

func (uc *UserUsecase) UpdateUser(id string , user *domain.User , role string) (interface{} ) {
	if user.Role != "" && role != "superAdmin" {
		return &domain.ErrorResponse{Message: "Only Super Admin Can Change a Role", Status: 403}
	}
	if user.Email == "" {
		return &domain.ErrorResponse{Message: "Email is Not editable", Status: 400}
	}

	updateduser , err := uc.UserRepo.UpdateUser(id  , user)

	if err != nil {
		return &domain.ErrorResponse{Message: "Cannot Update The user" , Status: 500}
	}

	return &domain.SingleUserResponse{Single_User: updateduser}


}