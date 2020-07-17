package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "D:\\go\\src\\l_file\\aa/ab.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//b1 := bufio.NewReader(file)
	//p := make([]byte, 1024)
	//n1, err := b1.Read(p)
	//fmt.Println(n1)
	//fmt.Println(string(p))

	//data, flag, err := b1.ReadLine()
	//fmt.Println(flag)
	//fmt.Println(err)
	//fmt.Println(data)
	//fmt.Println(string(data))

	//s1, err := b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)
	//
	//s1, err = b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)
	//
	//for {
	//	s1, err := b1.ReadString('\n')
	//	if err == io.EOF {
	//		fmt.Println("读取完毕")
	//		break
	//	}
	//	//fmt.Println(err)
	//	fmt.Println(s1)
	//}

	//data, err := b1.ReadBytes('\n')
	//fmt.Println(err)
	//fmt.Println(string(data))
	//
	//s2 := ""
	//fmt.Scanln(&s2)
	//fmt.Println(s2)
	b2 := bufio.NewReader(os.Stdin)
	s2, _ := b2.ReadString('\n')
	fmt.Println(s2)
}
