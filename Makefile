.PHONY: docker-compose-build-api
docker-compose-build-api:
	@docker-compose --file docker/docker-compose.yml build

.PHONY: docker-compose-up-api
docker-compose-up-api:
	@docker-compose --file docker/docker-compose.yml up

.PHONY: docker-compose-stop-api
docker-compose-stop-api:
	@docker-compose stop

.PHONY: dev
dev: docker-compose-build-api docker-compose-up-api

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build -o bin/main main.go
