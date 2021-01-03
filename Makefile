.PHONY: docker-compose-build-api
docker-compose-build-api:
	@docker-compose --file docker/docker-compose.yml build

.PHONY: docker-compose-up-api
docker-compose-up-api:
	@docker-compose --file docker/docker-compose.yml up

.PHONY: docker-compose-stop-api
docker-compose-stop-api:
	@docker-compose stop

.PHONY: run-api
run-api: docker-compose-build-api docker-compose-up-api

.PHONY: dev-docker-compose-build-api
dev-docker-compose-build-api:
	@docker-compose --file docker/dev.docker-compose.yml build

.PHONY: dev-docker-compose-up-api
dev-docker-compose-up-api:
	@docker-compose --file docker/dev.docker-compose.yml up

.PHONY: dev-docker-compose-stop-api
dev-docker-compose-stop-api:
	@docker-compose stop

.PHONY: run-dev
run-dev: dev-docker-compose-build-api dev-docker-compose-up-api

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build -o bin/main main.go
