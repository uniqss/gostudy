package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type testStruct struct {
	tint    int
	tstring string
	tbool   bool
	tnum1   int
	tnum2   int
}

var tmap sync.Map

var testMapKeys map[int]string = make(map[int]string)
var testMapCount int = 100000

func goroutineAddMap() {
	key := rand.Intn(testMapCount)
	value, exists := tmap.Load(key)
	if exists {
		// check
		v := value.(*testStruct)
		if !checkTestStruct(v, key) {
			return
		}
	} else {
		tmap.Store(key, &testStruct{tint: key, tstring: "tstring" + strconv.Itoa(key), tbool: key%2 == 0})
	}
}

func checkTestStruct(v *testStruct, key int) bool {
	if v.tint != key {
		logrus.Errorf("goroutineAddMap v.tint error v.tint:%d key:%d", v.tint, key)
		return false
	}
	tstring := "tstring" + strconv.Itoa(key)
	if v.tstring != tstring {
		logrus.Errorf("goroutineAddMap v.tstring error v.tstring:%s tstring:%s", v.tstring, tstring)
		return false
	}

	tbool := key%2 == 0
	if v.tbool != tbool {
		logrus.Errorf("goroutineAddMap v.tbool error v.tbool:%v tbool:%v", v.tbool, tbool)
		return false
	}

	if v.tnum1 != v.tnum2 {
		logrus.Errorf("goroutineAddMap v.tnum1 != v.tnum2 v.tnum1:%v v.tnum2:%v", v.tnum1, v.tnum2)
	}
	return true
}

func goroutineDelMap() {
	key := rand.Intn(testMapCount)
	value, exists := tmap.Load(key)
	if exists {
		// check
		v := value.(*testStruct)
		if !checkTestStruct(v, key) {
			return
		}
		tmap.Delete(key)
	}
}

func goroutineChangeMap() {
	key := rand.Intn(testMapCount)
	value, exists := tmap.Load(key)
	if exists {
		// check
		v := value.(*testStruct)
		if !checkTestStruct(v, key) {
			return
		}
		num := rand.Int()

		v.tnum1 = num
		v.tnum2 = num
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < testMapCount; i++ {
		testMapKeys[i] = "test" + strconv.Itoa(i)
	}

	working := true
	for i:=0;i < 1000;i++ {
		go func() {
			for working {
				goroutineAddMap()
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}

	for i:=0;i < 0;i++ {
		go func() {
			for working {
				goroutineDelMap()
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}

	for i:=0;i < 110000;i++ {
		go func() {
			for working {
				goroutineChangeMap()
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}

	go func() {
		i := 0
		for working {
			if i % 150 == 0 {
				i = 0
				fmt.Println("working...")
			}
			i++
			time.Sleep(time.Millisecond * 100)
		}
	}()

	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" || input == "q" || input == "quit" {

			working = false
			break
		}
	}

	time.Sleep(time.Second)
}
