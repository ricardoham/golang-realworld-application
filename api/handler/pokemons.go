package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ricardoham/pokedex-api/entity"
	services "github.com/ricardoham/pokedex-api/usecase/pokemon"
	usecase "github.com/ricardoham/pokedex-api/usecase/pokemon"
)

type pokemonsHandler struct {
	pokemonService *services.PokemonService
}

func NewPokemonsHandler(services *services.PokemonService) usecase.UseCase {
	return &pokemonsHandler{
		pokemonService: services,
	}
}

func (p *pokemonsHandler) CreatePokemon(echo echo.Context) error {
	var pokemon entity.Pokemon
	if err := echo.Bind(&pokemon); err != nil {
		log.Fatal("Error when binding the content ", err)
		return err
	}

	err := p.pokemonService.CreatePokemon(&pokemon)
	if err != nil {
		log.Fatal("Error ", err)
	}

	return echo.JSON(http.StatusCreated, "created")
}
