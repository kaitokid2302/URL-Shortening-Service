package main

import (
	"github.com/joho/godotenv"
	"github.com/kaitokid2302/URL-Shortening-Service/internal/cronjob"
	"github.com/kaitokid2302/URL-Shortening-Service/internal/handler"
)

func main() {
	godotenv.Load()
	go cronjob.Cronjob()

	handler.App()

}
