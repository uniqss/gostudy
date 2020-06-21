package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "D:\\go\\src\\l_file\\aa/ab.txt"
	//file, err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//bs := []byte{'A', 'B', 'C', 'D', 'E', 'F'}
	bs := []byte{'a', 'b', 'c', 'd'}
	n, err := file.Write(bs[:2])
	HandleErr(err)
	fmt.Println(n)

	n, err = file.WriteString("HelloWorld\n")
	HandleErr(err)
	fmt.Println(n)

	n, err = file.Write([]byte("Today\n"))
	HandleErr(err)
	fmt.Println(n)
}

func HandleErr(err error){
	if err != nil {
		fmt.Println(err)
	}
}