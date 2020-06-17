package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("自己创建玩的。。。。")
	fmt.Println(err1)

	fmt.Printf("%T\n", err1)

	err2 := fmt.Errorf("错误的信息码：%d", 100)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)

	err3 := checkAge(-1)
	//err3 := checkAge(30)
	if err3 != nil {
		fmt.Println(err3)
		return
	}

}

func checkAge(age int) error {
	if age < 0 {
		//return errors.New("年龄不合法")
		err := fmt.Errorf("您给定的年龄【%d】不合法", age)
		return err
	}
	fmt.Println("年龄是：", age)
	return nil
}
