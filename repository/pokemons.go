package repository

import (
	"github.com/ricardoham/pokedex-api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonsRepository struct {
	client     *mongo.Client
	collection string
	dbName     string
}

func NewPokemonsRepository() *PokemonsRepository {
	client, dbName := config.MongoConnection()
	collection := "pokemons"
	return &PokemonsRepository{
		client:     client,
		collection: collection,
		dbName:     dbName,
	}
}
