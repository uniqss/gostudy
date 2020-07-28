package main

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.87.210:27017"))
	if err != nil {
		fmt.Println("mongo.Connect err:", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("client.Ping err:", err)
	}

	collection := client.Database("testing").Collection("numbers")

	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	ctx = context.Background()

	oid, err := primitive.ObjectIDFromHex("5f1fe8de8e435fc23b7b4e25")
	if err != nil {
		fmt.Println("primitive.ObjectIDFromHex err:", err)
	}
	filter := bson.M{"_id": oid}
	//filter := bson.M{"_id": "5f1fc525987d0b7bd33c1c90"}
	//filter := bson.M{"uid": 16}
	opt := options.FindOptions{}

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, filter, &opt)
	if err != nil {
		fmt.Println("collection.Find err:", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		fmt.Println("printing result:")
		fmt.Println(result)
		fmt.Println("_id:", (result["_id"].((primitive.ObjectID))).String())
		fmt.Println("_id:", (result["_id"].((primitive.ObjectID))).Hex())
		oid := bson.ObjectId("asdf")
		oid.String()
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
