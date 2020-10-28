package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func GetCurrentGoroutineIdStr() string {
	var (
		buf [64]byte
		n = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)
	idField := strings.Fields(stk)[0]
	return idField
}

func pid(str string){
	fmt.Println(str, " pid:", GetCurrentGoroutineIdStr())
}

func main() {
	fmt.Println("main begin")
	pid("main begin ")

	time.AfterFunc(time.Second * 3, Callback)
	//timer := time.NewTimer(time.Second * 3)
	//timer.C

	time.Sleep(time.Second * 6)

	pid("main end ")
	fmt.Println("main end")
}

func Callback(){
	pid("Callback")
	time.AfterFunc(time.Second, Callback2)
}

func Callback2(){
	pid("Callback2")
}
