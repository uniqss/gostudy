package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "D:\\go\\src\\l_file\\aa/aa.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()

	bs := make([]byte, 4, 4)

	/*
		// 1
		n, err := file.Read(bs)
		fmt.Println(err)
		fmt.Println(n)
		fmt.Println(bs)
		fmt.Println(string(bs))

		// 2
		n, err = file.Read(bs)
		fmt.Println(err)
		fmt.Println(n)
		fmt.Println(bs)
		fmt.Println(string(bs))

		// 3
		n, err = file.Read(bs)
		fmt.Println(err)
		fmt.Println(n)
		fmt.Println(bs)
		fmt.Println(string(bs))

		// 4
		n, err = file.Read(bs)
		fmt.Println(err) // EOF
		fmt.Println(n) // 0
		fmt.Println(bs)
		fmt.Println(string(bs))
	*/

	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			break
		}
		fmt.Println(string(bs[:n]))
	}

}
