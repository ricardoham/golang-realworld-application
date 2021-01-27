package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	services "github.com/ricardoham/pokedex-api/usecase/client"
	usecase "github.com/ricardoham/pokedex-api/usecase/client"
)

type clientPokemonHandler struct {
	pokemonService *services.PokemonService
}

func NewClientPokemonHandler(service *services.PokemonService) usecase.UseCase {
	return &clientPokemonHandler{
		pokemonService: service,
	}
}

func (p *clientPokemonHandler) GetPokemon(echo echo.Context) error {
	var result interface{}
	var err error
	input := echo.Param("*")

	if input == "" {
		result, err = p.pokemonService.GetAllResultPokemonFromPokeApi()
		if err != nil {
			log.Fatal("Error during fetch pokemons, ", err)
			return echo.JSON(http.StatusBadRequest, err)
		}
	} else {
		result, err = p.pokemonService.GetPokemonFromPokeApi(input)
		if err != nil {
			log.Fatal("Error during fetch pokemon, ", err)
			return echo.JSON(http.StatusBadRequest, err)
		}
	}

	return echo.JSON(http.StatusOK, result)
}
