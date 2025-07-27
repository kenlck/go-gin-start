# go-gin-start

A starter Go backend API project using Gin, PostgreSQL (pgx), Squirrel, Atlas for migrations, JWT authentication, bcrypt password hashing, and godotenv for configuration.

## Features

- JWT-based authentication with login endpoint
- Password hashing and validation using bcrypt
- Middleware to protect routes using JWT
- Database migrations managed with Atlas
- Loads `DATABASE_URL` and `JWT_SECRET` from `.env`
- Multi-stage Dockerfile for containerized builds

## Tech Stack

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver
- [Squirrel](https://github.com/Masterminds/squirrel) - SQL query builder
- [Atlas](https://atlasgo.io/) - Database migrations
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
│   └── 20250727030018_add_users.sql
├── schema.pg.hcl
```

## Setup

1. **Clone the repo**
   `git clone <your-repo-url> && cd go-gin-start`

2. **Set environment variables**
   Edit `.env`:

   ```
   DATABASE_URL=postgres://user:password@localhost:5432/go_gin_start?sslmode=disable
   DEV_DATABASE_URL=postgres://user:password@localhost:5432/go_gin_start_dev?sslmode=disable
   JWT_SECRET=your_jwt_secret_key
   ```

   **Note:** Both `DATABASE_URL` and `DEV_DATABASE_URL` must be present in `.env` for migrations to work.

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

The Makefile automates common development and migration tasks.

**Available commands:**

- `make dev` — Run the app locally (with hot reload if air is installed)
- `make install` — Install dependencies
- `make migrate-create` — Create a migration using Atlas (see below)
- `make migrate-apply` — Apply migrations using DATABASE_URL from .env

**Usage example:**

```sh
make dev
make migrate-create
make migrate-apply
```

### Migration Create Command

`make migrate-create`

- Prompts for migration name (spaces replaced with underscores).
- Uses `DEV_DATABASE_URL` from `.env`.
- Errors if `DATABASE_URL` or `DEV_DATABASE_URL` is missing, or if both are the same.
- Example:
  ```sh
  make migrate-create
  ```

### Migration Apply Command

`make migrate-apply`

- Uses `DATABASE_URL` from `.env`.
- Errors if `DATABASE_URL` is missing.
- Example:
  ```sh
  make migrate-apply
  ```

## Generating Migrations

1. **Install Atlas CLI**
   See [Atlas docs](https://atlasgo.io/getting-started/) or use Homebrew:

   ```
   brew install ariga/tap/atlas
   ```

2. **Make changes to `schema.pg.hcl`**
   Edit this file to define your desired database schema.

3. **Create a migration**

   ```
   make migrate-create
   ```

4. **Apply migrations**
   ```
   make migrate-apply
   ```

## Development

- Edit code in `internal/` for features and logic.
- To change database structure, edit `schema.pg.hcl` (see Generating Migrations section).

---
