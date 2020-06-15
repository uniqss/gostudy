package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("aa" + strconv.Itoa(100))
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T, %t\n", b1, b1)

	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T, %s\n", ss1, ss1)

	s2 := "100"
	i2, err := strconv.ParseInt(s2, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T, %d\n", i2, i2)

	ss2 := strconv.FormatInt(i2, 10)
	fmt.Printf("%T, %s\n", ss2, ss2)

	i3, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T, %d\n", i3, i3)

	ss3 := strconv.Itoa(i3)
	fmt.Printf("%T, %s\n", ss3, ss3)
}
