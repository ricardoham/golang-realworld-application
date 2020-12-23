package main

import (
	"fmt"

	"github.com/ricardoham/pokedex-api/services"

	handler "github.com/ricardoham/pokedex-api/api/handler"
	"github.com/ricardoham/pokedex-api/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ricardoham/pokedex-api/config"
)

func main() {
	apiConfig := config.GetAPIConfig()
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	pokemonRepo := repository.NewPokemonsRepository()
	pokemonService := services.NewPokemonsService(pokemonRepo)
	pokemonHandler := handler.NewPokemonsHandler(pokemonService)

	fmt.Println(pokemonHandler) // TODO

	port := fmt.Sprintf(":%d", apiConfig.HostPort)
	echo.Logger.Fatal(echo.Start(port))
}
