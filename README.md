# Pokedex Golang API

## About
### The API
Golang RestAPI using Echo as web framework routing
MongoDB as data base
Redis DB for cache data base
Contains a Dockerfile and docker-compose to up the all containers(Go, Mongo and Redis)

### About the flux
This API simulates a simple flux of a PokeDex
- User Search a Pokemon or all Pokemons.
- Select a Pokemon.
- Create a Section.
- Search for all pokemons

## Running the application
There is a **Makefile** in the project that you can run with:
```
 Developer Mode(No Go Container, good for debug Go code)
 - make run-dev
 - make run

 Developer Mode(All containers)
 - make run-api

 Just Go code
 - make run
```

## cURls
````
Get All Pokemons From PokeAPI source

curl --location --request GET 'http://localhost:8080/v1/external/pokemons/'


Get single pokemon with id(int) or name(string)

curl --location --request GET 'http://localhost:8080/v1/external/pokemons/{bulbasaur}'


Get all Pokemons from podex-api:
curl --location --request GET 'http://localhost:8080/v1/pokemons'

Create Pokemon on podex-api:
curl --location --request POST 'http://localhost:8080/v1/pokemons' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Pikachu",
    "customName": "Spark"
}'
````
