package main

import (
	"fmt"
	"log"
	"net"
)

func tryconnect(i int) {
	fmt.Println("tryconnect i:", i)
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	//defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	var b []byte
	n, err := conn.Read(b)
	if err != nil {
		log.Fatal("Read error:", err)
	}

	if n > 0 {
		fmt.Println(string(b))
	} else {
		fmt.Println("n:", n)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		go tryconnect(i)
	}
}
