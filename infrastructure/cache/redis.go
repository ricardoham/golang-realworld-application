package cache

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

type Cache struct {
	redisClient *redis.Client
}

func NewRedisCache(redisClientConfig *redis.Client) *Cache {
	return &Cache{
		redisClient: redisClientConfig,
	}
}

func (c *Cache) Get(key string) (string, error) {
	cacherKey, err := c.redisClient.Get(key).Result()
	if err != nil {
		return "", err
	}

	if err == redis.Nil {
		log.Fatal("The key may not exist, ", err)
		return "", err
	}

	log.Println("Key Value, ", cacherKey)
	return cacherKey, nil
}

func (c *Cache) Set(key string, value string, expiresOn int) (bool, error) {
	err := c.redisClient.Set(key, value, time.Duration(expiresOn)*time.Second).Err()

	return true, err
}

func (c *Cache) Delete(key string) error {
	err := c.redisClient.Del(key).Err()
	if err == redis.Nil {
		log.Fatal("The key may not exist, ", err)
		return err
	}

	return err
}

func (c *Cache) Ping() (string, error) {
	pong, err := c.redisClient.Ping().Result()
	if err != nil {
		return "", err
	}

	return pong, nil
}
