package main

import "fmt"

func main() {
	s1 := Student{"张三", 18}
	fmt.Println(s1.name, s1.age)

	func(){
		fmt.Println("hello world...")
	}()

	s2 := struct{
		name string
		age int
	}{
		"李四",
		19,
	}
	fmt.Println(s2.name, s2.age)

	//w1 := Worker{"王二狗", 30}
	//fmt.Println(w1.name, w1.age)

	w2 := Worker{"王二狗", 30}
	fmt.Println(w2)
	fmt.Println(w2.string)
	fmt.Println(w2.int)
}

type Worker struct {
	//name string
	//age int
	string
	int
	// 匿名字段类型不能重复
	// string
}

type Student struct {
	name string
	age int
}
