package entity

import (
	"time"

	"github.com/ricardoham/pokedex-api/api/presenter"

	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewPokemon(name string, customName string, orderNumber int, sprite presenter.PokemonSprites, createdAt time.Time) (*presenter.Pokemon, error) {
	p := &presenter.Pokemon{
		ID:          newPokemonId().String(),
		Name:        name,
		CustomName:  customName,
		OrderNumber: orderNumber,
		Sprite:      sprite,
		CreatedAt:   createdAt,
	}

	return p, nil
}

func newPokemonId() ID {
	return ID(uuid.New())
}
