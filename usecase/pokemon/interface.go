package pokemon

import (
	"context"

	"github.com/labstack/echo"
	"github.com/ricardoham/pokedex-api/api/presenter"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pokemon interface {
	CreatePokemon(pokemon *presenter.SavePokemon) error
	GetPokemon(pokeID string) (presenter.Pokemon, error)
	GetAllPokemons() ([]*presenter.Pokemon, error)
	UpdatePokemon(pokeID string, updateData *presenter.UpdatePokemon) (int64, error)
	DeletePokemon(pokeID presenter.DeletePokemon) (int64, error)
}

type UseCase interface {
	CreatePokemon(echo echo.Context) error
	GetPokemon(echo echo.Context) error
	GetAllPokemons(echo echo.Context) error
	UpdatePokemon(echo echo.Context) error
	DeletePokemon(echo echo.Context) error
}

type Repository interface {
	Create(e *presenter.Pokemon) error
	FindAll(ctx context.Context, pokemons []*presenter.Pokemon) ([]*presenter.Pokemon, error)
	FindOne(ctx context.Context, pokeID string, pokemon *presenter.Pokemon) error
	Update(ctx context.Context, pokeID string, updateData *presenter.UpdatePokemon) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, pokeID presenter.DeletePokemon) (*mongo.DeleteResult, error)
}
