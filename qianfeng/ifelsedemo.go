package main

import "fmt"

func main(){
	/*
	score := 0
	fmt.Println("请输入成绩")
	fmt.Scanln(&score)
	if score >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}
	 */

	gender := "男"
	fmt.Scanln(&gender)
	if gender == "男" {
		fmt.Println("去男厕所")
	} else {
		fmt.Println("去女厕所")
	}

	fmt.Println("main over...")
}
