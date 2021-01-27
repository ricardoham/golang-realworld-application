package repository

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ricardoham/pokedex-api/api/presenter"
	"github.com/ricardoham/pokedex-api/config"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *PokemonsRepository) Create(e *presenter.Pokemon) error {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	_, err := coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatal("Error on repository", err)
		return err
	}
	return err
}

func (r *PokemonsRepository) FindAll(ctx context.Context, pokemons []*presenter.Pokemon) ([]*presenter.Pokemon, error) {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	cursor, err := coll.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var el presenter.Pokemon
		err := cursor.Decode(&el)
		if err != nil {
			return nil, err
		}

		pokemons = append(pokemons, &el)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return pokemons, err
}

func (r *PokemonsRepository) FindOne(ctx context.Context, pokeId uuid.UUID, pokemon *presenter.Pokemon) error {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	filter := bson.M{"id": pokeId}
	err := coll.FindOne(ctx, filter).Decode(&pokemon)
	if err != nil {
		return err
	}

	return nil
}

func (r *PokemonsRepository) Update(ctx context.Context, pokeId uuid.UUID, updateData *presenter.Pokemon) (*mongo.UpdateResult, error) {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	filter := bson.M{"id": pokeId}
	update := bson.D{
		{"$set", bson.M{
			"name":      updateData.Name,
			"updatedAt": time.Now(),
		}},
	}
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *PokemonsRepository) Delete(ctx context.Context, pokeId presenter.DeletePokemon) (*mongo.DeleteResult, error) {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	filter := bson.M{"id": pokeId.ID}
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
