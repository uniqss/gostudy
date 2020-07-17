package main

import (
	"context"
	"fmt"
	"redis_usage/redis_wrapper"
)

func main() {
	fmt.Println("hello world")
	var rc = redis_wrapper.GetRedisConnClient("asdf")
	ctx := context.Background()
	k := "asdf"
	v := rc.Get(ctx, k)
	fmt.Println(v)
	rc.Set(ctx, k, "qoweihpghigoifdjas", 0)
	v = rc.Get(ctx, k)
	fmt.Println(v)
	rc.Del(ctx, k)
}
