package favpokemon

import "github.com/labstack/echo"

type UseCase interface {
	CreateFavPokemon(echo echo.Context) error
	GetAllFavPokemons(echo echo.Context) error
	UpdateFavPokemon(echo echo.Context) error
	DeleteFavPokemon(echo echo.Context) error
}
