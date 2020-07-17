package main

import "fmt"

func main() {
	length, width := -1.1, -2.2
	area, err := rectArea(length, width)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*areaError);ok {
			if err.length < 0 {
				fmt.Printf("err:长度：%.2f,小于0\n", err.length)
			}
			if err.width < 0 {
				fmt.Printf("err:宽度：%.2f,小于0\n", err.width)
			}
		}
		return
	}
	fmt.Println("矩形的面积是：", area)
}

type areaError struct {
	msg    string
	length float64
	width  float64
}

func (e *areaError) Error() string {
	return e.msg
}
func (e *areaError) lengthNegative() bool {
	return e.length < 0
}
func (e *areaError) widthNegative() bool {
	return e.width < 0
}
func rectArea(length, width float64) (float64, error) {
	msg := ""
	if length < 0 {
		msg = "长度小于0"
	}
	if width < 0 {
		if msg == "" {
			msg = "宽度小于0"
		} else {
			msg += ",宽度也小于0"
		}
	}

	if msg != "" {
		return 0, &areaError{msg, length, width}
	}
	return length * width, nil
}
