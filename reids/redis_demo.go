package main

import (
	"context"
	"fmt"
	"log"


)
import "github.com/go-redis/redis/v8"

func main() {
	// 创建一个Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// 设置键值对
	ctx := context.Background()
	err := rdb.Set(ctx, "goKey", "goValue", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// 获取键值对
	val, err := rdb.Get(ctx, "goKey").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key", val)
	// Output: key value

	// 增加计数器
	counter, err := rdb.Incr(ctx, "counter").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("counter", counter)
	// Output: counter 1

	// 列出所有的键
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("keys", keys)
	// Output: keys [key counter]
}