package repository

import (
	"context"
	"log"

	"github.com/ricardoham/pokedex-api/config"
	"github.com/ricardoham/pokedex-api/entity"
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

func (r *PokemonsRepository) Create(e *entity.Pokemon) error {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	_, err := coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatal("Error on repository", err)
		return err
	}
	return err
}
