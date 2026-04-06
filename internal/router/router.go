package router

import (
	"rate-limiter/internal/handlers"
	"rate-limiter/internal/limiter"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetupRouter(rdb *redis.Client) *gin.Engine {
	r := gin.Default()

	r.Use(limiter.RateLimiter(rdb))
	r.GET("/", handlers.TestHandler)

	return r
}
