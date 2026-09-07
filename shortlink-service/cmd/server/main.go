package main

import (
	"shortlink-service/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	app.Run()
}
