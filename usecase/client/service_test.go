package client

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-redis/redis"
	"github.com/ricardoham/pokedex-api/api/presenter"
	"github.com/ricardoham/pokedex-api/infrastructure/cache"
	mocks "github.com/ricardoham/pokedex-api/mocks/redis"
	"github.com/stretchr/testify/mock"
)

func TestPokemonService(t *testing.T) {
	type service struct {
		client *http.Client
		cache  cache.Redis
	}

	type args struct {
		pokemon string
	}

	pokemonMock := &presenter.ClientPokemon{
		ID:   1,
		Name: "bulbasaur",
		Sprite: presenter.PokemonSprites{
			Front: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png",
		},
	}

	myTest := []struct {
		inputName      string
		service        service
		args           args
		expectedError  bool
		expectedOutput *presenter.ClientPokemon
	}{
		{
			inputName: "Should return a pokemon as result of GetPokemonFromPokeApi",
			service: func() service {
				want := "Success!"
				httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)
					w.Write([]byte(want))
				}))
				cache := &mocks.Redis{}

				cache.On("Get", mock.Anything, mock.Anything).Return(redis.Nil)
				cache.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
				return service{
					client: httpServer.Client(),
					cache:  cache,
				}
			}(),
			args: args{
				pokemon: "bulbasaur",
			},
			expectedError:  false,
			expectedOutput: pokemonMock,
		},
	}
	for _, tt := range myTest {
		t.Run(tt.inputName, func(t *testing.T) {
			p := PokemonService{
				tt.service.client,
				tt.service.cache,
			}
			result, err := p.GetPokemonFromPokeApi(tt.args.pokemon)
			if (err != nil) != tt.expectedError {
				t.Errorf("Error on test %v", err)
				return
			}
			if !reflect.DeepEqual(result, tt.expectedOutput) {
				t.Errorf("Error on test want %v have %v", tt.expectedOutput, result)
			}
		})
	}
}
