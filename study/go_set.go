package main

import "fmt"

func main() {
	set := make(map[string]bool)
	set["aaa"] = true
	set["bbb"] = false
	fmt.Println(set["aaa"])
	fmt.Println(set["bbb"])
	fmt.Println(set["ccc"])
}
