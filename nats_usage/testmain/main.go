package main

import (
	"fmt"
	"nats_usage/common"
)

func main() {
	fmt.Println("main begin.")

	common.NewSafeList()

	fmt.Println("main end.")
}
