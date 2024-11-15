package redis

import (
	"context"
	"fmt"

	"github.com/kaitokid2302/URL-Shortening-Service/internal/db"
)

func GetUrl(key string) string {
	ctx := context.Background()
	val, _ := Client.Get(ctx, key).Result()
	fmt.Printf("val: %v\n", val)
	var realUrl string
	if val == "" {
		realUrl, _ = db.FindByShortCode(key)

		Client.Set(ctx, key, realUrl, 0)
		return realUrl
	}
	return val
}
