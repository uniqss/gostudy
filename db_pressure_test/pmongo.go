package main

import (
	"context"
	"db_pressure_test/common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	op := common.OP_UPDATE

	if op == common.OP_INSERT {
		time.Sleep(time.Second * 3)
		connCount := 8
		start := printTimePre()
		for i := 0; i < connCount; i++ {
			go pressureTest(i, connCount)
		}
		time.Sleep(time.Second * 1)
		wg.Wait()
		printTimePost(start)
		//wrappers.CloseDB()
		//cancelFunc()
	} else if op == common.OP_DELETE {
		clearDB()
	} else if op == common.OP_UPDATE {
		time.Sleep(time.Second * 3)
		connCount := 8
		start := printTimePre()
		for i := 0; i < connCount; i++ {
			go ptestUpdateMongo1(i, connCount)
		}
		time.Sleep(time.Second * 1)
		wg.Wait()
		printTimePost(start)
	} else if op == common.OP_SELECT {
		
	}
}

func printTimePre() int64 {
	fmt.Println()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	start := time.Now().UnixNano()
	return start
}

func printTimePost(start int64) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	end := time.Now().UnixNano()
	fmt.Printf("cost is :%d \n", (end-start)/1000)
}

func clearDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.87.210:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("client.Ping err:", err)
	}

	collection := client.Database("testing").Collection("numbers")

	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//err = collection.Drop(ctx)
	//err = client.Database("testing").Drop(context.Background())
	err = collection.Drop(context.Background())
	filter := bson.M{}
	opts := options.DeleteOptions{}
	_, err = collection.DeleteMany(ctx, filter, &opts)
	if err != nil {
		fmt.Println(err)
	}
	//res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	//id := res.InsertedID
	//
	//fmt.Println(id)
}

func pressureTest(idx int, connCount int) {
	wg.Add(1)
	insertCount := 10000000

	users := make([]common.User, insertCount/connCount)
	for i := 0; i < insertCount/connCount; i++ {
		users[i].UId = int32(i*connCount + idx + 1)
		users[i].InfoJson = fmt.Sprintf("{\"InfoId\":\"%v\"}", users[i].UId)
	}

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
		fmt.Println("client.Ping err:", err)
	}

	collection := client.Database("testing").Collection("numbers")

	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	ctx = context.Background()
	for i := idx; i < insertCount/connCount; i++ {
		//res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		//id := res.InsertedID
		//fmt.Println(id)
		_, err := collection.InsertOne(ctx, bson.M{"uid": users[i].UId, "info_json": users[i].InfoJson})
		if err != nil {
			fmt.Println("collection.InsertOne err", err)
		}
	}
	defer cancel()
	wg.Done()
}

func ptestUpdateMongo1(idx int, connCount int){
	wg.Add(1)
	insertCount := 10000000

	users := make([]common.User, insertCount/connCount)
	for i := 0; i < insertCount/connCount; i++ {
		users[i].UId = int32(rand.Intn(insertCount/connCount))
		users[i].InfoJson = fmt.Sprintf("{\"InfoId\":\"%v\"}", users[i].UId)
	}

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
		fmt.Println("client.Ping err:", err)
	}

	collection := client.Database("testing").Collection("numbers")

	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	ctx = context.Background()
	for i := idx; i < insertCount/connCount; i++ {
		//res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		//id := res.InsertedID
		//fmt.Println(id)
		_, err := collection.UpdateOne(ctx, bson.M{"uid":users[i].UId}, bson.M{"$set":bson.M{"uid": users[i].UId, "info_json": users[i].InfoJson}})
		//_, err := collection.InsertOne(ctx, bson.M{"uid": users[i].UId, "info_json": users[i].InfoJson})
		if err != nil {
			log.Error("collection.UpdateOne err", err)
		}
	}
	defer cancel()
	wg.Done()
}