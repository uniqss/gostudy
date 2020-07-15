package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "D:\\go\\src\\l_file\\aa/cc.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w1 := bufio.NewWriter(file)
	//n, err := w1.WriteString("hello world")
	//fmt.Println(err)
	//fmt.Println(n)
	//w1.Flush()

	for i := 1; i <= 10000; i++ {
		w1.WriteString(fmt.Sprintf("%d:hello", i))
	}
	w1.Flush()
}
