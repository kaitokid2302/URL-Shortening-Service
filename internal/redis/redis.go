package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func init() {
	redisNetwork := os.Getenv("redis_network")
	// Create a new Redis client
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", redisNetwork), // Redis server address
		Password: "",                                   // no password set
		DB:       0,                                    // default DB
	})

	ctx := context.Background()
	_, er := Client.Ping(ctx).Result()
	if er != nil {
		panic(er)
	}
}

func ResetKey(key string) {
	ctx := context.Background()
	Client.Del(ctx, key)
	Client.Del(ctx, "count"+key)
}
