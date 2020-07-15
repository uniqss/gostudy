package main

import "fmt"

func main() {
	/*
	array:[size]数据类型
	slice:[]数据类型
	map:map[key的类型]value的类型

	slice,map引用类型
	make()创建的都是引用类型  ,slice,map,channel都是引用类型
	 */

	map1 := make(map[int]string)
	map2 := make(map[string]float64)
	fmt.Printf("%T\n", map1)
	fmt.Printf("%T\n", map2)

	map3 := make(map[string]map[string]string)
	m1 := make(map[string]string)
	m1["name"] = "王二狗"
	m1["age"] = "30"
	m1["salary"] = "3000"
	map3["hr"] = m1

	m2 := map[string]string{"name":"ruby", "age":"28", "salary":"8000"}
	map3["总经理"] = m2
	fmt.Println(map3)

	map4 := make(map[string]string)
	map4["王二狗"] = "矮穷矬"
	map4["李小花"] = "白富美"
	map4["ruby"] = "住在隔壁"
	fmt.Println(map4)

	map5 := map4
	fmt.Println(map5)
	map5["王二狗"] = "高富帅"
	fmt.Println(map4, map5)
}
