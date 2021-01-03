package repository

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ricardoham/pokedex-api/config"
	"github.com/ricardoham/pokedex-api/entity"
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

func (r *PokemonsRepository) Create(e *entity.Pokemon) error {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	_, err := coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatal("Error on repository", err)
		return err
	}
	return err
}

func (r *PokemonsRepository) FindAll(ctx context.Context, pokemons []*entity.Pokemon) ([]*entity.Pokemon, error) {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	cursor, err := coll.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var el entity.Pokemon
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

func (r *PokemonsRepository) Update(ctx context.Context, pokeId uuid.UUID, updateData *entity.Pokemon) (*mongo.UpdateResult, error) {
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

func (r *PokemonsRepository) Delete(ctx context.Context, pokeId entity.DeletePokemon) (*mongo.DeleteResult, error) {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	filter := bson.M{"id": pokeId.ID}
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
