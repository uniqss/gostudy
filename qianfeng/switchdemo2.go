package main

import "fmt"

func main(){
	switch {
	case true:
		fmt.Println("true....")
	case false:
		fmt.Println("false...")
	}
	score := 88
	switch {
	case score >=0 && score < 60:
		fmt.Println(score, "不及格")
	case score >= 60 && score < 70:
		fmt.Println(score, "及格")
	case score >= 70 &&score < 80:
		fmt.Println(score, "中等")
	case score >= 80 &&score < 90:
		fmt.Println("良好")
	case score >= 90 &&score < 100:
		fmt.Println("良好")
	default:
		fmt.Println(score, "成绩有误。。。")
	}

	letter := "A"
	switch letter {
	case "A", "E", "I", "O", "U":
		fmt.Println(letter, "是元音")
	case "M", "N":
		fmt.Println(letter, "是M或N")
	default:
		fmt.Println("其他")
	}

	year := 2020
	month := 2
	day := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if (year % 400 == 0||year %4 == 0 &&year % 100 != 0) {
			day = 29
		} else {
			day = 28
		}
	default:
		fmt.Println("月份有误")
	}
	fmt.Println(day)

	switch language := "java";language {
	case "golang":
		fmt.Println("GO")
	case "java":
		fmt.Println("Java")
	default:
		fmt.Println("other")
	}
}