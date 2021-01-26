package main

import (
	"fmt"
	"log"

	"github.com/ricardoham/pokedex-api/infrastructure/cache"

	"github.com/go-redis/redis"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ricardoham/pokedex-api/api/handler"
	"github.com/ricardoham/pokedex-api/config"
	repository "github.com/ricardoham/pokedex-api/infrastructure/repository"
	services "github.com/ricardoham/pokedex-api/usecase/favpokemon"
	pokeApiService "github.com/ricardoham/pokedex-api/usecase/pokemon"
)

func main() {
	apiConfig := config.GetAPIConfig()
	redisConfig := config.GetRedisConfig()
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	redisClientConfig, err := redisConfiguration(redisConfig.Addr, redisConfig.Password, redisConfig.DB)
	if err != nil {
		log.Fatal("Error ocurred with Redis Client ", err)
	}

	cache := cache.NewRedisCache(redisClientConfig)

	pokeAPIService := pokeApiService.NewPokemonService()
	pokeAPIHandler := handler.NewPokemonHandler(pokeAPIService)

	pokemonRepo := repository.NewPokemonsRepository()
	pokemonService := services.NewFavPokemonsService(pokemonRepo, cache, pokeAPIService)
	pokemonHandler := handler.NewFavPokemonsHandler(pokemonService)

	echo.GET("/v1/pokemons/*", pokeAPIHandler.GetPokemon)

	echoGroup := echo.Group("/v1/favpokemons")
	echoGroup.POST("", pokemonHandler.CreateFavPokemon)
	echoGroup.GET("", pokemonHandler.GetAllFavPokemons)
	echoGroup.PUT("/:id", pokemonHandler.UpdateFavPokemon)
	echoGroup.DELETE("", pokemonHandler.DeleteFavPokemon)

	port := fmt.Sprintf(":%d", apiConfig.HostPort)
	echo.Logger.Fatal(echo.Start(port))
}

func redisConfiguration(host string, password string, db int) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
		PoolSize: 50000,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	log.Println("Establish connection with Redis Client.")

	return redisClient, nil
}
