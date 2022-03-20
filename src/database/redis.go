package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

func SetupCacheChannel() {
	CacheChannel = make(chan string)

	go func(ch chan string) {

		for {
			time.Sleep(1 * time.Second)
			key := <-ch
			fmt.Printf("cache cleared %s\n", key)
			Cache.Del(context.Background(), key)
		}

	}(CacheChannel)

}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key
	}

}
