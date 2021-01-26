package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ricardoham/pokedex-api/api/presenter"
	"github.com/ricardoham/pokedex-api/infrastructure/cache"
)

type PokemonService struct {
	*http.Client
	cache *cache.Cache
}

func NewPokemonService(cache *cache.Cache) *PokemonService {
	return &PokemonService{
		Client: &http.Client{
			Timeout: time.Second,
		},
		cache: cache,
	}
}

func (p *PokemonService) GetPokemonFromPokeApi(pokemon string) (*presenter.Pokemon, error) {
	var pokemonResult *presenter.Pokemon

	err := p.cache.Get("pokeApi", &pokemonResult)
	if err == nil {
		return pokemonResult, nil
	}

	bodyResult, err := p.doRequest(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bodyResult, &pokemonResult); err != nil {
		return nil, err
	}

	isSetted, err := p.cache.Set("pokeApi", pokemonResult, 320)
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

	err := p.cache.Get("allPokeApi", &pokemonResult)
	if err == nil {
		return pokemonResult, nil
	}

	bodyResult, err := p.doRequest(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/"))
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
