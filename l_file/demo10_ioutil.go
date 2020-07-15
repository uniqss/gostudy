package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//fileName := "D:\\go\\src\\l_file\\aa/aa.txt"
	//data,err := ioutil.ReadFile(fileName)
	//fmt.Println(err)
	//fmt.Println(string(data))

	//fileName := "D:\\go\\src\\l_file\\aa/bbb.txt"
	////s1 := "helloworld面朝大海春暖花开"
	//s1 := "床前明月光，地上鞋三双"
	//err := ioutil.WriteFile(fileName, []byte(s1), os.ModePerm)
	//fmt.Println(err)

	//s2 := "王二狗和李小花是好朋友，hello world, hahahaha"
	//r1 := strings.NewReader(s2)
	//data, err := ioutil.ReadAll(r1)
	//fmt.Println(err)
	//fmt.Println(string(data))

	//dirName := "D:\\go\\src\\l_file\\"
	//fileInfos, err := ioutil.ReadDir(dirName)
	//fmt.Println(err)
	//fmt.Println(len(fileInfos))
	//for i := 0; i < len(fileInfos); i++ {
	//	fmt.Printf("第%d个：名字：%s，是否是目录：%t\n", i, fileInfos[i].Name(), fileInfos[i].IsDir())
	//}

	dir, err := ioutil.TempDir("D:/go/src/l_file/aa", "Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cleanup(dir)
	fmt.Println(dir)

	file, err := ioutil.TempFile(dir, "Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer os.Remove(file.Name())
	fmt.Println(file.Name())
	// 一定要close，否则会说
	// The process cannot access the file because it is being used by another process
	defer file.Close()
}

func cleanup(dir string){
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Println(err)
	}
}

