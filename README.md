## Money Manager Server

A minimal Go backend for the Money Manager project. The server loads environment variables, connects to PostgreSQL using `pgx`, and exposes a small HTTP API used during development.

## Project Overview

Key points:

- Module-based Go project
- HTTP server listens on port `8000`
- Loads `.env` with `godotenv`
- Connects to PostgreSQL with a `pgxpool` connection pool
- Routes are registered under `/api/v1`
- Provides a health check endpoint

## Last Completed Work

- Implemented `main.go` to load environment and connect DB
- Added `internal/database` with `pgxpool` connection
- Created routing and a health check handler at `/api/v1/health`

## Project Structure

```text
go.mod
main.go
README.md
internal/
  database/
    database.go
  handlers/
    health.handaler.go
  routes/
    routes.go
```

## Getting Started

### Requirements

- Go (1.20+ recommended)
- PostgreSQL database

### Environment

Create a `.env` file in the repository root with at least:

```env
DATABASE_URL=postgres://user:password@localhost:5432/money_manager?sslmode=disable
```

### Install dependencies

```bash
go mod tidy
```

### Run locally

```bash
go run main.go
```

Server will be reachable at `http://localhost:8000`.

## API

### Health Check

- Endpoint: `GET /api/v1/health`
- Response body: `Api is healthy` (plain text)

## Troubleshooting

- If the server exits with a DB connection error, confirm `DATABASE_URL` and that Postgres is reachable.
- Missing `.env` will cause `godotenv.Load()` to return an error; ensure `.env` exists when running locally.

## Contributing

Small project — feel free to open issues or PRs. For changes affecting DB behavior, add clear setup steps in the README.

## License

This repository has no license specified. Add a `LICENSE` file if you want to share it publicly.
