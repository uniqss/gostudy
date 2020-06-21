package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10 // 100张票

var mutex sync.Mutex
var wg sync.WaitGroup
func main() {
	wg.Add(4)

	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait()
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "售完，没有票了")
			break
		}
		mutex.Unlock()
	}
}
