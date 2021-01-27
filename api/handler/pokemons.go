package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/ricardoham/pokedex-api/api/presenter"
	services "github.com/ricardoham/pokedex-api/usecase/pokemon"
	usecase "github.com/ricardoham/pokedex-api/usecase/pokemon"
)

type pokemonsHandler struct {
	pokemonService services.Pokemon
}

func NewPokemonsHandler(services *services.PokemonService) usecase.UseCase {
	return &pokemonsHandler{
		pokemonService: services,
	}
}

func (p *pokemonsHandler) CreatePokemon(echo echo.Context) error {
	var pokemon presenter.SavePokemon
	if err := echo.Bind(&pokemon); err != nil {
		log.Fatal("Error when binding the content ", err)
		return err
	}

	err := p.pokemonService.CreatePokemon(&pokemon)
	if err != nil {
		log.Println("Error ", err)
		return echo.JSON(http.StatusBadRequest, "Couldn't create Pokemon")
	}

	return echo.JSON(http.StatusCreated, "created")
}

func (p *pokemonsHandler) GetPokemon(echo echo.Context) error {
	pokeID := uuid.MustParse(echo.Param("id"))

	result, err := p.pokemonService.GetPokemon(pokeID)
	if err != nil {
		log.Println("Error during fetch data ", err)
		return echo.JSON(http.StatusBadRequest, "error during fetch data")
	}

	return echo.JSON(http.StatusOK, result)
}

func (p *pokemonsHandler) GetAllPokemons(echo echo.Context) error {
	pokemons, err := p.pokemonService.GetAllPokemons()
	if err != nil {
		log.Fatal("Error during fetch the data ", err)
		return echo.JSON(http.StatusBadRequest, "error")
	}

	return echo.JSON(http.StatusOK, pokemons)
}

func (p *pokemonsHandler) UpdatePokemon(echo echo.Context) error {
	pokeId := uuid.MustParse(echo.Param("id"))
	updatePokemon := new(presenter.Pokemon)
	if err := echo.Bind(updatePokemon); err != nil {
		log.Fatal("Error when binding the content ", err)
		return err
	}

	result, err := p.pokemonService.UpdatePokemon(pokeId, updatePokemon)
	if err != nil {
		log.Fatal("Error when update data ", err)
		return echo.NoContent(http.StatusBadRequest)
	}

	return echo.JSON(http.StatusOK, result)
}

func (p *pokemonsHandler) DeletePokemon(echo echo.Context) error {
	var pokeId presenter.DeletePokemon
	if err := echo.Bind(&pokeId); err != nil {
		log.Fatal("Error when binding the content ", err)
		return err
	}

	result, err := p.pokemonService.DeletePokemon(pokeId)
	if err != nil {
		log.Fatal("Error when deleting data ", err)
		return echo.NoContent(http.StatusBadRequest)
	}

	return echo.JSON(http.StatusOK, result)
}
