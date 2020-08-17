package main

import (
	"fmt"
	"time"
)

var (
	ticker *time.Ticker = nil
	unixnano int64 = 0
)

func main() {

	ticker = time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	unixnano = time.Now().UnixNano()

	for {
		select {
		case <-ticker.C:
			onLoop()
		}
	}
}

func onLoop() {
	fmt.Println("onLoop time:", (time.Now().UnixNano() - unixnano) / 1000000)
	time.Sleep(99500*time.Microsecond)
}
