package limiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var luaScript = redis.NewScript(`
local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local requested = tonumber(ARGV[4])

local data = redis.call("HMGET", key, "tokens", "timestamp")
local tokens = tonumber(data[1])
local last = tonumber(data[2])

if tokens == nil then
  tokens = capacity
  last = now
end

local delta = math.max(0, now - last)
local refill = delta * refill_rate
tokens = math.min(capacity, tokens + refill)

if tokens < requested then
  return {0, tokens}
end

tokens = tokens - requested

redis.call("HMSET", key, "tokens", tokens, "timestamp", now)
redis.call("EXPIRE", key, 3600)

return {1, tokens}
`)

func AllowRequest(ctx context.Context, rdb *redis.Client, key string, capacity int, refillRate float64) (bool, int64, error) {
	now := time.Now().Unix()

	result, err := luaScript.Run(ctx, rdb, []string{key},
		capacity,
		refillRate,
		now,
		1,
	).Result()

	if err != nil {
		return false, 0, err
	}

	res := result.([]interface{})
	allowed := res[0].(int64)
	tokens := res[1].(int64)

	return allowed == 1, tokens, nil
}
