package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.87.210:27017"))
	if err != nil {
		fmt.Println("mongo.Connect err:", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Error("client.Ping err:", err)
	}

	collection := client.Database("testing").Collection("numbers")

	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	ctx = context.Background()
	//res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	//id := res.InsertedID
	//fmt.Println(id)

	userId := 50000000
	infoJson := "{\"uinfo\":{\"level\":123, \"name\":\"肉山大王\", \"score\":100}, \"Level1\":10000101, \"Level2\":10000101, \"Level3\":10000101}"
	//infoJson2 := "{\"uinfo\":{\"level\":234, \"name\":\"肉山大王\", \"score\":345}}"

	// insert
	_, err = collection.InsertOne(ctx, bson.M{"uid": userId, "info_json": infoJson, "quest1":123})
	if err != nil {
		log.Error("collection.InsertOne err:", err)
	}

	var result bson.M

	// find
	fOne := collection.FindOne(ctx, bson.M{"uid": userId}, &options.FindOneOptions{})
	if fOne.Err() != nil {
		log.Warn("collection.FindOne fOne.Err():", fOne.Err())
	} else {
		err = fOne.Decode(&result)
		if err != nil {
			log.Error("fOne.Decode err:", err)
		}
		fmt.Println(result)
	}

	// delete
	_, err = collection.DeleteMany(ctx, bson.M{"uid": userId})
	if err != nil {
		fmt.Println("collection.DeleteMany err:", err)
	}
	// delete
	_, err = collection.DeleteMany(ctx, bson.M{"uid": userId + 1})
	if err != nil {
		fmt.Println("collection.DeleteMany err:", err)
	}

	// find
	fOne = collection.FindOne(ctx, bson.M{"uid": userId}, &options.FindOneOptions{})
	if fOne.Err() != nil {
		log.Warn("collection.FindOne fOne.Err():", fOne.Err())
	} else {
		err = fOne.Decode(&result)
		if err != nil {
			log.Error("fOne.Decode err:", err)
		}
		fmt.Println(result)
	}

	// insert
	_, err = collection.InsertOne(ctx, bson.M{"uid": userId, "info_json": infoJson, "quest1":123})
	if err != nil {
		log.Error("collection.InsertOne err:", err)
	}

	// find
	fOne = collection.FindOne(ctx, bson.M{"uid": userId}, &options.FindOneOptions{})
	if fOne.Err() != nil {
		log.Warn("collection.FindOne fOne.Err():", fOne.Err())
	} else {
		err = fOne.Decode(&result)
		if err != nil {
			log.Error("fOne.Decode err:", err)
		}
		fmt.Println(result)
	}

	// update
	_, err = collection.UpdateOne(ctx, bson.M{"uid":userId}, bson.M{"$set":bson.M{"quest1":345, "quest2": "helloworld"}})
	if err != nil {
		log.Error("collection.UpdateOne err", err)
	}

	// find
	fOne = collection.FindOne(ctx, bson.M{"uid": userId}, &options.FindOneOptions{})
	if fOne.Err() != nil {
		log.Warn("collection.FindOne fOne.Err():", fOne.Err())
	} else {
		err = fOne.Decode(&result)
		if err != nil {
			log.Error("fOne.Decode err:", err)
		}
		fmt.Println(result)
	}

	defer cancel()
}
