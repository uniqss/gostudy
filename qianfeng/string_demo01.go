package main

import "fmt"

func main() {
	s1 := "hello中国"
	s2 := "hello world"
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(len(s1))
	fmt.Println(len(s2))

	fmt.Println(s2[0])
	a := 'h'
	b := 104
	fmt.Printf("%c %c %c\n", s2[0], a, b)

	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c\t", s2[i])
	}

	for _, v := range s2 {
		fmt.Printf("%c", v)
	}

	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1)
	fmt.Println(s3)

	s4 := "abcdef"
	slice2 := []byte(s4)
	fmt.Println(slice2)

	fmt.Println(s4)
	//s4[3] = 'a'// cannot assign to s4[3]
}
