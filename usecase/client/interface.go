package client

import (
	"github.com/labstack/echo"
	"github.com/ricardoham/pokedex-api/api/presenter"
)

type UseCase interface {
	GetPokemon(echo echo.Context) error
}

type ClientPokemon interface {
	GetPokemonFromPokeApi(pokemon string) (*presenter.ClientPokemon, error)
	GetAllResultPokemonFromPokeApi() (*presenter.Result, error)
}
