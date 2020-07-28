package main

import (
	"context"
	"db_pressure_test/common"
	"db_pressure_test/wrappers"
	"fmt"
	"runtime"
	"sync"
	"time"
	//fastprefork "github.com/valyala/fasthttp/prefork"
)

func DB() {
	w := common.AcquireWorld()

	//w := new(common.World)

	id := 1
	fmt.Println("wrappers.Db.Query sql:", wrappers.WorldSelectSQL, " id:", id)
	rows, err := wrappers.Db.Query(context.Background(), wrappers.WorldSelectSQL, id)

	if err != nil {
		fmt.Println("wrappers.Db.Query err:", err)
		return
	}
	//err = rows.Scan(&w.ID, &w.RandomNumber) // nolint:errcheck
	//if err != nil {
	//	fmt.Println("rows.Scan err:", err)
	//}
	for rows.Next() {
		err := rows.Scan(&w.ID, &w.RandomNumber)
		if err != nil {
			fmt.Println(err)
		}
		break
	}
	fmt.Println(w)

	//data, _ := json.Marshal(w)
	//fmt.Println(data)

	common.ReleaseWorld(w)
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

func pressureTest(idx int, connCount int) {
	wg.Add(1)
	//u := common.User{UId: 1, InfoJson: ""}
	ctx := context.Background()
	insertCount := 10000000

	bInsert := 1 == 1
	if bInsert {
		users := make([]common.User, insertCount/connCount)
		for i := 0; i < insertCount/connCount; i++ {
			users[i].UId = int32(i*connCount + idx + 1)
			users[i].InfoJson = fmt.Sprintf("{\"InfoId\":\"%v\"}", users[i].UId)
		}
		for i := 0; i < insertCount / connCount; i++ {
			//u.UId = int32(i + 1)
			//u.InfoJson = fmt.Sprintf("{\"InfoId\":\"%v\"}", u.UId)
			//_, err := wrappers.Db.Exec(ctx, "insert into pressure_test(uid, info_json) values($1, $2)", u.UId, u.InfoJson)
			_, err := wrappers.Db.Exec(ctx, "insert into pressure_test(uid, info_json) values($1, $2)", users[i].UId, users[i].InfoJson)
			if err != nil {
				//fmt.Println("wrappers.Db.Exec err:", err, " u.UId:", u.UId, " u.InfoJson:", u.InfoJson)
				fmt.Println("wrappers.Db.Exec err:", err, " u.UId:", users[i].UId, " u.InfoJson:", users[i].InfoJson)
			}
		}
	}
	wg.Done()
}

func numCPU() int {
	n := runtime.NumCPU()
	if n == 0 {
		n = 8
	}

	return n
}

var wg sync.WaitGroup

func main() {
	maxConn := numCPU() * 4
	//if fastprefork.IsChild() {
	//	maxConn = numCPU()
	//}
	fmt.Println("maxConn:", maxConn)
	//ctx, cancelFunc := context.WithCancel(context.Background())

	//go func() {
	err := wrappers.InitDB(maxConn)
	if err != nil {
		fmt.Println("wrappers.InitDB err:", err)
	}
	//}()
	//for i:=1;i < 10;i++ {
	//	DB()
	//}
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
}
