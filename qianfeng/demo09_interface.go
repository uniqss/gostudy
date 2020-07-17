package main

import "fmt"

func main() {
	var a1 A = Cat{"花猫"}
	var a2 A = Person{"王二狗", 30}
	var a3 A = "haha"
	var a4 A = 100
	fmt.Println(a1, a2, a3, a4)
	test1(a1)
	test1(a2)
	test1(a3)
	test1(a4)

	test2(a3)

	map1 := make(map[string]interface{})
	map1["name"] = "李小花"
	map1["age"] = 30
	map1["friend"] = Person{"Jerry", 10}
	fmt.Println(map1)

	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, 100, "abc")
	fmt.Println(slice1)

	test3(slice1)
}

func test3(s1 []interface{}) {
	for i := 0; i < len(s1); i++ {
		fmt.Printf("第%d个数据：%v\n", i + 1, s1[i])
	}
}

func test1(a A) {
	fmt.Println(a)
}

func test2(a interface{}) {
	fmt.Println(a)
}

type A interface {
}

type Cat struct {
	color string
}
type Person struct {
	name string
	age  int
}
