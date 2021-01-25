package main

import "log"

func SSSS(arr []interface{}) []interface{} {
	for i := 5; i < 10; i++ {
		var tmp uint32
		tmp = uint32(i)
		arr = append(arr, tmp)
	}

	return arr
}

func main() {
	var s []interface{}
	for i := 0; i < 5; i++ {
		var tmp uint32
		tmp = uint32(i)
		s = append(s, tmp)
	}

	s = SSSS(s)

	log.Println(s)
}
