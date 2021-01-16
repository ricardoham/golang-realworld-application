package pokemon

import (
	"github.com/labstack/echo"
	"github.com/ricardoham/pokedex-api/api/presenter"
)

type UseCase interface {
	GetPokemon(echo echo.Context) error
}

type Pokemon interface {
	GetPokemonFromPokeApi(pokemon string) (*presenter.Pokemon, error)
	GetAllResultPokemonFromPokeApi() (*presenter.Result, error)
}
