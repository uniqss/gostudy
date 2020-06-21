package main

import (
	"fmt"
	"os"
)

func main() {
	// FileInfo
	fileInfo, err := os.Stat("filedemo.go")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Mode())
}
