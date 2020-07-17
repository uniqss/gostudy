package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 断点续传
	srcFile := "D:\\go\\src\\l_file\\aa/g.zip"
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1:]
	fmt.Println(destFile)
	tempFile := destFile + "temp.txt"
	fmt.Println(tempFile)

	file1, err := os.Open(srcFile)
	HandleErr(err)
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandleErr(err)
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	HandleErr(err)

	defer file1.Close()
	defer file2.Close()

	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := file3.Read(bs)
	HandleErr(err)
	countStr := string(bs[:n1])
	count, err := strconv.ParseInt(countStr, 10, 64)
	HandleErr(err)
	fmt.Println(count)

	// step2:设置读、写的位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 4096, 4096)
	n2 := -1 // 读取的数据量
	n3 := -1 // 写出的数据量
	total := int(count) //读取的总量

	// step3:复制文件
	for{
		n2, err = file1.Read(data)
		if err == io.EOF||n2 == 0 {
			fmt.Println("文件复制完毕。总大小：", total)
			file3.Close()
			os.Remove(tempFile)
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3

		// 将复制的总量，存储到零时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		//// 假装断电
		//if total >8000 {
		//	panic("断电了")
		//}
	}
}

func HandleErr(err error){
	if err != nil {
		fmt.Println(err)
	}
}