package utils

import "fmt"

func Count() {
	fmt.Println("utils.Count() ")
}


func init(){
	fmt.Println("utils包里的init()")
}

func init(){
	fmt.Println("utils包里的init() 第2个。。")
}