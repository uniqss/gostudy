package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"strings"
)

const EntityDefinitionFile = "entity.json"

func ReadAllDefinitions() {
	files, err := ioutil.ReadDir("./definition/")
	if err != nil {
		fmt.Println("ioutil.ReadDir err:", err)
		return
	}
	for _, f := range files {
		fName := f.Name()
		if strings.HasSuffix(fName, ".json") {
			//fmt.Println("fName:", fName)
			//definitionName := strings.TrimSuffix(fName, ".json")
			//fmt.Println("definitionName:", definitionName)

			if fName == "entity.json" && 1 == 0 {
				// 处理 entity 这里是definition，并不是实例
				jsonFile, err := os.Open("./definition/" + fName)
				if err != nil {
					fmt.Println("os.Open failed. err:", err, " fName:", fName)
					return
				}
				defer jsonFile.Close()
				byteValue, _ := ioutil.ReadAll(jsonFile)
				result := gjson.GetBytes(byteValue, "entities")
				result.ForEach(func(key, value gjson.Result) bool {
					fmt.Println("key.String():", key.String(), "value.String():", value.String())
					for i:= 0; i < len(value.Array()); i++ {
						c := value.Array()[i]
						fmt.Println("component value:", c.String())
					}
					return true
				})
			} else {
				// 处理 component 这里是definition，并不是实例
				jsonFile, err := os.Open("./definition/" + fName)
				if err != nil {
					fmt.Println("os.Open failed. err:", err, " fName:", fName)
					return
				}
				defer jsonFile.Close()
				byteValue, _ := ioutil.ReadAll(jsonFile)
				properties := gjson.GetBytes(byteValue, "properties")
				properties.ForEach(func(key, value gjson.Result) bool {
					//fmt.Println("key.String():", key.String(), "value.String():", value.String())
					fmt.Println("key.String():", key.String())
					value.ForEach(func(pKey, pType gjson.Result) bool {
						fmt.Println("pKey.String():", pKey.String(), " pType.String()", pType.String())
						return true
					})
					return true
				})
			}
		}
	}
}

func main() {
	ReadAllDefinitions()
}
