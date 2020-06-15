package main

import "fmt"

func main(){
	fmt.Println("main")
	var s1 string
	s1 = "王二狗"
	fmt.Printf("%T, %s\n", s1, s1)

	s2 := `hello world`
	fmt.Printf("%T, %s\n", s2, s2)

	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T, %d\n", v1, v1)
	fmt.Printf("%T, %s\n", v2, v2)

	v3 := '中'
	fmt.Printf("%T, %d %c\n", v3, v3, v3)

	fmt.Println("\"helloworld\"")
	fmt.Println("hell\nwor\tld")

	fmt.Println(`"helloworld"`)
	fmt.Println("hello`wor`ld")
}
