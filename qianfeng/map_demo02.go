package main

import (
	"fmt"
	"sort"
)

func main() {
	// for range
	map1 := make(map[int]string)
	map1[1] = "红孩儿"
	map1[2] = "小钻风"
	map1[3] = "白骨精"
	map1[4] = "白素贞"
	map1[5] = "王二狗"
	map1[6] = "银角大王"

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println("=================")
	for i := 1; i <= len(map1); i++ {
		fmt.Println(i, "==>", map1[i])
	}

	// 获取key 排序 遍历
	keys := make([]int, 0, len(map1))
	fmt.Println(keys)

	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, "=====>", map1[key])
	}

	s1 := []string{"Apple", "Windows", "Orange", "abc", "王二狗", "acd", "acc"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
}
