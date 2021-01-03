package entity

import (
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID

type Pokemon struct {
	ID        string    `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type DeletePokemon struct {
	ID string `json:"id" bson:"id"`
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
