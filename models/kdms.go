package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dem struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	User
}
