package handler

import (
	"github.com/ricardoham/pokedex-api/services"
)

type PokemonsInterface interface {
} // TODO

type pokemonsHandler struct {
	pokemonsService *services.PokemonsService
}

func NewPokemonsHandler(services *services.PokemonsService) PokemonsInterface {
	return &pokemonsHandler{
		pokemonsService: services,
	}
}
