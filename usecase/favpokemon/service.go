package favpokemon

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ricardoham/pokedex-api/api/presenter"
	"github.com/ricardoham/pokedex-api/entity"
	repository "github.com/ricardoham/pokedex-api/infrastructure/repository"
	services "github.com/ricardoham/pokedex-api/usecase/pokemon"
)

type FavPokemonService struct {
	repository     *repository.FavPokemonsRepository
	pokeAPIService *services.PokemonService
}

func NewFavPokemonsService(
	repository *repository.FavPokemonsRepository,
	pokeAPIService *services.PokemonService) *FavPokemonService {
	return &FavPokemonService{
		repository:     repository,
		pokeAPIService: pokeAPIService,
	}
}

func (s *FavPokemonService) CreateFavPokemon(pokemon *presenter.SaveFavPokemon) error {
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

func (s *FavPokemonService) GetAllFavPokemons() ([]*presenter.FavPokemon, error) {
	var pokemons []*presenter.FavPokemon
	ctx := context.TODO()
	pokemons, err := s.repository.FindAll(ctx, pokemons)

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (s *FavPokemonService) UpdateFavPokemon(pokeId uuid.UUID, updateData *presenter.FavPokemon) (int64, error) {
	ctx := context.TODO()
	result, err := s.repository.Update(ctx, pokeId, updateData)
	if err != nil {
		return 0, err
	}

	return result.MatchedCount, nil
}

func (s *FavPokemonService) DeleteFavPokemon(pokeId presenter.DeleteFavPokemon) (int64, error) {
	ctx := context.TODO()
	deleteResult, err := s.repository.Delete(ctx, pokeId)
	if err != nil {
		return 0, err
	}

	return deleteResult.DeletedCount, nil
}
