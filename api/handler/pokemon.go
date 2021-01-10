package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	services "github.com/ricardoham/pokedex-api/usecase/pokemon"
	usecase "github.com/ricardoham/pokedex-api/usecase/pokemon"
)

type pokemonHandler struct {
	pokemonService *services.PokemonService
}

func NewPokemonHandler(service *services.PokemonService) usecase.UseCase {
	return &pokemonHandler{
		pokemonService: service,
	}
}

func (p *pokemonHandler) GetPokemon(echo echo.Context) error {
	input := echo.Param("*")
	result, err := p.pokemonService.GetPokemonFromPokeApi(input)
	if err != nil {
		log.Fatal("Error during fetch pokemon, ", err)
		return echo.JSON(http.StatusBadRequest, err)
	}

	return echo.JSON(http.StatusOK, result)
}
