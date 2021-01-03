package pokemon

import (
	"context"
	"time"

	"github.com/google/uuid"
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
	p, err := entity.NewPokemon(pokemon.Name, time.Now())
	if err != nil {
		return err
	}

	return s.repository.Create(p)
}

func (s *PokemonService) GetAllPokemons() ([]*entity.Pokemon, error) {
	var pokemons []*entity.Pokemon
	ctx := context.TODO()
	pokemons, err := s.repository.FindAll(ctx, pokemons)

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (s *PokemonService) UpdatePokemon(pokeId uuid.UUID, updateData *entity.Pokemon) (int64, error) {
	ctx := context.TODO()
	result, err := s.repository.Update(ctx, pokeId, updateData)
	if err != nil {
		return 0, err
	}

	return result.MatchedCount, nil
}

func (s *PokemonService) DeletePokemon(pokeId entity.DeletePokemon) (int64, error) {
	ctx := context.TODO()
	deleteResult, err := s.repository.Delete(ctx, pokeId)
	if err != nil {
		return 0, err
	}

	return deleteResult.DeletedCount, nil
}
