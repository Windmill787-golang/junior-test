include .env

docker-build:
	@docker-compose -f deployments/docker-compose.yml up --build

docker-up:
	@docker-compose -f deployments/docker-compose.yml up -d

docker-down:
	@docker-compose -f deployments/docker-compose.yml down

migrate-up:
	@migrate -path ./internal/database/migrations/ -database 'postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=$(DB_SSLMODE)' up

migrate-down:
	@migrate -path ./internal/database/migrations/ -database 'postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=$(DB_SSLMODE)' down

build:
	@go build -o junior-test cmd/junior-test/main.go

run:
	@go build -o junior-test cmd/junior-test/main.go && ./junior-test

swag:
	@swag init -g ./internal/app/app.go