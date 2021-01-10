package entity

import (
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID

type FavPokemon struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeleteFavPokemon struct {
	ID string
}

func NewFavPokemon(name string, createdAt time.Time) (*FavPokemon, error) {
	p := &FavPokemon{
		ID:        newFavPokemonId().String(),
		Name:      name,
		CreatedAt: createdAt,
	}

	return p, nil
}

func newFavPokemonId() ID {
	return ID(uuid.New())
}
