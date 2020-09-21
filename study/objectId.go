package main

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"github.com/globalsign/mgo/bson"
)

func readMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	fmt.Println("hostname:", hostname)
	if err1 != nil {
		_, err2 := io.ReadFull(rand.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}

func main() {
	mid := readMachineId()
	fmt.Println(mid)

	oid := bson.NewObjectId()
	fmt.Println(oid.Hex())
	fmt.Println(string(oid))

}
