# go-gin-start

A starter Go backend API project using Gin, PostgreSQL (pgx), Squirrel, golang-migrate, JWT authentication, bcrypt password hashing, and godotenv for configuration.

## Features

- JWT-based authentication with login endpoint
- Password hashing and validation using bcrypt
- Middleware to protect routes using JWT
- Automatic migration on app startup using golang-migrate
- Loads `DATABASE_URL` and `JWT_SECRET` from `.env`
- Multi-stage Dockerfile for containerized builds

## Tech Stack

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver
- [Squirrel](https://github.com/Masterminds/squirrel) - SQL query builder
- [golang-migrate](https://github.com/golang-migrate/migrate) - DB migrations
- [JWT](https://github.com/golang-jwt/jwt) - Authentication
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password hashing
- [godotenv](https://github.com/joho/godotenv) - .env loader

## Project Structure

```
.
├── .env
├── Dockerfile
├── go.mod / go.sum
├── main.go
├── internal/
│   ├── auth/
│   ├── db/
│   ├── handler/
│   └── model/
├── migrations/
│   └── 001_create_users.sql
```

## Setup

1. **Clone the repo**
   `git clone <your-repo-url> && cd go-gin-start`

2. **Set environment variables**
   Edit `.env`:

   ```
   DATABASE_URL=postgres://user:password@localhost:5432/go_gin_start?sslmode=disable
   JWT_SECRET=your_jwt_secret_key
   ```

3. **Install dependencies**

   ```
   go mod tidy
   ```

4. **Run PostgreSQL**
   Make sure PostgreSQL is running and accessible via `DATABASE_URL`.

## Running Locally

```
go run main.go
```

- The server runs on `http://localhost:8080`
- On startup, migrations in `migrations/` are applied automatically.

## Docker

Build and run with Docker:

```
docker build -t go-gin-start .
docker run --env-file .env -p 8080:8080 go-gin-start
```

## Endpoints

- `POST /login` - Login with username and password, returns JWT
- `GET /api/me` - Protected route, returns current user info (requires Authorization header)

## Environment Variables

- `DATABASE_URL` - PostgreSQL connection string
- `JWT_SECRET` - Secret key for JWT signing

## Makefile Commands

The Makefile automates common development and migration tasks. Migration commands automatically load `DATABASE_URL` from your `.env` file.

**Available commands:**

- `make dev` — Run the app locally
- `make migrate-up` — Apply all pending migrations (uses .env)
- `make migrate-down` — Roll back the last migration (uses .env)
- `make migrate-create` — Create a new migration interactively

**Usage example:**

```sh
make dev
make migrate-up
make migrate-create
```

You can override `DATABASE_URL` if needed:

```sh
make migrate-up DATABASE_URL="your_connection_string"
```

## Generating Migrations

1. **Install golang-migrate CLI**
   See [golang-migrate docs](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) or use Homebrew:

   ```
   brew install golang-migrate
   ```

2. **Create a new migration**

   ```
   migrate create -ext sql -dir migrations -seq <migration_name>
   ```

   Example:

   ```
   migrate create -ext sql -dir migrations -seq add_email_to_users
   ```

3. **Edit the generated .sql files in `migrations/` as needed.**

## Development

- Edit code in `internal/` for features and logic.
- Add new migrations in `migrations/` as needed.

---
