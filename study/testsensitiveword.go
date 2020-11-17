package main

import (
	"fmt"
	textcensor "github.com/kai1987/go-text-censor"
)

func InitSensitiveWords(word string) bool {
	var words []string
	words = append(words, "你好")
	words = append(words, "世界")
	words = append(words, "一二三")
	words = append(words, "四五六")
	words = append(words, "12345")
	words = append(words, word)

	textcensor.InitWords(words, false)
	punctuation := "0123456789abcdefghijklmnopqrstuvwxyz !\"#$%&'()*+,-./:;<=>?@[\\]^_`|~，。？；：”’￥（）——、！……•"
	textcensor.SetPunctuation(punctuation)
	return true
}

func CheckWordSensitive(str string) bool {
	//for _, value := range DataTables.GetDT_Sensitive_Words_Config().DT_Sensitive_Words_ConfigItems {
	//	if strings.Contains(str, value.Sensitive) {
	//		return false
	//	}
	//}

	return textcensor.IsPass(str, true)
}

func _CheckWord(word string) {
	ok := CheckWordSensitive(word)
	fmt.Println("CheckWordSensitive word:", word, "\t\t", ok)
}

func main() {
	var ok bool
	ok = InitSensitiveWords("")
	fmt.Println("InitSensitiveWords ok:", ok)

	_CheckWord("我好")
	_CheckWord("你好")
	_CheckWord("你*好")
	_CheckWord("你好帅")
	_CheckWord("世界")
	_CheckWord("你好世界")
	_CheckWord("11911")

	ok = InitSensitiveWords("11911")
	fmt.Println("InitSensitiveWords ok:", ok)

	_CheckWord("我好")
	_CheckWord("你好")
	_CheckWord("你*好")
	_CheckWord("你好帅")
	_CheckWord("世界")
	_CheckWord("你好世界")
	_CheckWord("11911")

	ok = InitSensitiveWords("11911")
	fmt.Println("InitSensitiveWords ok:", ok)

	_CheckWord("我好")
	_CheckWord("你好")
	_CheckWord("你*好")
	_CheckWord("你好帅")
	_CheckWord("世界")
	_CheckWord("你好世界")
	_CheckWord("11911")
}
