# Money Manager Server

Money Manager Server is a Go backend for a personal finance app. The current codebase boots an HTTP server, loads environment variables from `.env`, connects to PostgreSQL, and exposes a simple root endpoint to confirm the API is running.

## Requirements

- Go 1.26.2 or newer
- PostgreSQL database

## Current Behavior

- Loads environment variables with `godotenv`
- Opens a PostgreSQL connection through `config.ConnectDB()`
- Serves `GET /` with the response `Money Manager API is running!`
- Uses a shared database handle in `config.DB`

## Environment Variables

Create a `.env` file in the project root with the database credentials used by `config/db.go`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=money_manager
```

## Install

```bash
go mod tidy
```

## Run

```bash
go run main.go
```

The server listens on port `8080`.

## Project Layout

- `main.go` - application entry point and HTTP bootstrap
- `config/` - database connection setup
- `models/` - data models for users, transactions, and investments
- `repository/` - database access helpers
- `handlers/`, `services/`, `router/`, `middleware/` - reserved for future API layers

## Data Models

The repository currently defines models for:

- `User`
- `Transaction`
- `Investment`

## Notes

- The repository layer currently includes user database helpers in `repository/user.repo.go`.
- If you add API routes, document them here so the README stays aligned with the code.

