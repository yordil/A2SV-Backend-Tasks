package repository

import (
	"context"
	"task7/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositoryImpl struct {
	client  *mongo.Client
	dbname string
	collection string 
}

func NewTaskRepositoryImpl(db *mongo.Client , dbname string,  collection string) domain.TaskRepository {
	return &TaskRepositoryImpl{
		client: db,
		dbname: dbname,
		collection: collection, 
	}
}




// NewUserRepositoryImpl initializes the repository with a MongoDB collection
func (tr *TaskRepositoryImpl) CreateTask(task domain.Task) (domain.Task, error) {
	
	collection := tr.client.Database(tr.dbname).Collection(tr.collection)
	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return domain.Task{}, err 
	}

	return task, nil
}

func (tr * TaskRepositoryImpl) GetAllTask() ([]domain.Task, error) {
	var tasks []domain.Task
	collection := tr.client.Database(tr.dbname).Collection(tr.collection)
	cursor, err := collection.Find(context.Background(), bson.M{})
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


func (tr *TaskRepositoryImpl) GetTaskByUserID(id string) ([]domain.Task, error) {
	collection := tr.client.Database(tr.dbname).Collection(tr.collection)
	var tasks []domain.Task
	
	
	cursor, err := collection.Find(context.Background(), bson.M{"userid": id})
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

func (tr *TaskRepositoryImpl) DeleteTask(id string) (error) {
	collection := tr.client.Database(tr.dbname).Collection(tr.collection)
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection.DeleteOne(context.Background(), bson.M{"_id": idPrimitive})

	return nil

}


func (tr *TaskRepositoryImpl) GetTaskByID(id string, role string) (domain.Task, error) {
	var task domain.Task
	collection := tr.client.Database(tr.dbname).Collection(tr.collection)
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}
	err = collection.FindOne(context.Background(), bson.M{"_id": idPrimitive}).Decode(&task)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (tr *TaskRepositoryImpl) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	collection := tr.client.Database(tr.dbname).Collection(tr.collection)
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
	err = collection.FindOneAndUpdate(context.Background(), bson.M{"_id": idPrimitive}, update).Decode(&updatedTask)
	
	if err != nil {
		return domain.Task{}, err
	}

	return updatedTask, nil
}



func(tr * TaskRepositoryImpl) UserIdGetter(userid string , taskid string) bool {
	collection := tr.client.Database(tr.dbname).Collection(tr.collection)
    // get userid from the task collection
    var task domain.Task

    objID, err := primitive.ObjectIDFromHex(taskid)
    if err != nil {
        return false
    }
    // context
    err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
    if err != nil {
        return false
    }
	
    return task.USERID == userid 

}

