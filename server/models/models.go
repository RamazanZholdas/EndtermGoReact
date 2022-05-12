package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoList struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Task   string             `bson:"task" json:"task,omitempty"`
	Status string             `bson:"status" json:"status,omitempty"`
}
