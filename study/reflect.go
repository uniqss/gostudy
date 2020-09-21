package main

import (
	"fmt"
	"reflect"
)

func pTypeName(val interface{}) {
	tName := reflect.TypeOf(val).Name()
	fmt.Println(tName)
}

func AttributesDefSafeTrans(typeDef string, valInt *int64, valString *string, val interface{}) {
	if typeDef == "string" {
		*valInt = 0
		*valString = val.(string)
	} else if typeDef == "int64" {
		typeName := reflect.TypeOf(val).Name()
		if typeName == "int64" {
			*valInt = val.(int64)
		} else if typeName == "int32" {
			*valInt = int64(val.(int32))
		} else if typeName == "uint32" {
			*valInt = int64(val.(uint32))
		}
		*valString = ""
	} else if typeDef == "int32" {
		typeName := reflect.TypeOf(val).Name()
		if typeName == "int64" {
			*valInt = val.(int64)
		} else if typeName == "int32" {
			*valInt = int64(val.(int32))
		} else if typeName == "uint32" {
			*valInt = int64(val.(uint32))
		}
		*valString = ""
	}
}

func main() {
	i32 := int32(1234)
	pTypeName(i32)

	ui32 := uint32(2345)
	pTypeName(ui32)

	i64 := int64(3242)
	pTypeName(i64)

	var idInt int64
	var idString string
	AttributesDefSafeTrans("int32", &idInt, &idString, i32)
	fmt.Println(idInt)
	fmt.Println(idString)
}
