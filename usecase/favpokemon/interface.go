package favpokemon

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/ricardoham/pokedex-api/api/presenter"
	"go.mongodb.org/mongo-driver/mongo"
)

type FavPokemon interface {
	CreateFavPokemon(pokemon *presenter.SaveFavPokemon) error
	GetAllFavPokemons() ([]*presenter.FavPokemon, error)
	UpdateFavPokemon(pokeId uuid.UUID, updateData *presenter.FavPokemon) (int64, error)
	DeleteFavPokemon(pokeId presenter.DeleteFavPokemon) (int64, error)
}

type UseCase interface {
	CreateFavPokemon(echo echo.Context) error
	GetAllFavPokemons(echo echo.Context) error
	UpdateFavPokemon(echo echo.Context) error
	DeleteFavPokemon(echo echo.Context) error
}

type Repository interface {
	Create(e *presenter.FavPokemon) error
	FindAll(ctx context.Context, pokemons []*presenter.FavPokemon) ([]*presenter.FavPokemon, error)
	Update(ctx context.Context, pokeId uuid.UUID, updateData *presenter.FavPokemon) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, pokeId presenter.DeleteFavPokemon) (*mongo.DeleteResult, error)
}
