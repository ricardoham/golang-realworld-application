package pokemon

import "github.com/labstack/echo"

type UseCase interface {
	CreatePokemon(echo echo.Context) error
	GetAllPokemons(echo echo.Context) error
	UpdatePokemon(echo echo.Context) error
	DeletePokemon(echo echo.Context) error
}
