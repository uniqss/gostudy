package main

import "fmt"

func main() {
	/*
		int:0
		float 0.0->0
		string:""
		array:[000]

		slice:nil
		map:nil
	*/

	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]int{"Go":88, "Python":77, "Java":66}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	//map1[1] = "hello" // panic: assignment to entry in nil map

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	if map1 == nil {
		map1 = make(map[int]string)
	}
	map1[1] = "hello world"
	map1[2] = " world"
	map1[3] = "memeda"
	map1[4] = "王二狗"
	map1[5] = "ruby"
	map1[6] = "三胖思密达"

	// 如果key存在，获取数值，如果key不存在，获取的是类型的默认值(零值)
	fmt.Println(map1)
	fmt.Println(map1[4])
	fmt.Println(map1[40])//""
	value, ok := map1[40]
	if ok {
		fmt.Println("对应的数值是：", value)
	} else {
		fmt.Println("操作的key值不存在，获取到的是零值：", value)
	}

	fmt.Println(map1)
	map1[3] = "如花"
	fmt.Println(map1)

	delete(map1, 3)
	fmt.Println(map1)

	// key不存在也不会崩
	delete(map1, 3000)
	fmt.Println(map1)

	//len  map没有cap
	fmt.Println(len(map1))
}
