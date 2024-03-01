include .env
export

build:
	@docker-compose -f deployments/docker-compose.yml up --build

docker-up:
	@docker-compose -f deployments/docker-compose.yml up -d

docker-down:
	@docker-compose -f deployments/docker-compose.yml down

migrate-up:
	@migrate -path ./internal/database/migrations/ -database 'postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=$(DB_SSLMODE)' up

migrate-down:
	@migrate -path ./internal/database/migrations/ -database 'postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=$(DB_SSLMODE)' down

run:
	@go build -o app cmd/junior-test/main.go && ./app

swag:
	@swag init