package entity

import "github.com/google/uuid"

type ID = uuid.UUID

type Pokemon struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

func NewPokemon(name string) (*Pokemon, error) {
	p := &Pokemon{
		ID:   newPokemonId().String(),
		Name: name,
	}

	return p, nil
}

func newPokemonId() ID {
	return ID(uuid.New())
}
