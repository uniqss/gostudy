package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "hello world"
	fmt.Println(strings.Contains(s1, "abc"))
	fmt.Println(strings.Contains(s1, "ello"))

	fmt.Println(strings.ContainsAny(s1, "abcd"))

	fmt.Println(strings.Count(s1, "ll"))

	s2 := "20190525课堂笔记.txt"
	if strings.HasPrefix(s2, "2019") {
		fmt.Println("19年的文件。。")
	}

	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("文本文档")
	}

	fmt.Println(strings.Index(s1, "lloo"))
	fmt.Println(strings.IndexAny(s1, "abcd"))

	fmt.Println(strings.LastIndex(s1, "l"))

	ss1 := []string{"abc", "world", "hello", "ruby"}
	s3 := strings.Join(ss1, "*")
	fmt.Println(s3)

	s4 := "123,4566,aaa,bbb,34234"
	ss2 := strings.Split(s4, ",")
	for i:= 0;i < len(ss2);i++ {
		fmt.Println(ss2[i])
	}

	s5 := strings.Repeat("hello", 5)
	fmt.Println(s5)

	s6 := strings.Replace(s1, "l", "*", -1)
	fmt.Println(s6)

	s7 := "Hello World***123..."
	fmt.Println(strings.ToLower(s7))
	fmt.Println(strings.ToUpper(s7))

	// substr
	// str[start:end]    [start, end)
	fmt.Println(s1)
	s8 := s1[0:5]
	fmt.Println(s8)
}
