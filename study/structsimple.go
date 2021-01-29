package main

// 外部接口
type DT_Player struct {
	Id     int32
	Name   string
	Age    uint32
	Height uint32
}

// 内部使用，不要暴出去
type _DT_Player_Key2 struct {
	Id   int32
	Name string
}

type _DT_Player_Key3 struct {
	Id   int32
	Name string
	Age  uint32
}

// 万一暴的接口不能满足外面的需求，从这里拿全部数据，但是注意不要修改！！！ sb go不能const限定
var DT_Player_map = make(map[int32]*DT_Player)
// 内部使用，不要暴出去
var _DT_Player_map2 = make(map[_DT_Player_Key2]*DT_Player)
var _DT_Player_map3 = make(map[_DT_Player_Key3]*DT_Player)

// 内部使用，不要暴出去
func _AAAAddItem(item *DT_Player) {
	DT_Player_map[item.Id] = item
	_DT_Player_map2[_DT_Player_Key2{item.Id, item.Name}] = item
	_DT_Player_map3[_DT_Player_Key3{item.Id, item.Name, item.Age}] = item
}

// 外部接口
func LoadConfig () {
	var a1 = new(DT_Player)
	a1.Id = 123
	a1.Name = "张三"
	a1.Age = 18
	_AAAAddItem(a1)

	var a2 = new(DT_Player)
	a2.Id = 456
	a2.Name = "李四"
	a2.Age = 17
	_AAAAddItem(a2)
}

// 外部接口
func AAAFindByKey1(id int32) *DT_Player {
	return DT_Player_map[id]
}
func AAAFindByKey2(id int32, name string) *DT_Player {
	return _DT_Player_map2[_DT_Player_Key2{id, name}]
}
func AAAFindByKey3(id int32, name string, age uint32) *DT_Player {
	return _DT_Player_map3[_DT_Player_Key3{id, name, age}]
}

func main() {
	LoadConfig()

	// ...

	AAAFindByKey1(123)
	AAAFindByKey2(123, "张三")
	AAAFindByKey3(123, "张三", 18)
}
