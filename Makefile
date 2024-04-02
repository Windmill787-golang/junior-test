include .env

docker-build:
	@echo "Building docker containers..."
	@docker-compose -f deployments/docker-compose.yml -p junior-test up --build

docker-up:
	@echo "Starting docker containers..."
	@docker-compose -f deployments/docker-compose.yml -p junior-test up -d

docker-down:
	@docker-compose -f deployments/docker-compose.yml -p junior-test down

docker-shell:
	@docker exec -it postgres sh -c 'psql -U $(POSTGRES_USERNAME) -d $(POSTGRES_DATABASE)'

migrate-up:
	@migrate -path ./internal/database/migrations/ -database 'postgres://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSLMODE)' up

migrate-down:
	@migrate -path ./internal/database/migrations/ -database 'postgres://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSLMODE)' down

build:
	@echo "Building project..."
	@go build -o junior-test cmd/junior-test/main.go

run:
	@echo "Running project..."
	@go build -o junior-test cmd/junior-test/main.go && ./junior-test

swag:
	@swag init -g ./internal/app/app.go
