package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-redis/redis"
	"github.com/ricardoham/pokedex-api/api/presenter"
	"github.com/ricardoham/pokedex-api/infrastructure/cache"
	mockHttpClient "github.com/ricardoham/pokedex-api/mocks/client"
	mockRedis "github.com/ricardoham/pokedex-api/mocks/redis"
	"github.com/stretchr/testify/mock"
)

func TestPokemonService(t *testing.T) {
	type service struct {
		httpClient HTTPClient
		cache      cache.Redis
		url        string
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
				json := `{
					"id": 1,
					"name": "bulbasaur",
					"sprites": {
							"front_default": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png"
					}
			}`
				mockBody := ioutil.NopCloser(bytes.NewReader([]byte(json)))
				httpRes := &http.Response{
					StatusCode: 200,
					Body:       mockBody,
				}
				httpClient := &mockHttpClient.HTTPClient{}
				cache := &mockRedis.Redis{}
				cache.On("Get", mock.Anything, mock.Anything).Return(redis.Nil)
				httpClient.On("Do", mock.Anything).Return(httpRes, nil)
				cache.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
				return service{
					httpClient: httpClient,
					cache:      cache,
					url:        "http://localhost:8080/",
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
				tt.service.httpClient,
				tt.service.cache,
				tt.service.url,
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
