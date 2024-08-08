package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	// GetUserByEmail(email string) (User, error)
	GetAllUsers() ([]User, error)
	DeleteUser(id string) 
	// UpdateUser(user User) (User, error)
	Login(email string, password string) (User, error)
	
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}
