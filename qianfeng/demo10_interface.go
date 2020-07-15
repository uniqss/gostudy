package main

import (
	"fmt"
	"math"
)

func main() {
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)

	var c1 Circle = Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.radius)

	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var s2 Shape
	s2 = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())

	testShape(t1)
	testShape(c1)
	testShape(s1)

	getType(t1)
	getType(c1)
	getType(s1)

	var t2 *Triangle = &Triangle{3, 4, 2}
	fmt.Printf("t2:%T, %p, %p\n", t2, &t2, t2)
	getType(t2)
	getType(*t2)
	getType2(t2)
	getType2(t1)
}

func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形。。", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形。。", ins.radius)
	case *Triangle:
		fmt.Println("三角形指针:", ins.a, ins.b, ins.c)
	}
}

func getType(s Shape) {
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Println("是三角形指针类型，三边是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径：", ins.radius)
	} else if ins, ok := s.(*Circle); ok {
		fmt.Println("是圆形指针类型，半径：", ins.radius)
	} else {
		fmt.Println("不知道的类型")
	}
}

type Shape interface {
	peri() float64 // 周长
	area() float64 // 面积
}

type Triangle struct {
	//a float64
	//b float64
	//c float64
	a, b, c float64
}

func testShape(s Shape) {
	fmt.Printf("周长：%.2f, 面积:%.2f", s.peri(), s.area())
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
