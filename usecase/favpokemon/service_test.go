package favpokemon

import (
	"testing"

	"github.com/stretchr/testify/mock"

	mockClient "github.com/ricardoham/pokedex-api/mocks/client"
	mockFavpokemon "github.com/ricardoham/pokedex-api/mocks/favpokemon"

	"github.com/ricardoham/pokedex-api/api/presenter"
	"github.com/ricardoham/pokedex-api/usecase/client"
)

// Basesd on TDT(Table Driven Tests)
func TestFavPokemonService(t *testing.T) {
	type service struct {
		repository     Repository
		pokeAPIService client.ClientPokemon
	}

	type args struct {
		pokemon *presenter.SaveFavPokemon
	}

	myTest := []struct {
		inputName     string
		service       service
		args          args
		expectedError bool
	}{
		{
			inputName: "Should Return Create Pokemon when CreateFavPokemon",
			service: func() service {
				pokemonMock := &presenter.ClientPokemon{
					ID:   1,
					Name: "bulbasaur",
					Sprite: presenter.PokemonSprites{
						Front: "test",
					},
				}
				repository := &mockFavpokemon.Repository{}
				pokeAPIService := &mockClient.ClientPokemon{}

				pokeAPIService.On("GetPokemonFromPokeApi", mock.Anything).Return(pokemonMock, nil)
				repository.On("Create", mock.Anything).Return(nil)

				return service{
					repository:     repository,
					pokeAPIService: pokeAPIService,
				}
			}(),
			args: args{
				pokemon: &presenter.SaveFavPokemon{
					Name:       "Bulbasaur",
					CustomName: "Test",
				},
			},
			expectedError: false,
		},
	}
	for _, tt := range myTest {
		t.Run(tt.inputName, func(t *testing.T) {
			p := FavPokemonService{
				repository:     tt.service.repository,
				pokeAPIService: tt.service.pokeAPIService,
			}
			err := p.CreateFavPokemon(tt.args.pokemon)
			if (err != nil) != tt.expectedError {
				t.Errorf("Error on test %v", err)
				return
			}
		})
	}
}
