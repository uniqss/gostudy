package main

import (
	"fmt"
	"time"
)

func main() {
	// After其实就是 return NewTimer(d).C
	ch := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch)
	fmt.Println(time.Now())

	time2 := <-ch
	fmt.Println(time2)
}
