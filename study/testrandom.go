package main

import (
	"fmt"
	"github.com/seehuhn/mt19937"
	"math/rand"
	"time"
)

func main() {
	var randArr []int32
	_ = randArr

	var repeateCheck = make(map[int32]bool)

	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	for i:=0;i < 10000000000;i++ {
		//r := rand.Int31()
		r := rng.Int31()
		if repeateCheck[r] {
			fmt.Println("repeated i:", i)
			break
		}
		repeateCheck[r] = true
	}

	//for i := 0; i < 1000000; i++ {
	//	r := rand.Int31()
	//	randArr = append(randArr, r)
	//	repeateCheck[]
	//}

}
