package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	USERID 		string 				`json:"user_id"`
	Title       string             `json:"title" `
	Description string             `json:"description" `
	DueDate     time.Time          `json:"due_date" `
	Status      string             `json:"status"`
}
