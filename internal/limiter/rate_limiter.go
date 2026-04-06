package limiter

import (
	"fmt"
	"net/http"

	"rate-limiter/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimiter(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "rate:" + c.ClientIP()

		capacity := 100
		refillRate := 1.0

		allowed, tokens, err := AllowRequest(config.Ctx, rdb, key, capacity, refillRate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
			c.Abort()
			return
		}

		//headers
		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", capacity))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", tokens))
		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
