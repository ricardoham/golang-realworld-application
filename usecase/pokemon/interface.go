package pokemon

import "github.com/labstack/echo"

type UseCase interface {
	CreatePokemon(echo echo.Context) error
}
