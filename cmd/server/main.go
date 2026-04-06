package main

import (
	"log"
	"os"

	"rate-limiter/internal/config"
	"rate-limiter/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	rdb := config.NewRedisClient()

	r := router.SetupRouter(rdb)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(r.Run(":" + port))
}
