package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	s1 := Student{"王二狗", 19, "美国克莱登大学"}
	fmt.Printf("%T\n", s1)
	p1 := &s1
	fmt.Printf("%T\n", p1)
	fmt.Println(s1.Name)
	fmt.Println((*p1).Name, p1.Name)

	//改变数值
	value := reflect.ValueOf(&s1)
	if value.Kind() == reflect.Ptr {
		newValue := value.Elem()
		fmt.Println(newValue.CanSet())

		f1 := newValue.FieldByName("Name")
		f1.SetString("李小花")
		f3 := newValue.FieldByName("School")
		f3.SetString("幼儿园")
		fmt.Println(s1)
	}
}
