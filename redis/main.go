package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	addr := flag.String("addr", "127.0.0.1:6379", "redis host and port")
	flag.Parse()

	client = redis.NewClient(&redis.Options{
		Addr: *addr,
	})
	fmt.Println("connected to", client)

	// push 10 items onto test key
	for i := 0; i < 10; i++ {
		lpush(client, "test", strconv.Itoa(i))
	}

	// read the first 5 items from the test key
	items := lrange(client, "test", 0, 5)
	fmt.Println(items)

	// delete the list
	del(client, "test")
}

func lpush(client *redis.Client, key string, value string) {
	client.LPush(key, value)
}

func lrange(client *redis.Client, key string, start int64, end int64) []string {
	items, err := client.LRange(key, start, end).Result()
	if err != nil {
		panic(err)
	}
	return items
}

func del(client *redis.Client, key string) {
	client.Del(key)
}
