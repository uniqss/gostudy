package main

import "fmt"

type ITimer interface {
	OnTimer()
}

type TimerKey struct {
	key     *ITimer
	timerId int
}

var g_map map[TimerKey]string = map[TimerKey]string{}

func main() {

	k1 := TimerKey{
		key:     nil,
		timerId: 1234,
	}
	g_map[k1] = "sadf"

	k1.timerId = 4578
	g_map[k1] = "jjjjjj"

	k2 := TimerKey{
		key:     nil,
		timerId: 1234,
	}

	fmt.Println(g_map)

	fmt.Println(g_map[k2])
}
