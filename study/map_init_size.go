package main

import "fmt"

func main() {
	m := make(map[string]string, 2)
	m["aaa"] = "bbb"
	m["ccc"] = "ddd"
	m["fff"] = "ggg"

	m["xxx"] = "yyy"
	m["你好"] = "世界"

	fmt.Println(m)
}
