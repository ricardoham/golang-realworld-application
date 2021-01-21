package cache

import "github.com/go-redis/redis"

type Cache struct {
	redisClient *redis.Client
}

func NewRedisCache(redisClientConfig *redis.Client) *Cache {
	return &Cache{
		redisClient: redisClientConfig,
	}
}
