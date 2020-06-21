package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	fileName1 := "D:\\go\\src\\l_file\\demo02_file.go"
	fileName2 := "demo02_file.go"
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	fmt.Println("获取父目录：", path.Join(fileName1, ".."))

	//err := os.Mkdir("D:\\go\\src\\l_file\\aa", os.ModePerm)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//err :=os.MkdirAll("D:\\go\\src\\l_file\\bb\\cc\\dd", os.ModePerm)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("create dir ok")
	//
	//file1, err := os.Create("D:\\go\\src\\l_file\\aa/ab.txt")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("create file ok file1:", file1)

	//file2, err := os.Create("D:\\go\\src\\l_file\\aa/bb.txt")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println(file2)

	//file3, err := os.Open(fileName1)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//fmt.Println(file3)

	//file4, err := os.OpenFile(fileName1, os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("file4:", file4)
	//
	//file4.Close()

	//err := os.Remove("D:\\go\\src\\l_file\\aa/ab.txt")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("删除成功")

	// 慎用，不入回收站
	err := os.RemoveAll("D:\\go\\src\\l_file\\aa")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("删除成功")
}
