package pokemon

import "github.com/labstack/echo"

type UseCase interface {
	GetPokemon(echo echo.Context) error
}
