package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Printf("kind:%s, type:%s\n", value.Kind(), value.Type())

	value2 := reflect.ValueOf(fun2)
	value3 := reflect.ValueOf(fun3)
	fmt.Printf("kind:%s, type:%s\n", value2.Kind(), value2.Type())
	fmt.Printf("kind:%s, type:%s\n", value3.Kind(), value3.Type())

	value.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf("hello world")})
	resultValue := value3.Call([]reflect.Value{reflect.ValueOf(123), reflect.ValueOf("帅帅帅")})
	fmt.Printf("%T\n", resultValue)
	fmt.Println(len(resultValue))
	rv := resultValue[0]
	fmt.Printf("kind:%s type:%s\n", rv.Kind(), rv.Type())

	s := rv.Interface().(string)
	fmt.Println(s)
}

func fun1() {
	fmt.Println("我是函数fun1，无参的")
}
func fun2(i int, s string) {
	fmt.Println("我是函数fun2, 有参的", i, s)
}
func fun3(i int, s string) string {
	fmt.Println("我是函数fun3, 有参的，也有返回类型", i, s)
	return s + strconv.Itoa(i)
}
