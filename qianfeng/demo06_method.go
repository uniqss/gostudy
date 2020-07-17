package main

import "fmt"

func main() {
	w1 := Worker{"王二狗", 30, "男"}
	w1.work()

	w2 := Worker{name: "ruby", age: 34, gender: "女"}
	fmt.Printf("%T\n", w2)
	w2.work()

	w2.rest()
	w1.rest()

	w1.printInfo()

	c1 := Cat{"yellow", 3}
	c1.printInfo()
}

type Worker struct {
	name   string
	age    int
	gender string
}

type Cat struct {
	color string
	age int
}

func (w Worker) work() {
	fmt.Println(w.name, "在工作。。。")
}

func (p *Worker) rest() {
	fmt.Println(p.name, "在休息。。。")
}

func (p* Worker) printInfo(){
	fmt.Printf("工人 姓名:%s 年龄:%d 性别:%s\n", p.name, p.age, p.gender)
}

func (p *Cat) printInfo() {
	fmt.Printf("猫 颜色：%s 年龄:%d", p.color, p.age)

}
