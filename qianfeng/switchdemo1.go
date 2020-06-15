package main

import "fmt"

func main() {
	num := 0
	switch num {
	case 1:
		fmt.Println("第1季度")
	case 2:
		fmt.Println("第2季度")
	case 3:
		fmt.Println("第3季度")
	case 4:
		fmt.Println("第4季度")
	default:
		fmt.Println("数据有误，必须为1-4的整数")
	}

	num1 := 0
	num2 := 0
	oper := ""
	fmt.Println("请输入一个整数num1")
	fmt.Scan(&num1)
	fmt.Println("请输入一个整数num2")
	fmt.Scan(&num2)
	fmt.Println("请输入操作： + - * /")
	fmt.Scanln(&oper)
	switch oper {
	case "+":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	}

	fmt.Println("main over...")
}
