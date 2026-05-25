## Money Manager Server

A small Go backend for the Money Manager project. It loads environment variables, connects to PostgreSQL through `pgxpool`, and exposes a basic HTTP API under `/api/v1`.

## Features

- HTTP server on port `8000`
- Environment loading with `godotenv`
- PostgreSQL connection via `pgxpool`
- JSON API responses
- Health check and user registration routes

## Project Structure

```text
go.mod
main.go
README.md
internal/
  database/
    database.go
  handlers/
    health.handler.go
    user.handler.go
  models/
    user.model.go
  routes/
    routes.go
  utils/
    json.go
```

## Requirements

- Go 1.26 or newer
- PostgreSQL database

## Environment

Create a `.env` file in the repository root with:

```env
DATABASE_URL=postgres://user:password@localhost:5432/money_manager?sslmode=disable
```

## Install

```bash
go mod tidy
```

## Run

```bash
go run main.go
```

The server listens on `http://localhost:8000`.

## API Endpoints

### Health Check

- `GET /api/v1/health`
- Returns a JSON response with `success`, `message`, and optional `data`

### Register User

- `POST /api/v1/users/register`
- Request body:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secret"
}
```

- Returns the created user record with `id`, `created_at`, and `updated_at`

## Notes

- The app exits if `.env` cannot be loaded.
- The database connection must succeed before the server starts.
