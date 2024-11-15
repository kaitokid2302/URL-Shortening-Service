package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Url       string `gorm:"index"`
	ShortCode string `gorm:"index"`
	Count     *Count
}

type Count struct {
	gorm.Model
	Count int
	UrlID uint `gorm:"index"` // Foreign key
}

var Db *gorm.DB

func init() {
	godotenv.Load()
	dbNetwork := os.Getenv("db_network")
	dsn := fmt.Sprintf("host=%s user=gorm password=gorm dbname=gorm port=5432", dbNetwork)
	var er error
	Db, er = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// FullSaveAssociations: true,
		// Logger:               logger.Default.LogMode(logger.Info),
	})
	if er != nil {
		panic(er)
	}
	// Db.Migrator().DropTable(&Count{})
	// Db.Migrator().DropTable(&Url{})

	if err := Db.AutoMigrate(&Url{}, &Count{}); err != nil {
		fmt.Printf("Migration error: %v\n", err)
	}
}
func FindByShortCode(key string) (string, *Url) {
	fmt.Printf("key: %v\n", key)
	var url Url
	Db.Preload("Count").Where("short_code = ?", key).First(&url)
	// to json, except count

	var f map[string]interface{} = map[string]interface{}{
		"id":        url.ID,
		"url":       url.Url,
		"shortCode": url.ShortCode,
		"createdAt": url.CreatedAt,
		"updatedAt": url.UpdatedAt,
	}

	jsonData, _ := json.Marshal(f)
	return string(jsonData), &url
}
