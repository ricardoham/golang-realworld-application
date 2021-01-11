package presenter

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
