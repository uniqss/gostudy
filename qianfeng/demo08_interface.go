package main

import "fmt"

func main() {
	m1 := Mouse{"罗技小红"}
	fmt.Println(m1.name)

	f1 := FlashDisk{"闪迪64GB"}
	fmt.Println(f1.name)

	testInterface(m1)
	testInterface(f1)

	var usb USB
	usb = m1
	usb.start()
	usb.end()
	//fmt.Println(usb.name) // usb.name undefined (type USB has no field or method name)

	f1.deleteData()

	// 函数接受接口类型的参数，就是多态（作为返回值也一样）
	var arr [3]USB
	arr[0] = m1
	arr[1] = f1
	fmt.Println(arr)

	// 鸭子类型
}

type USB interface {
	start() // usb设备开始工作
	end()   // USB设备停止工作
}

type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "可以开始工作了，点点点。。。")
}

func (m Mouse) end() {
	fmt.Println(m.name, "结束工作了，可以安全退出")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "准备开始工作了，可以进行数据存储")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "可以弹出了")
}

func testInterface(usb USB) {
	usb.start()
	usb.end()
}

func (f FlashDisk) deleteData() {
	fmt.Println(f.name, "可以删除数据了")
}
