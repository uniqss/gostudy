package main

import (
	"encoding/json"
	"fmt"
)

type Test2 struct {
	XXX string
	YYY int32
}

type Test1 struct {
	AAA         string
	BBB         int32
	Propertymap map[string]string
	Inner1      *Test2
}

type Test1PropertySync struct {
	Propertymap map[string]string
}

func main() {
	t1 := new(Test1)
	t1.AAA = "aaaa"
	t1.BBB = 12345
	t1.Inner1 = new(Test2)
	t1.Inner1.XXX = "xxxxx"
	t1.Inner1.YYY = 2341341
	t1.Propertymap = make(map[string]string)
	t1.Propertymap["hp"] = "asdf"
	str, err := json.Marshal(t1)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}
	fmt.Println(string(str))

	t2 := new(Test1)
	err = json.Unmarshal(str, t2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t2)
	fmt.Println(t2.Inner1)

	fmt.Println("============================================================")
	t3 := new(Test1)
	jsonstr := "{\"AAA_KeyWrong\":\"aaaa\",\"BBB\":12345,\"Propertymap\":{\"hp\":\"asdf\"},\"Inner1\":{\"XXX\":\"xxxxx\",\"YYY\":2341341}}"
	err = json.Unmarshal(([]byte)(jsonstr), t3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t3)
	fmt.Println(t3.Inner1)
}
