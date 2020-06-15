package main

import "fmt"

func main() {
	sex := "泰国"
	if sex == "男" {
		fmt.Println("去男厕所")
	} else {
		if sex == "女" {
			fmt.Println("去女厕所")
		} else {
			fmt.Println("我也不知道了")
		}
	}
}
