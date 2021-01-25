package main

import (
	"log"
	"reflect"
)

type DT_Rift_Config struct {
	INT uint32
	STRING string
	ArrINT []uint32
	ArrSTRING []string
}

func (dt *DT_Rift_Config) Fuck(){
	log.Println("Fuck")
}

func (dt *DT_Rift_Config) Fuck2(data map[string]interface{}){
	log.Println("Fuck2 data:", data)
}

func main(){
	var fakeArrInt []uint32
	var fakeArrString []string

	var baseStruct interface{}
	baseStruct = &DT_Rift_Config{}

	base := reflect.New(reflect.TypeOf(baseStruct).Elem()).Elem()

	var args []interface{}
	for i := 0;i < base.NumField();i++{
		baseField := base.Field(i)
		baseFieldAddr := baseField.Addr()
		baseInterface := baseFieldAddr.Interface()

		if baseField.Kind() == reflect.Slice {
			if baseField.Type() == reflect.TypeOf(fakeArrInt) {
				log.Println("11111")
			} else if baseField.Type() == reflect.TypeOf(fakeArrString) {
				log.Println("22222")
			}
			args = append(args, baseInterface)
		} else {
			args = append(args, baseInterface)
		}
	}

	//reflect.ValueOf(&baseStruct).MethodByName("Geeks").Call([]reflect.Value{})

	v1 := reflect.ValueOf(&baseStruct)
	v2 := reflect.ValueOf(baseStruct)

	var o DT_Rift_Config

	v3 := reflect.ValueOf(o)
	v4 := reflect.ValueOf(&o)

	log.Println("v1:", v1, " v2:", v2, " v3:", v3, " v4:", v4)

	//v1.MethodByName("Fuck").Call([]reflect.Value{})
	v2.MethodByName("Fuck").Call([]reflect.Value{})
	//v3.MethodByName("Fuck").Call([]reflect.Value{})
	v4.MethodByName("Fuck").Call([]reflect.Value{})

	reflect.ValueOf(baseStruct).MethodByName("Fuck").Call([]reflect.Value{})

	reflect.ValueOf(baseStruct).MethodByName("Fuck").Call([]reflect.Value{})

	log.Println("args:", args)
}
