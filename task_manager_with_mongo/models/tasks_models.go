package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Title       string             `json:"title" `
	Description string             `json:"description" `
	// DueDate     time.Time          `json:"due_date" `
	Status      string             `json:"status"`
}
