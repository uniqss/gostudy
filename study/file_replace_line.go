package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	inputFile := ""
	if len(args) > 1 {
		inputFile = args[1]
	} else {
		log.Fatal("usage: executable inputFileName outputFileName")
		return
	}
	outputFile := inputFile
	if len(args) > 2 {
		outputFile = args[2]
	}
	fmt.Println("inputFile:", inputFile, " outputFile:", outputFile)
	iFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("os.Open failed. inputFile:", inputFile, " err:", err)
		return
	}
	defer iFile.Close()
	oFile, err := os.OpenFile(outputFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("os.OpenFile failed. err:", err)
		return
	}
	defer oFile.Close()

	dataWriter := bufio.NewWriter(oFile)

	scanner := bufio.NewScanner(iFile)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		text := scanner.Text()
		if text == "var $protobuf = require(\"protobufjs/minimal\");" {
			text = "var $protobuf = protobuf;"
		}
		dataWriter.WriteString(text + "\n")
	}
	dataWriter.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
