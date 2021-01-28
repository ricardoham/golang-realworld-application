package presenter

import "time"

type Pokemon struct {
	ID          string         `json:"id" bson:"id"`
	Name        string         `json:"name" bson:"name"`
	OrderNumber int            `json:"orderNumber" bson:"orderNumber"`
	CustomName  string         `json:"customName" bson:"customName"`
	Favorite    bool           `json:"favorite" bson:"favorite"`
	Sprite      PokemonSprites `json:"sprites" bson:"sprites"`
	CreatedAt   time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type SavePokemon struct {
	Name       string `json:"name" bson:"name"`
	CustomName string `json:"customName" bson:"customName"`
	Favorite   bool   `json:"favorite" bson:"favorite"`
}

type UpdatePokemon struct {
	CustomName string `json:"customName" bson:"customName"`
	Favorite   bool   `json:"favorite" bson:"favorite"`
}

type DeletePokemon struct {
	ID string `json:"id" bson:"id"`
}
