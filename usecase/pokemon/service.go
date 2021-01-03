package pokemon

import (
	"github.com/ricardoham/pokedex-api/entity"
	repository "github.com/ricardoham/pokedex-api/infrastructure/repository"
)

type PokemonService struct {
	repository *repository.PokemonsRepository
}

func NewPokemonsService(repository *repository.PokemonsRepository) *PokemonService {
	return &PokemonService{
		repository: repository,
	}
}

func (s *PokemonService) CreatePokemon(pokemon *entity.Pokemon) error {
	p, err := entity.NewPokemon(pokemon.Name)
	if err != nil {
		return err
	}

	return s.repository.Create(p)
}
