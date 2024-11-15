package redis

import (
	"context"

	"github.com/kaitokid2302/URL-Shortening-Service/internal/db"
)

func IncreaseCount(key string) {
	key = "count" + key
	Client.Incr(context.Background(), key)
}

func GetCount(key string) int {
	exist := Client.Exists(context.Background(), "count"+key).Val()
	if exist == 0 {
		_, url := db.FindByShortCode(key)
		Client.IncrBy(context.Background(), "count"+key, int64(url.Count.Count))
		return url.Count.Count
	}
	x, _ := Client.Get(context.Background(), "count"+key).Int64()
	return int(x)
}
