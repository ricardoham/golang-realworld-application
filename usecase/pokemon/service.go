package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ricardoham/pokedex-api/api/presenter"
)

type PokemonService struct {
	http.Client
}

func NewPokemonService() *PokemonService {
	return &PokemonService{
		Client: http.Client{
			Timeout: time.Second,
		},
	}
}

func (p *PokemonService) GetPokemonFromPokeApi(pokemon string) (*presenter.Pokemon, error) {
	var pokemonResult *presenter.Pokemon

	bodyResult, err := p.doRequest(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bodyResult, &pokemonResult); err != nil {
		return nil, err
	}
	return pokemonResult, nil
}

func (p *PokemonService) GetAllResultPokemonFromPokeApi() (*presenter.Result, error) {
	var pokemonResult *presenter.Result

	bodyResult, err := p.doRequest(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/"))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bodyResult, &pokemonResult); err != nil {
		return nil, err
	}
	return pokemonResult, nil
}

func (p *PokemonService) doRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(
		"GET",
		url,
		nil)
	if err != nil {
		return nil, err
	}

	res, err := p.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
