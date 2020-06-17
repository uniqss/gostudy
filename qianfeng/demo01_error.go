package main

import (
	"fmt"
	"os"
)

func main() {
	//f, err := os.Open("./test.txt")
	f, err := os.Open("demo01_error.go")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.op:", ins.Op)
			fmt.Println("2.Path:", ins.Path)
			fmt.Println("3.Err", ins.Err)
		}
		return
	}

	fmt.Println(f.Name(), "打开文件成功")
}
