package repository

import (
	"context"
	"errors"
	"fmt"
	"task7/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserRepositoryImpl struct {
	collection *mongo.Collection
}

// NewUserRepositoryImpl initializes the repository with a MongoDB collection
func NewUserRepositoryImpl(collection *mongo.Collection) domain.UserRepository {
	return &UserRepositoryImpl{collection: collection}
}

// CreateUser inserts a new user into the MongoDB collection
func (repo *UserRepositoryImpl) CreateUser(user domain.User) (domain.User, error) {
	// Check if user already exists by email
	var existingUser domain.User
	err := repo.collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return domain.User{}, errors.New("user with this email already exists")
	}
	if err != mongo.ErrNoDocuments {
		return domain.User{}, err
	}

	
	fmt.Println(user , "***************************************")
	_, err = repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}


func (repo *UserRepositoryImpl) Login(email string, password string) (domain.User, error) {
	var user domain.User
	err := repo.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return domain.User{}, errors.New("user not found")
	}
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}

func (repo *UserRepositoryImpl) DeleteUser(id string) {
	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	repo.collection.DeleteOne(context.Background(), bson.M{"_id": idPrimitive})

}

func (repo *UserRepositoryImpl) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	cursor, err := repo.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user domain.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil

}