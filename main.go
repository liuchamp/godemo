package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.193:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.TODO())
	col := client.Database("casbin").Collection("casbin_rule")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cor, err := col.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	if cor.Next(ctx) {

	}
}
