package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ricardoham/pokedex-api/api/handler"
	"github.com/ricardoham/pokedex-api/config"
	repository "github.com/ricardoham/pokedex-api/infrastructure/repository"
	services "github.com/ricardoham/pokedex-api/usecase/favpokemon"
)

func main() {
	apiConfig := config.GetAPIConfig()
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	pokemonRepo := repository.NewPokemonsRepository()
	pokemonService := services.NewFavPokemonsService(pokemonRepo)
	pokemonHandler := handler.NewFavPokemonsHandler(pokemonService)

	echoGroup := echo.Group("/v1/pokemons")
	echoGroup.POST("", pokemonHandler.CreateFavPokemon)
	echoGroup.GET("", pokemonHandler.GetAllFavPokemons)
	echoGroup.PUT("/:id", pokemonHandler.UpdateFavPokemon)
	echoGroup.DELETE("", pokemonHandler.DeleteFavPokemon)

	port := fmt.Sprintf(":%d", apiConfig.HostPort)
	echo.Logger.Fatal(echo.Start(port))
}
