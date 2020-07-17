package main

import "fmt"

func main() {
	b1 := Book{}
	b1.bookName = "西游记"
	b1.price = 12.34

	s1 := Student{}
	s1.name = "王二狗"
	s1.age = 18
	s1.book = b1
	fmt.Println(b1)
	fmt.Println(s1)

	fmt.Println(s1.book.bookName, s1.book.price)

	s2 := Student{"李小花", 19, Book{"Go语言是怎样练成的", 23.4}}
	fmt.Println(s2)

	s1.book.bookName = "红楼梦"
	fmt.Println(s1)
	fmt.Println(b1)

	b4 := Book{bookName: "呼啸山庄", price: 33.44}
	s4 := Student2{name :"王二麻子", age:18, book:&b4}
	fmt.Println(b4)
	fmt.Println(s4.book)

	s4.book.bookName = "雾都孤儿"
	fmt.Println(b4)
	fmt.Println(s4.book)
}

type Book struct {
	bookName string
	price    float64
}

type Student struct {
	name string
	age  int
	book Book
}

type Student2 struct {
	name string
	age int
	book *Book
}
