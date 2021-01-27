package presenter

type ClientPokemon struct {
	ID     int            `json:"id"`
	Name   string         `json:"name"`
	Sprite PokemonSprites `json:"sprites"`
}

type PokemonSprites struct {
	Front string `json:"front_default"`
}

type Result struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Result   []ResultAllPokemon `json:"results"`
}

type ResultAllPokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
