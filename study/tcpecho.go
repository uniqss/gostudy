package main

import (
	"fmt"
	"log"
	"net"
	"time"
)
//func Uint16ToBytes(n uint16) []byte {
//	return []byte{
//		byte(n),
//		byte(n >> 8),
//	}
//}

//func BytesToUint16(array []byte) uint16 {
//	var data uint16 =0
//	for i:=0;i< len(array);i++  {
//		data = data+uint16(uint(array[i])<<uint(8*i))
//	}
//
//	return data
//}

func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}

//func Uint64ToBytes(n uint64) []byte {
//	return []byte{
//		byte(n),
//		byte(n >> 8),
//		byte(n >> 16),
//		byte(n >> 24),
//		byte(n >> 32),
//		byte(n >> 40),
//		byte(n >> 48),
//		byte(n >> 56),
//	}
//}

func tcpEchoHandler(conn net.Conn) {
	fmt.Println("tcpEchoHandler new connection.")
	//fmt.Fprintf(conn, "%s", time.Now().String())
	b := make([]byte, 1024)
	conn.Read(b)
	fmt.Println(string(b))
	var msglen uint32
	msglen = uint32(len(b))
	conn.Write(Uint32ToBytes(msglen))
	conn.Write(b)
	time.Sleep(time.Second * 5)
	conn.Close()
}

func main() {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("new connection.")
		go tcpEchoHandler(conn)
	}
}
