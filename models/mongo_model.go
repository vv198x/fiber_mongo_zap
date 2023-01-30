package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mongo struct {
	ID   primitive.ObjectID `bson:"_id"`
	Decs []PurchaseOrder    `json:"decs,omitempty"`
	Lock bool               `json:"lock"`
}
