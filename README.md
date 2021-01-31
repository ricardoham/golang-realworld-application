# Pokedex Golang API
[![CircleCI](https://circleci.com/gh/ricardoham/pokedex-api.svg?style=svg)](https://circleci.com/gh/ricardoham/pokedex-api)

[![GitHub license](https://img.shields.io/github/license/ricardoham/pokedex-api?style=for-the-badge)](https://github.com/ricardoham/pokedex-api/blob/main/LICENSE)
## About 📖
### API 🌟
- Using [PokeApi](https://github.com/PokeAPI/pokeapi) for data source
- Golang RestAPI using Echo as web framework routing
- MongoDB as data base
- Redis DB for cache data base for requests(include external request)
- Contains a Dockerfile and docker-compose to up the all containers(Go, Mongo and Redis)

### The flow 🔛
This API simulates a simple flux of a PokeDex
- User Search a Pokemon or all Pokemons.
- Select a Pokemon.
- Create a Section.
- Search for all pokemons

## Running the application ▶
There is a **Makefile** in the project that you can run with:

Developer Mode(No Go Container, good for debug Go code)
```
 make run-dev
 make run
```
 Developer Mode(All containers)
 ```
 make run-api
 ```
  Standalone Go build
```
 make run
```

### Local host results:
```
API port: 8080

MongoDB: 27017

Redis: 6379
```
## cURls 🛠
Get All Pokemons from PokeAPI source
```
curl --location --request GET 'http://localhost:8080/v1/external/pokemons/'
```
Get single pokemon with id(int) or name(string)
```
curl --location --request GET 'http://localhost:8080/v1/external/pokemons/{id or name}'
```

Get all Pokemons from podex-api:
```
curl --location --request GET 'http://localhost:8080/v1/pokemons'
```
Get a single with id(string) Pokemon from podex-api:
```
curl --location --request GET 'http://localhost:8080/v1/pokemons/{id}
```
Create Pokemon on podex-api:
```
curl --location --request POST 'http://localhost:8080/v1/pokemons' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Pikachu",
    "customName": "Spark"
    "favorite": true
}'
```

Delete Pokemon from podex-api:
```
curl --location --request DELETE 'http://localhost:8080/v1/pokemons' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "2e82f233-5629-41f5-8aa7-133fa34194a8"
}'
```
