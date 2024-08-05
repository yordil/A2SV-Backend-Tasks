package data

import (
	"auth/config"
	"auth/models"
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var user_collection *mongo.Collection

func init() {
	client, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	var dbname = os.Getenv("DB_NAME")
	var user_collection_name = os.Getenv("user")

	user_collection = client.Database(dbname).Collection(user_collection_name)
}

func Register(user *models.User) error  {

	hashedPassword , err :=  bcrypt.GenerateFromPassword([]byte(user.Password) , bcrypt.DefaultCost)

	if err != nil {
		return err
	}

    email , _ := GetUserByEmail(user.Email)

	if email != nil { 
		return errors.New("users with this email already exists")
	}

	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    user.ID = primitive.NewObjectID()
    _ , err = user_collection.InsertOne(ctx, user)
    
	if err != nil {
        return err
    }

    return nil
}

func Login(user *models.User) (*models.User) {

	email  := user.Email

	var existing *models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := user_collection.FindOne(ctx, bson.M{"email": email}).Decode(&existing)

	if err != nil  || bcrypt.CompareHashAndPassword([]byte(existing.Password) , []byte(user.Password)) != nil {
		return nil
	}

	return existing 
}

func GetAllUser() ([]*models.User , error) {
	
	var allUsers []*models.User
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := user_collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        allUsers = append(allUsers, &user)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return allUsers, nil

}

func UpdateUser(id string , user *models.User)  (*models.User , error){
	
	objID , err := primitive.ObjectIDFromHex(id)
    if err != nil { 
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    update := bson.M{
        "$set": bson.M{
            "password": user.Password,
            "role": user.Role,     
        },
    }

    if user.ID != primitive.NilObjectID {
       return nil, errors.New("cannot update task ID")
    }
    
    var updatedUser models.User
    err = collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update).Decode(&updatedUser)
    if err != nil {  
        return nil, err
    }

    return &updatedUser, nil
}

func DeleteUser(id string) error {	

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    
	defer cancel()

    objID, err := primitive.ObjectIDFromHex(id)

    if err != nil { 
        return err
    }
    _, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        return  err
    }

    return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := user_collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}