package main

import "fmt"

func main() {
	// switch break
	// fallthrough 穿透
	n := 2
	switch n {
	case 1:
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
	case 2:
		fmt.Println("我是熊二")
		fmt.Println("我是熊二")
		break
		fmt.Println("我是熊二")
	case 3:
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
	}

	fmt.Println("-------------------------")
	m := 2
	switch m {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fallthrough
	case 3:
		fmt.Println("第三季度")
		fallthrough
	case 4:
		fmt.Println("第四季度")

	}

	fmt.Println("main over")
}
