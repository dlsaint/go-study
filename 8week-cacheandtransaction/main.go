package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
)

const N = 100000 // 插入十万条数据
const B = 200 //字节

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.0.77:6379",
		DB:   1,
	})
	client.FlushDB(ctx)

	// pipeline 批量插入数据
	pipeline := client.Pipeline()
	for i := range [N]struct{}{} {
		// 每条数据 value 大小为 100b
		pipeline.Set(ctx, fmt.Sprintf("test_key:%d", i), GenerateValueWithBytes(B), 0)
	}
	if _, err := pipeline.Exec(ctx); err != nil {
		panic(err)
	}

	fmt.Println("done")
}

// 生成指定 byte 大小的随机字符串
func GenerateValueWithBytes(size int) string {
	var res []byte
	str := "ABCDEFGHIJKLMNOPQRSTUVWSYZ0123456789"
	for i := 0; i < size; i++ {
		res = append(res, str[rand.Intn(len(str))])
	}
	return string(res)
}
