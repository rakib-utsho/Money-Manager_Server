## Money Manager Server

A small Go backend for the Money Manager project. The server loads environment variables, connects to PostgreSQL with `pgx`, and exposes a health check endpoint.

## Project Overview

Current features:

- Go module-based project structure
- HTTP server running on port `8000`
- Environment variable loading with `godotenv`
- PostgreSQL connection through `pgxpool`
- API routing under `/api/v1`
- Health check endpoint for basic server verification

## Last Completed Work

The last work completed in this repository was:

- Added the application entry point in `main.go`
- Wired environment loading and database connection startup
- Registered the API route group
- Added a health check handler at `/api/v1/health`

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

## Setup

### Requirements

- Go installed locally
- PostgreSQL database
- `.env` file with `DATABASE_URL`

### Environment Variables

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://user:password@localhost:5432/money_manager?sslmode=disable
```

### Install Dependencies

```bash
go mod tidy
```

## Run the Server

```bash
go run main.go
```

The server starts on `http://localhost:8000`.

## API Endpoints

### Health Check

- `GET /api/v1/health`

Response:

```text
Api is healthy
```

## Notes

- The database connection is established during startup.
- If `DATABASE_URL` is missing or invalid, the server will stop on launch.
