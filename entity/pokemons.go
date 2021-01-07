package entity

import (
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID

type Pokemon struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeletePokemon struct {
	ID string
}

func NewPokemon(name string, createdAt time.Time) (*Pokemon, error) {
	p := &Pokemon{
		ID:        newPokemonId().String(),
		Name:      name,
		CreatedAt: createdAt,
	}

	return p, nil
}

func newPokemonId() ID {
	return ID(uuid.New())
}
