package repository

//go:generate mockgen -destination=infrastructure/repository/mocks/fav_pokemons.go -package=mock_repository github.com/ricardoham/pokedex-api/infrastructure/repository FavPokemonsRepository

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

type FavPokemonsRepository struct {
	client     *mongo.Client
	collection string
	dbName     string
}

func NewPokemonsRepository() *FavPokemonsRepository {
	client, dbName := config.MongoConnection()
	collection := "pokemons"
	return &FavPokemonsRepository{
		client:     client,
		collection: collection,
		dbName:     dbName,
	}
}

func (r *FavPokemonsRepository) Create(e *presenter.FavPokemon) error {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	_, err := coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatal("Error on repository", err)
		return err
	}
	return err
}

func (r *FavPokemonsRepository) FindAll(ctx context.Context, pokemons []*presenter.FavPokemon) ([]*presenter.FavPokemon, error) {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	cursor, err := coll.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var el presenter.FavPokemon
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

func (r *FavPokemonsRepository) Update(ctx context.Context, pokeId uuid.UUID, updateData *presenter.FavPokemon) (*mongo.UpdateResult, error) {
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

func (r *FavPokemonsRepository) Delete(ctx context.Context, pokeId presenter.DeleteFavPokemon) (*mongo.DeleteResult, error) {
	coll := r.client.Database(r.dbName).Collection(r.collection)
	filter := bson.M{"id": pokeId.ID}
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
