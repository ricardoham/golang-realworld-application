package pokemon

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ricardoham/pokedex-api/api/presenter"
	"github.com/ricardoham/pokedex-api/entity"
	"github.com/ricardoham/pokedex-api/infrastructure/cache"
	"github.com/ricardoham/pokedex-api/usecase/client"
)

type PokemonService struct {
	repository     Repository
	pokeAPIService client.ClientPokemon
	cache          *cache.Cache
}

func NewPokemonsService(
	repository Repository,
	cache *cache.Cache,
	pokeAPIService client.ClientPokemon) *PokemonService {
	return &PokemonService{
		repository:     repository,
		pokeAPIService: pokeAPIService,
		cache:          cache,
	}
}

func (s *PokemonService) CreatePokemon(pokemon *presenter.SavePokemon) error {
	r, err := s.pokeAPIService.GetPokemonFromPokeApi(pokemon.Name)
	if err != nil {
		return err
	}

	p, err := entity.NewFavPokemon(pokemon.Name, pokemon.CustomName, r.ID, r.Sprite, time.Now())
	if err != nil {
		return err
	}

	return s.repository.Create(p)
}

func (s *PokemonService) GetPokemon(pokeId uuid.UUID) (*presenter.Pokemon, error) {
	var pokemon *presenter.Pokemon
	ctx := context.TODO()

	err := s.repository.FindOne(ctx, pokeId, pokemon)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

func (s *PokemonService) GetAllPokemons() ([]*presenter.Pokemon, error) {
	var pokemons []*presenter.Pokemon
	ctx := context.TODO()

	err := s.cache.Get("favPokemon", &pokemons)
	if err == nil {
		return pokemons, nil
	}

	pokemons, err = s.repository.FindAll(ctx, pokemons)
	if err != nil {
		return nil, err
	}

	isSetted, err := s.cache.Set("favPokemon", pokemons, 320)
	if err != nil {
		log.Println("Failed to set new cache on Redis", err)
		return pokemons, nil
	}
	if isSetted {
		log.Println("Cache is set")
	}

	return pokemons, nil
}

func (s *PokemonService) UpdatePokemon(pokeId uuid.UUID, updateData *presenter.Pokemon) (int64, error) {
	ctx := context.TODO()
	result, err := s.repository.Update(ctx, pokeId, updateData)
	if err != nil {
		return 0, err
	}

	return result.MatchedCount, nil
}

func (s *PokemonService) DeletePokemon(pokeId presenter.DeletePokemon) (int64, error) {
	ctx := context.TODO()
	deleteResult, err := s.repository.Delete(ctx, pokeId)
	if err != nil {
		return 0, err
	}

	return deleteResult.DeletedCount, nil
}
