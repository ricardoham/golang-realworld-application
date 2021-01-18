package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mocks "github.com/ricardoham/pokedex-api/mocks/favpokemon"
	"github.com/stretchr/testify/mock"

	"github.com/labstack/echo"

	"github.com/ricardoham/pokedex-api/api/presenter"
	services "github.com/ricardoham/pokedex-api/usecase/favpokemon"
)

func TestFavPokemonsHandler(t *testing.T) {
	type handler struct {
		pokemonService services.FavPokemon
	}

	type args struct {
		pokemon presenter.SaveFavPokemon
	}
	body := strings.NewReader(`{"name": "Pikachu", "customName": "Test"}`)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080", body)
	w := httptest.NewRecorder()
	r.Header.Set("content-type", "application/json")

	echo := echo.New()
	echoContext := echo.NewContext(r, w)

	myTest := []struct {
		inputName       string
		handler         handler
		args            args
		expectedError   bool
		excpectedOutput int
	}{
		{
			inputName: "Should return 200 when CreateFavPokemon",
			handler: func() handler {
				pokemonService := &mocks.FavPokemon{}
				pokemonService.On("CreateFavPokemon", mock.Anything).Return(nil)

				return handler{
					pokemonService: pokemonService,
				}
			}(),
			expectedError:   false,
			excpectedOutput: http.StatusCreated,
		},
	}
	for _, tt := range myTest {
		t.Run(tt.inputName, func(t *testing.T) {
			p := favPokemonsHandler{
				pokemonService: tt.handler.pokemonService,
			}
			err := p.CreateFavPokemon(echoContext)
			if (err != nil) != tt.expectedError {
				t.Errorf("Error on test %v", err)
				return
			}
			if tt.excpectedOutput != w.Code {
				t.Errorf("Unexpected error in the request %v - wants %v", w.Code, tt.excpectedOutput)
			}
		})
	}
}
