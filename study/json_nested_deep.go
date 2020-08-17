package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json, "name.last")
	println(value.String())

	name := gjson.Get(json, "name").String()
	fmt.Println(name)
	last := gjson.Get(name, "last")
	fmt.Println(last)
}