package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
)

type Smou struct {
	SID  string `json:"_id" bson:"_id"`
	Scd  string `json:"scd" bson:"scd"`
	Smut string `json:"sd" bson:"sd"`
}
type TestDemo struct {
	ID  primitive.ObjectID `json:"_id" bson:"_id"`
	Sid string             `json:"sd" bson:"sd"`
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.193:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.TODO())

	client.Database("test").Collection("sdmo").DeleteMany(context.TODO(), bson.M{})
	sduls()
	inster(client)
}

func sduls() {
	input, err := ioutil.ReadFile("sdi.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var ss []TestDemo
	err = json.Unmarshal([]byte(input), &ss)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bs, _ := json.Marshal(ss)
	fmt.Println(string(bs))
}
func inster(client *mongo.Client) {
	input, err := ioutil.ReadFile("input.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var ss []Smou
	err = json.Unmarshal([]byte(input), &ss)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bs, err := json.Marshal(ss)
	fmt.Println(string(bs))
	var doc []interface{}
	for _, e := range ss {
		doc = append(doc, e)
	}

	sr, err := client.Database("test").Collection("sdmo").InsertMany(context.TODO(), doc)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range sr.InsertedIDs {
		fmt.Println(v)
	}
}
