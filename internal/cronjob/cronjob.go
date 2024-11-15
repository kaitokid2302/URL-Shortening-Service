package cronjob

import (
	"fmt"
	"time"

	"github.com/kaitokid2302/URL-Shortening-Service/internal/db"
	"github.com/kaitokid2302/URL-Shortening-Service/internal/redis"
	"gorm.io/gorm"
)

func Cronjob() {
	// Implement cronjob here
	id := 0
	for {
		fmt.Printf("\"lam\": %v\n", "lam")
		Db := db.Db

		// search by cursor, limit 1000

		var urls []db.Url
		// cursor pagination, order by id
		Db.Preload("Count").Where("id > ?", id).Order("id").Limit(1000).Find(&urls)
		if len(urls) == 0 {
			id = 0
		}
		fmt.Printf("urls: %v\n", urls)

		for i := range urls {
			x := redis.GetCount(urls[i].ShortCode)
			if x > urls[i].Count.Count {
				// update count
				urls[i].Count.Count = x
				// Db.Save(&urls[i])
			}
			id = int(urls[i].ID)
		}
		// Db.Save(&urls)
		// save full association

		Db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&urls)
		time.Sleep(60 * time.Second)
	}
}
