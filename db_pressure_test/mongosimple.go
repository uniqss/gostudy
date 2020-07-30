package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
	"time"
)

var __start int64
func printTimePre() {
	fmt.Println()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	__start = time.Now().UnixNano()
}
func printTimePost() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	__end := time.Now().UnixNano()
	fmt.Printf("cost is :%d us    %d ms    %d s \n", (__end-__start)/1000, (__end-__start)/1000000, (__end-__start)/1000000000)
}

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

	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()
	err = client.Ping(ctx2, readpref.Primary())
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
	ins := bson.M{"uid": userId, "info_json": infoJson, "quest1": 123}
	var oid primitive.ObjectID

	var result bson.M

	if 1 == 0 {
		// insert
		fmt.Println("calling collection.InsertOne. ins:", ins)
		insertResult, err := collection.InsertOne(ctx, ins)
		if err != nil {
			log.Error("collection.InsertOne err:", err)
		} else {
			fmt.Println("collection.InsertOne insertResult.InsertedID:", insertResult.InsertedID)
			oid = insertResult.InsertedID.(primitive.ObjectID)
		}
	}
	if 1 == 1 {
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

	if 1 == 1 {
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
	}

	if 1 == 0 {
		// find by _uid
		fOne := collection.FindOne(ctx, bson.M{"_id": oid}, &options.FindOneOptions{})
		if fOne.Err() != nil {
			log.Warn("collection.FindOne fOne.Err():", fOne.Err())
		} else {
			err = fOne.Decode(&result)
			if err != nil {
				log.Error("fOne.Decode err:", err)
			}
			fmt.Println(result)
		}
	}

	// insert
	fmt.Println("calling collection.InsertOne. ins:", ins)
	insertResult, err := collection.InsertOne(ctx, ins)
	if err != nil {
		log.Error("collection.InsertOne err:", err)
	} else {
		fmt.Println("collection.InsertOne insertResult.InsertedID:", insertResult.InsertedID)
		oid = insertResult.InsertedID.(primitive.ObjectID)
	}

	if 1 == 0 {
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
	}

	if 1 == 1 {
		// find by _uid
		fOne := collection.FindOne(ctx, bson.M{"_id": oid}, &options.FindOneOptions{})
		if fOne.Err() != nil {
			log.Warn("collection.FindOne fOne.Err():", fOne.Err())
		} else {
			err = fOne.Decode(&result)
			if err != nil {
				log.Error("fOne.Decode err:", err)
			}
			fmt.Println(result)
		}
	}

	if 1 == 0 {
		time.Sleep(time.Nanosecond)
	}

	if 1 == 0 {
		// update
		_, err = collection.UpdateOne(ctx, bson.M{"uid": userId}, bson.M{"$set": bson.M{"quest1": 345, "quest2": "helloWorld"}})
		if err != nil {
			log.Error("collection.UpdateOne err", err)
		}
	} else {
		// update by _uid
		_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": bson.M{"quest1": 345, "quest2": "helloWorld"}})
		if err != nil {
			log.Error("collection.UpdateOne err", err)
		}
	}

	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" {
			break
		}

		if input == "f" {
			printTimePre()
			if 1 == 1 {
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
			}
			printTimePost()
		}
		if input == "s" {
			printTimePre()
			if 1 == 1 {
				// find by _uid
				fOne := collection.FindOne(ctx, bson.M{"_id": oid}, &options.FindOneOptions{})
				if fOne.Err() != nil {
					log.Warn("collection.FindOne fOne.Err():", fOne.Err())
				} else {
					err = fOne.Decode(&result)
					if err != nil {
						log.Error("fOne.Decode err:", err)
					}
					fmt.Println(result)
				}
			}
			printTimePost()
		}
	}
}
