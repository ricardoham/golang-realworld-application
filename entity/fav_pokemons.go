package entity

import (
	"time"

	"github.com/ricardoham/pokedex-api/api/presenter"

	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewFavPokemon(name string, customName string, orderNumber int, sprite presenter.PokemonSprites, createdAt time.Time) (*presenter.FavPokemon, error) {
	p := &presenter.FavPokemon{
		ID:          newFavPokemonId().String(),
		Name:        name,
		CustomName:  customName,
		OrderNumber: orderNumber,
		Sprite:      sprite,
		CreatedAt:   createdAt,
	}

	return p, nil
}

func newFavPokemonId() ID {
	return ID(uuid.New())
}
