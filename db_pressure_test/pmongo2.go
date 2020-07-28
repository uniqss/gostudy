package main

import (
	"db_pressure_test/common"
	"db_pressure_test/wrappers"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	bInsert := 1 == 1
	if bInsert {
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
	} else {
		clearDB()
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
	collection := wrappers.MongoDB.C("ptest_tbl")
	err := collection.DropCollection()
	if err != nil {
		fmt.Println("clearDB collection.DropCollection err:", err)
	}
}

func pressureTest(idx int, connCount int) {
	wg.Add(1)
	insertCount := 10000000

	bInsert := 1 == 1
	if bInsert {
		users := make([]common.User, insertCount/connCount)
		for i := 0; i < insertCount/connCount; i++ {
			users[i].UId = int32(i*connCount + idx + 1)
			users[i].InfoJson = fmt.Sprintf("{\"InfoId\":\"%v\"}", users[i].UId)
		}

		collection := wrappers.MongoDB.C("ptest_tbl")
		for i := idx; i < insertCount/connCount; i++ {
			err := collection.Insert(bson.M{"UId": users[i].UId, "InfoJson": users[i].InfoJson})
			//_, err := collection.UpsertId(bson.NewObjectId(), bson.M{"UId": users[i].UId, "InfoJson": users[i].InfoJson})
			if err != nil {
				fmt.Println("pressureTest collection.UpsertId err:", err)
			}
		}
	}
	wg.Done()
}
