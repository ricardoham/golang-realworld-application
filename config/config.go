package config

import (
	"fmt"
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
	fmt.Println("Apiport", apiPort)
	if err != nil {
		apiPort = 8080
	}

	return apiPort
}

func GetMongoConfig() *MongoConfig {
	mongoAddress := os.Getenv("MONGODB_ADDRS")
	mongoDatabaseName := os.Getenv("MONGODB_NAME")
	mongoTimeout := 10 * time.Second

	return &MongoConfig{
		Address:      mongoAddress,
		DatabaseName: mongoDatabaseName,
		Timeout:      mongoTimeout,
	}
}
