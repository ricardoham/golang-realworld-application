package pokemon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ricardoham/pokedex-api/entity"
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

func (p *PokemonService) GetPokemonFromPokeApi(pokemon string) ([]*entity.Pokemon, error) {
	var pokemonResult []*entity.Pokemon

	dataMarshlead, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon),
		bytes.NewBuffer(dataMarshlead))
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

	if err = json.Unmarshal(body, &pokemonResult); err != nil {
		return nil, err
	}

	return pokemonResult, nil
}
