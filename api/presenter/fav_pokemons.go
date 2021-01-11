package presenter

import "time"

type FavPokemon struct {
	ID         string         `json:"id" bson:"id"`
	Name       string         `json:"name" bson:"name"`
	CustomName string         `json:"customName" bson:"customName"`
	Sprite     PokemonSprites `json:"sprites" bson:"sprites"`
	CreatedAt  time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type SaveFavPokemon struct {
	Name       string `json:"name" bson:"name"`
	CustomName string `json:"customName" bson:"customName"`
}

type DeleteFavPokemon struct {
	ID string `json:"id" bson:"id"`
}
