package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Test struct {
	Name string
	Oks  string
	Smod Sim
	Next *Test `json:"next,omitempty"`
}

type Sim struct {
	Name string
	Mod  string
}

type Menu struct {
	Id       primitive.ObjectID `bson:"id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Status   int                `bson:"status" json:"status"`
	Order    int                `bson:"order" json:"order"`
	Icon     string             `bson:"icon" json:"icon"`
	ParentId primitive.ObjectID `bson:"parent_id" json:"parentId"`
}
