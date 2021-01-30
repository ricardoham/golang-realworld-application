package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ricardoham/pokedex-api/api/presenter"
	"github.com/ricardoham/pokedex-api/infrastructure/cache"
)

type PokemonService struct {
	httpClient HTTPClient
	cache      cache.Redis
	url        string
}

func NewPokemonService(cache cache.Redis) *PokemonService {
	return &PokemonService{
		httpClient: &http.Client{},
		cache:      cache,
		url:        "https://pokeapi.co/api/v2/pokemon/",
	}
}

func (p *PokemonService) GetPokemonFromPokeApi(pokemon string) (*presenter.ClientPokemon, error) {
	var pokemonResult *presenter.ClientPokemon

	redisKey := fmt.Sprintf("pokeApi-%s", pokemon)
	err := p.cache.Get(redisKey, pokemonResult)
	if err == nil {
		return pokemonResult, nil
	}

	bodyResult, err := p.doRequest(fmt.Sprintf("%s%s", p.url, strings.ToLower(pokemon)))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bodyResult, &pokemonResult); err != nil {
		return nil, err
	}

	isSetted, err := p.cache.Set("pokeApi", pokemonResult, 60)
	if err != nil {
		log.Println("Failed to set new cache on Redis", err)
		return pokemonResult, nil
	}
	if isSetted {
		log.Println("Cache is set")
	}

	return pokemonResult, nil
}

func (p *PokemonService) GetAllResultPokemonFromPokeApi() (*presenter.Result, error) {
	var pokemonResult *presenter.Result

	err := p.cache.Get("allPokeApi", pokemonResult)
	if err == nil {
		return pokemonResult, nil
	}

	bodyResult, err := p.doRequest(fmt.Sprintf("%s", p.url))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bodyResult, &pokemonResult); err != nil {
		return nil, err
	}

	isSetted, err := p.cache.Set("allPokeApi", pokemonResult, 320)
	if err != nil {
		log.Println("Failed to set new cache on Redis", err)
		return pokemonResult, nil
	}
	if isSetted {
		log.Println("Cache is set")
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

	res, err := p.httpClient.Do(req)
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
