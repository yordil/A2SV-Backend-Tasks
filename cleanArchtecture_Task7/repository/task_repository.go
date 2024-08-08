package repository

import (
	"context"
	"fmt"
	"task7/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepositoryImpl struct {
	collection *mongo.Collection
}




func NewTaskRepositoryImpl(collection *mongo.Collection) domain.TaskRepository {
	return &taskRepositoryImpl{collection: collection}
}

func (tr *taskRepositoryImpl) CreateTask(task domain.Task) (domain.Task, error) {
	_, err := tr.collection.InsertOne(context.Background(), task)
	if err != nil {
		return domain.Task{}, err 
	}

	return task, nil
}

func (tr * taskRepositoryImpl) GetAllTask() ([]domain.Task, error) {
	var tasks []domain.Task
	cursor, err := tr.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task domain.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	return tasks, nil
}


func (tr *taskRepositoryImpl) GetTaskByUserID(id string) ([]domain.Task, error) {
	var tasks []domain.Task
	
	
	cursor, err := tr.collection.Find(context.Background(), bson.M{"userid": id})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task domain.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	return tasks, nil
	
}

func (tr *taskRepositoryImpl) DeleteTask(id string) (error) {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	tr.collection.DeleteOne(context.Background(), bson.M{"_id": idPrimitive})

	return nil

}


func (tr *taskRepositoryImpl) GetTaskByID(id string, role string) (domain.Task, error) {
	var task domain.Task
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}
	err = tr.collection.FindOne(context.Background(), bson.M{"_id": idPrimitive}).Decode(&task)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (tr *taskRepositoryImpl) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}

	 update := bson.M{
        "$set": bson.M{
            "title": task.Title,
            "description": task.Description,     
            "duedate": task.DueDate,     
            "status": task.Status,    
        },
    }
	var updatedTask domain.Task
	err = tr.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": idPrimitive}, update).Decode(&updatedTask)
	
	if err != nil {
		return domain.Task{}, err
	}

	return updatedTask, nil
}



func(tr * taskRepositoryImpl) UserIdGetter(userid string , taskid string) bool {

    // get userid from the task collection
    var task domain.Task

    objID, err := primitive.ObjectIDFromHex(taskid)
    if err != nil {
        return false
    }
    // context
    err = tr.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
    if err != nil {
        return false
    }
	fmt.Println(task.USERID , userid ," ************")
    return task.USERID == userid 

}

