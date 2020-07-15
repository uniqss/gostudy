package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dirname := "D:/go/src/l_file/"
	listFiles(dirname, 0)
}

func listFiles(dirName string, depth int) {
	s := "|--"
	for i := 0; i < depth; i++ {
		s = "|  " + s
	}
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		fileName := dirName + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, fileName)
		if fi.IsDir() {
			listFiles(fileName, depth+1)
		}
	}
}
