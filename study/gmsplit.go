package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "setprop level  80"
	s := strings.Split(str, " ")
	for id, e := range s {
		fmt.Println(strconv.Itoa(id) + "\t" + e)
	}
}
