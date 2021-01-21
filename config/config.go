package config

import (
	"os"
	"strconv"
	"sync"
	"time"
)

type APIConfig struct {
	HostPort int
}

type MongoConfig struct {
	Address      string
	DatabaseName string
	Timeout      time.Duration
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

var apiConfig *APIConfig
var onceConfig sync.Once

// Return the instance of APIConfig
func GetAPIConfig() *APIConfig {
	onceConfig.Do(func() {
		apiConfig = &APIConfig{
			HostPort: getAPIHostPort(),
		}
	})

	return apiConfig
}

func getAPIHostPort() int {
	apiPort, err := strconv.Atoi(os.Getenv("HOST_PORT"))
	if err != nil {
		apiPort = 8080
	}

	return apiPort
}

func GetMongoConfig() *MongoConfig {
	mongoTimeout := 10 * time.Second

	if os.Getenv("CUR_ENV") == "" {
		return &MongoConfig{
			Address:      MONGODB_ADDRS,
			DatabaseName: MONGODB_NAME,
			Timeout:      mongoTimeout,
		}
	}

	return &MongoConfig{
		Address:      os.Getenv("MONGODB_ADDRS"),
		DatabaseName: os.Getenv("MONGODB_NAME"),
		Timeout:      mongoTimeout,
	}
}

func GetRedisConfig() *RedisConfig {
	if os.Getenv("CUR_ENV") == "" {
		return &RedisConfig{
			Addr:     REDIS_ADDRS,
			Password: REDIS_PASSWORD,
			DB:       REDIS_DB,
		}
	}

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		db = 0
	}

	return &RedisConfig{
		Addr:     os.Getenv("REDIS_ADDRS"),
		Password: os.Getenv("MONGODB_NAME"),
		DB:       db,
	}
}
