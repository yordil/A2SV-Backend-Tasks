package services

import (
	"context"
	"errors"
	"log"
	"time"

	"taskmanager/config"
	"taskmanager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// Init initializes the database connection and assigns the collection
func init() {
    client , err := config.ConnectDB()
    if err != nil { 
        log.Fatal(err)
    }
    collection = client.Database("taskmanager").Collection("tasks")
}

// GetTasks fetches all tasks from the database and returns them as a slice
func GetTasks() ([]*models.Task, error) {
    if collection == nil {
        return nil, errors.New("collection is not initialized")
    }

    var tasks []*models.Task
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var task models.Task
        if err := cursor.Decode(&task); err != nil {
            return nil, err
        }
        tasks = append(tasks, &task)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return tasks, nil
}

func CreateTask(task *models.Task) error {
    if collection == nil {
        return errors.New("collection is not initialized")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    task.ID = primitive.NewObjectID()
    _, err := collection.InsertOne(ctx, task)
    if err != nil {
        return err
    }

    return nil
}

func GetTaskByID(id string) (*models.Task, error) { 
    if collection == nil {
        return nil, errors.New("collection is not initialized")
    }

    var task models.Task
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
    if err != nil {
        return nil, err
    }

    return &task, nil


}

func DeleteTask(id string) (string , error) {
    if collection == nil {
        log.Fatal("collection is not initialized")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(id)

    if err != nil { 
        return "Invalid ID", err
    }
    _, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        return "Task Not Found", err
    }

    return "Task Deleted", nil
}



func UpdateTask(id string , task *models.Task) (*models.Task, error) {
    if collection == nil {
        return nil, errors.New("collection is not initialized")
    }
    objID , err := primitive.ObjectIDFromHex(id)
    if err != nil { 
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    update := bson.M{
        "$set": bson.M{
            "title":       task.Title,
            "description": task.Description,
            "status":      task.Status,
            // Add other fields here as needed
        },
    }

    if task.ID != primitive.NilObjectID {
       return nil, errors.New("cannot update task ID")
    }
    
    var updatedTask models.Task
    err = collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update).Decode(&updatedTask)
    if err != nil {  
        return nil, err
    }

    return &updatedTask, nil
}