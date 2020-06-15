package main

import (
	"fmt"
	"log"
	"net"
)

func tryconnect(i int) {
	fmt.Println("tryconnect i:", i)
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
		conn.Close()
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
	for i := 0; i < 10000; i++ {
		go tryconnect(i)
	}
}
