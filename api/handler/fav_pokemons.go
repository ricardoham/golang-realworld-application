package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/ricardoham/pokedex-api/api/presenter"
	services "github.com/ricardoham/pokedex-api/usecase/favpokemon"
	usecase "github.com/ricardoham/pokedex-api/usecase/favpokemon"
)

type favPokemonsHandler struct {
	pokemonService services.FavPokemon
}

func NewFavPokemonsHandler(services *services.FavPokemonService) usecase.UseCase {
	return &favPokemonsHandler{
		pokemonService: services,
	}
}

func (p *favPokemonsHandler) CreateFavPokemon(echo echo.Context) error {
	var pokemon presenter.SaveFavPokemon
	if err := echo.Bind(&pokemon); err != nil {
		log.Fatal("Error when binding the content ", err)
		return err
	}

	err := p.pokemonService.CreateFavPokemon(&pokemon)
	if err != nil {
		log.Fatal("Error ", err)
	}

	return echo.JSON(http.StatusCreated, "created")
}

func (p *favPokemonsHandler) GetAllFavPokemons(echo echo.Context) error {
	pokemons, err := p.pokemonService.GetAllFavPokemons()
	if err != nil {
		log.Fatal("Error during fetch the data ", err)
		return echo.JSON(http.StatusBadRequest, "error")
	}

	return echo.JSON(http.StatusOK, pokemons)
}

func (p *favPokemonsHandler) UpdateFavPokemon(echo echo.Context) error {
	pokeId := uuid.MustParse(echo.Param("id"))
	updatePokemon := new(presenter.FavPokemon)
	if err := echo.Bind(updatePokemon); err != nil {
		log.Fatal("Error when binding the content ", err)
		return err
	}

	result, err := p.pokemonService.UpdateFavPokemon(pokeId, updatePokemon)
	if err != nil {
		log.Fatal("Error when update data ", err)
		return echo.NoContent(http.StatusBadRequest)
	}

	return echo.JSON(http.StatusOK, result)
}

func (p *favPokemonsHandler) DeleteFavPokemon(echo echo.Context) error {
	var pokeId presenter.DeleteFavPokemon
	if err := echo.Bind(&pokeId); err != nil {
		log.Fatal("Error when binding the content ", err)
		return err
	}

	result, err := p.pokemonService.DeleteFavPokemon(pokeId)
	if err != nil {
		log.Fatal("Error when deleting data ", err)
		return echo.NoContent(http.StatusBadRequest)
	}

	return echo.JSON(http.StatusOK, result)
}
