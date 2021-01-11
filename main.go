package main

import (
	"fmt"

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
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	pokeAPIService := pokeApiService.NewPokemonService()
	pokeAPIHandler := handler.NewPokemonHandler(pokeAPIService)

	pokemonRepo := repository.NewPokemonsRepository()
	pokemonService := services.NewFavPokemonsService(pokemonRepo, pokeAPIService)
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
