package client

import (
	"net/http"

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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
