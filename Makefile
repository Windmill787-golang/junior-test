include .env

docker-build:
	@docker-compose -f deployments/docker-compose.yml -p junior-test up --build

docker-up:
	@docker-compose -f deployments/docker-compose.yml -p junior-test up -d

docker-down:
	@docker-compose -f deployments/docker-compose.yml -p junior-test down

migrate-up:
	@migrate -path ./internal/database/migrations/ -database 'postgres://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSLMODE)' up

migrate-down:
	@migrate -path ./internal/database/migrations/ -database 'postgres://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSLMODE)' down

build:
	@go build -o junior-test cmd/junior-test/main.go

run:
	@go build -o junior-test cmd/junior-test/main.go && ./junior-test

swag:
	@swag init -g ./internal/app/app.go