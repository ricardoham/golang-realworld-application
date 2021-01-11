package entity

import (
	"time"

	"github.com/ricardoham/pokedex-api/api/presenter"

	"github.com/google/uuid"
)

type ID = uuid.UUID

type FavPokemon struct {
	ID          string
	Name        string
	OrderNumber int
	CustomName  string
	Sprites     presenter.PokemonSprites
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DeleteFavPokemon struct {
	ID string
}

func NewFavPokemon(name string, customName string, orderNumber int, sprite presenter.PokemonSprites, createdAt time.Time) (*FavPokemon, error) {
	p := &FavPokemon{
		ID:          newFavPokemonId().String(),
		Name:        name,
		CustomName:  customName,
		OrderNumber: orderNumber,
		Sprites:     sprite,
		CreatedAt:   createdAt,
	}

	return p, nil
}

func newFavPokemonId() ID {
	return ID(uuid.New())
}
