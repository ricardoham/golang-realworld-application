package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ricardoham/pokedex-api/api/handler"
	"github.com/ricardoham/pokedex-api/config"
	repository "github.com/ricardoham/pokedex-api/infrastructure/repository"
	services "github.com/ricardoham/pokedex-api/usecase/pokemon"
)

func main() {
	apiConfig := config.GetAPIConfig()
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	pokemonRepo := repository.NewPokemonsRepository()
	pokemonService := services.NewPokemonsService(pokemonRepo)
	pokemonHandler := handler.NewPokemonsHandler(pokemonService)

	echoGroup := echo.Group("/v1/pokemons")
	echoGroup.POST("", pokemonHandler.CreatePokemon)
	echoGroup.GET("", pokemonHandler.GetAllPokemons)
	echoGroup.PUT("/:id", pokemonHandler.UpdatePokemon)
	echoGroup.DELETE("", pokemonHandler.DeletePokemon)

	port := fmt.Sprintf(":%d", apiConfig.HostPort)
	echo.Logger.Fatal(echo.Start(port))
}
