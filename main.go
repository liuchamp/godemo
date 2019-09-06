package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"windplatform/webbackend/webserver/util"
)

type DemoDec struct {
	UpdateTime time.Time
	CreateTime time.Time
}

const timeform = "2006-01-02 15:04:05"

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.193:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.TODO())
	col := client.Database("test").Collection("tezt")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ddc := DemoDec{}
	ddc.CreateTime = time.Now()
	ddc.UpdateTime = time.Now()

	_, err = col.InsertOne(ctx, ddc)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"createtime": bson.M{"$lt": time.Now().Add(-2 * time.Hour), "$gt": time.Now().Add(-4 * time.Hour)}}
	opt := options.FindOptions{}
	var lim int64 = 101
	var sikp int64 = 0
	opt.Limit = &lim
	opt.Skip = &sikp
	cor, err := col.Find(ctx, filter, &opt)
	if err != nil {
		panic(err)
	}
	defer cor.Close(ctx)

	for cor.Next(ctx) {
		sm := new(DemoDec)
		if err := cor.Decode(sm); err != nil {
			panic(err)
		}
		fmt.Println(util.TimeFormat(sm.CreateTime), "------>", util.TimeFormat(sm.UpdateTime))
	}

}
