package main

import "fmt"

func main() {
	defer func(){
		if msg := recover(); msg != nil{
			fmt.Println(msg, "程序恢复啦")
		}
	}()
	funA()
	defer myprint("defer main 3...")
	funB()
	defer myprint("defer main 4...")

	fmt.Println("main over...")
}
func myprint(s string) {
	fmt.Println(s)
}
func funA() {
	fmt.Println("我是一个函数funA()...")
}
func funB() {
	fmt.Println("我是函数funB()...")
	defer myprint("defer funB 1")
	for i := 1; i <= 3; i++ {
		fmt.Println("i:", i)
		if i == 2 {
			// panic发生时，前面的defer全部会被调用，后面的不会
			panic("funB函数，恐慌了")
		}
	}
	defer myprint("defer funB 2")
}
