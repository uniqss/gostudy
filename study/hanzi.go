package main

import (
	"fmt"
	"unicode"
)

func IsChinese(r int32) bool {
	if unicode.Is(unicode.Scripts["Han"], r) {
		return true
	}
	return false
}

func IsLetter(r int32) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func IsDigit(r int32) bool {
	return r >= '0' && r <= '9'
}

func mainSimple(){
	str := "你好，世界！asdf1234ASDF"
	for idx, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Println(idx, " 是汉字")
		} else {
			fmt.Println(idx, " 不是汉字")
		}
	}
}

func main(){
	str := "你好，世界！asdf1234ASDF"
	for idx, r := range str {
		if IsChinese(r) {
			fmt.Println(idx, " 是汉字")
		} else if IsDigit(r) {
			fmt.Println(idx, " 是数字")
		} else if IsLetter(r) {
			fmt.Println(idx, " 是字母")
		} else {
			fmt.Println(idx, " 是火星文")
		}
	}
}
