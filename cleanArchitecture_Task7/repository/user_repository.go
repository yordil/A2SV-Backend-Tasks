package repository

import (
	"context"
	"errors"
	"task7/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserRepositoryImpl struct {
	client  *mongo.Client
	dbname string
	collection string 
}

// NewUserRepositoryImpl initializes the repository with a MongoDB collection
func NewUserRepositoryImpl(db *mongo.Client , dbname string,  collection string) domain.UserRepository {
	return &UserRepositoryImpl{
		client: db,
		dbname: dbname,
		collection: collection, 
	}
}

func (repo *UserRepositoryImpl) CreateUser(user domain.User) (domain.User, error) {
	// Check if user already exists by email
	collection := repo.client.Database(repo.dbname).Collection(repo.collection)

	var existingUser domain.User

	err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return domain.User{}, errors.New("user with this email already exists")
	}
	if err != mongo.ErrNoDocuments {
		return domain.User{}, err
	}
	
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}


func (repo *UserRepositoryImpl) Login(email string, password string) (domain.User, error) {
	var user domain.User

	_ , err := repo.UserEmailGetter(email , &user)
	
	if err != nil { 
		return domain.User{} , err
	}

	return user , nil

}

func (repo *UserRepositoryImpl) DeleteUser(id string) {
	collection := repo.client.Database(repo.dbname).Collection(repo.collection)
	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	collection.DeleteOne(context.Background(), bson.M{"_id": idPrimitive})

}

func (repo *UserRepositoryImpl) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	collection := repo.client.Database(repo.dbname).Collection(repo.collection)
	cursor, err := collection.Find(context.Background(), bson.M{})
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


func (repo *UserRepositoryImpl) UserEmailGetter(email string  , user *domain.User) ( domain.User , error) {
	collection := repo.client.Database(repo.dbname).Collection(repo.collection)
	
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return domain.User{} , err
	}
	if err != nil {
		return domain.User{} , err
	}
	return *user , nil
}

func (repo *UserRepositoryImpl) UpdateUser(id string , user *domain.User) (domain.User , error){
	collection := repo.client.Database(repo.dbname).Collection(repo.collection)
	var updatedUser domain.User
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{} , err
	}
	 update := bson.M{
        "$set": bson.M{
            "password": user.Password,
        },
    }
	err = collection.FindOneAndUpdate(context.Background() ,bson.M{"_id": idPrimitive}, update).Decode(&updatedUser)

	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}


