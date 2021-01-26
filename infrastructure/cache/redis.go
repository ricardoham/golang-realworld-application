package cache

import (
	"encoding/json"
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

func (c *Cache) Get(key string, data interface{}) error {
	serialize, err := c.redisClient.Get(key).Result()
	if err == redis.Nil {
		log.Println("The key may not exist, ", err)
		return err
	}

	err = json.Unmarshal([]byte(serialize), &data)
	if err != nil {
		return err
	}

	log.Println("Key Value, ", data)
	return nil
}

func (c *Cache) Set(key string, value interface{}, expiresOn int) (bool, error) {
	serialize, err := json.Marshal(value)
	if err != nil {
		log.Println("Failed to set new cache on Redis", err)
	}
	err = c.redisClient.Set(key, string(serialize), time.Duration(expiresOn)*time.Second).Err()

	return true, err
}

func (c *Cache) Delete(key string) error {
	err := c.redisClient.Del(key).Err()
	if err == redis.Nil {
		log.Println("The key may not exist, ", err)
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
