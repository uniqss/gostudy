package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	ctx = context.Background()
)

func main() {
	fmt.Print("etcd test main begin......")
	if true {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
			DialTimeout: 5 * time.Second,
		})
		defer cli.Close()

		if err != nil {
			// handle error!
			fmt.Println("err:", err)
		}
		key := "aaa"
		value := "bbb"

		r, err := cli.Get(ctx, key)
		if err != nil {
			fmt.Println("cli.Get err:", err)
			return
		}
		if len(r.Kvs) != 1 {
			fmt.Println("cli.Get count error. count:", len(r.Kvs))
		}
		fmt.Println(string(r.Kvs[0].Value))

		_, err = cli.Put(ctx, key, value)
		if err != nil {
			fmt.Println("cli.Put err:", err)
		}

		iiiii := 0 == 0
		if iiiii {
			r, err = cli.Get(ctx, key)
			if err != nil {
				fmt.Println("cli.Get err:", err)
				return
			}
			if len(r.Kvs) != 1 {
				fmt.Println("cli.Get count error. count:", len(r.Kvs))
			}
			fmt.Println(string(r.Kvs[0].Value))
		}

	}
	fmt.Print("etcd test main end......")
}
