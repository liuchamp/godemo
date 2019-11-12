package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.193:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.TODO())

	col := client.Database("test").Collection("sdmo")
	indexs := col.Indexes()

	ctx := context.TODO()
	i := mongo.IndexModel{
		Keys: bson.M{"sd": 1},
	}

	// 创建唯一索引，需要保证记录中，有对应的字段，否则不能创建，重复创建时，不会报错
	i.Options = options.Index().SetUnique(true)
	s, er := indexs.CreateOne(ctx, i, options.CreateIndexes())
	if er != nil {
		log.Fatal(err)
	} else {
		log.Println(s)
	}
	cor, err := indexs.List(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer cor.Close(ctx)

	for cor.Next(ctx) {
		im := new(mongo.IndexModel)
		if e := cor.Decode(im); e != nil {
			log.Println(e)
			return
		} else {
			log.Println(im)
		}

	}
}
