.PHONY: dev migrate migrate-up migrate-down migrate-create

dev:
	@command -v air >/dev/null 2>&1 && air || (echo "Air not installed. Running without hot reload."; go run main.go)

start:
	go run main.go

install:
	go mod tidy

migrate-create:
	@if ! grep -q '^DATABASE_URL=' .env; then \
		echo "Error: DATABASE_URL not found in .env"; exit 1; \
	fi; \
	if ! grep -q '^DEV_DATABASE_URL=' .env; then \
		echo "Error: DEV_DATABASE_URL not found in .env"; exit 1; \
	fi; \
	db_url=$$(grep '^DATABASE_URL=' .env | cut -d'=' -f2-); \
	dev_db_url=$$(grep '^DEV_DATABASE_URL=' .env | cut -d'=' -f2-); \
	if [ "$$db_url" = "$$dev_db_url" ]; then \
		echo "Error: DEV_DATABASE_URL and DATABASE_URL must not be the same"; exit 1; \
	fi; \
	read -p "Migration name: " name; \
	name=$${name// /_}; \
	atlas migrate diff $$name --dir "file://migrations" --to "file://schema.pg.hcl" --dev-url "$$dev_db_url"

migrate-apply:
	@if ! grep -q '^DATABASE_URL=' .env; then \
		echo "Error: DATABASE_URL not found in .env"; exit 1; \
	fi; \
	db_url=$$(grep '^DATABASE_URL=' .env | cut -d'=' -f2-); \
	atlas migrate apply --dir "file://migrations" --url "$$db_url"
