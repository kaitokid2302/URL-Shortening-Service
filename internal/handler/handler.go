package handler

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/URL-Shortening-Service/internal/db"
	"github.com/kaitokid2302/URL-Shortening-Service/internal/redis"
)

func App() {
	app := gin.Default()
	app.POST("/shorten", HandlePostUrl)
	app.GET("/shorten/:shortCode", HandleGetUrl)
	app.PUT("/shorten/:shortCode", HandlePutUrl)
	app.DELETE("/shorten/:shortCode", HanldeDeleteUrl)
	app.GET("/shorten/:shortCode/stats", HandleGetCount)

	app.Run(":8080")
}

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)

	x := fmt.Sprintf("%x", bs)
	// get 10 characters

	return x[:8]
}

func HandlePostUrl(c *gin.Context) {
	fmt.Printf("\"lam\": %v\n", "lam")
	var f interface{}
	c.ShouldBindJSON(&f)
	postBody, ok := f.(map[string]interface{})

	if !ok {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	url, ok := postBody["url"].(string)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	insert := db.Url{
		Url:       url,
		ShortCode: Sha256(url),
		Count: &db.Count{
			Count: 0,
		},
	}

	db.Db.Create(&insert)
	/*
			"id": "1",
		"url": "https://www.example.com/some/long/url",
		"shortCode": "abc123",
		"createdAt": "2021-09-01T12:00:00Z",
		"updatedAt" : "2021-09-01T12:00:00Z"
	*/
	c.JSON(200, gin.H{"id": insert.ID, "url": insert.Url, "shortCode": insert.ShortCode, "createdAt": insert.CreatedAt, "updatedAt": insert.UpdatedAt})
}

func HandleGetUrl(c *gin.Context) {
	shortCode := c.Param("shortCode")
	x := redis.GetUrl(shortCode)

	var f map[string]interface{}
	json.Unmarshal([]byte(x), &f)
	c.JSON(200, f)
	redis.IncreaseCount(shortCode)
}

func HandlePutUrl(c *gin.Context) {
	shortCode := c.Param("shortCode")
	var f map[string]interface{}
	c.ShouldBindJSON(&f)

	if url, ok := f["url"].(string); ok {
		_, thing := db.FindByShortCode(shortCode)
		thing.Url = url
		db.Db.Save(thing)
		// delete from redis
		redis.ResetKey(shortCode)

		c.JSON(200, gin.H{"id": thing.ID, "url": thing.Url, "shortCode": thing.ShortCode, "createdAt": thing.CreatedAt, "updatedAt": thing.UpdatedAt})
		return
	}
	c.JSON(400, gin.H{"error": "Invalid request"})
}

func HanldeDeleteUrl(c *gin.Context) {
	shortCode := c.Param("shortCode")
	_, thing := db.FindByShortCode(shortCode)
	db.Db.Delete(thing)
	redis.ResetKey(shortCode)

	if thing.ID == 0 {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.JSON(204, gin.H{"message": "No content"})
}

func HandleGetCount(c *gin.Context) {
	shortCode := c.Param("shortCode")
	count := redis.GetCount(shortCode)
	thing := redis.GetUrl(shortCode)
	/*
			"id" : "1" ,
		"url": "https://www.example.com/some/long/url",
		"shortCode": "abc123",
		"createdAt": "2021-09-01T12:00 : 00Z",
		"updatedAt" : "2021-09-01T12:00:002" ,
		"accessCount": 10
	*/
	var f map[string]interface{}
	json.Unmarshal([]byte(thing), &f)
	c.JSON(200, gin.H{"id": f["id"], "url": f["url"], "shortCode": f["shortCode"], "createdAt": f["createdAt"], "updatedAt": f["updatedAt"], "accessCount": count})
}
