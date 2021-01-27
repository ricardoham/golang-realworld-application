package client

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ricardoham/pokedex-api/api/presenter"
)

func TestPokemonService(t *testing.T) {
	type service struct {
		client *http.Client
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
				return service{
					client: httpServer.Client(),
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
