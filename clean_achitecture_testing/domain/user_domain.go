package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	DeleteUser(id string) 
	Login(email string, password string) (User, error)
	UserEmailGetter(email string  , user *User) (User , error)
	UpdateUser(id string  , user *User) (User , error)
	
}


type UserUsecase interface {
    RegisterUser(user  User) (interface{} )
    Login(email string, password string) (interface{} )
    DeleteUser(id string) (interface{} )
    GetAllUsers() (interface{})
	UpdateUser(id string , user *User , role string) (interface{} )
}


type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}

type SingleUserResponse struct {
	Single_User User
}

type AllUserResponse struct {
	All_User []User
}