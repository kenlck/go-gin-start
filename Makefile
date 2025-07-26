.PHONY: dev migrate migrate-up migrate-down migrate-create

dev:
	@command -v air >/dev/null 2>&1 && air || (echo "Air not installed. Running without hot reload."; go run main.go)

start:
	go run main.go

install:
	go mod tidy

migrate:
	@bash -c 'source .env && migrate -database "$$DATABASE_URL" -path ./migrations up'

migrate-up:
	@bash -c 'source .env && migrate -database "$$DATABASE_URL" -path ./migrations up'

migrate-down:
	@bash -c 'source .env && migrate -database "$$DATABASE_URL" -path ./migrations down'

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name
