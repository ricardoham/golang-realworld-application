package services

import (
	"github.com/ricardoham/pokedex-api/repository"
)

type PokemonsService struct {
	repository *repository.PokemonsRepository
}

func NewPokemonsService(repository *repository.PokemonsRepository) *PokemonsService {
	return &PokemonsService{
		repository: repository,
	}
}
