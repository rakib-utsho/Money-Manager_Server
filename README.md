# Money Manager Server

Simple backend API for managing personal finances and investments.

Prerequisites
- Go 1.26 or later

Quick start
1. Copy or create a `.env` with your database settings (see `config/db.go`).
2. Fetch dependencies:

```bash
go mod download
```

3. Run the server:

```bash
go run main.go
```

Project layout
- `main.go` — application entrypoint
- `config/` — database and configuration helpers
- `models/` — domain models (User, Transaction, Investment)
- `handlers/` — HTTP handlers
- `repository/`, `services/`, `router/` — application layers

Notes
- The project uses the `pgx`/`pq` Postgres drivers; update `config/db.go` to point at your database.
- Models include JSON tags for API responses; adjust DB tags as needed for your chosen data access library.
## Money Manager Server

A robust Go backend for the Money Manager project. It loads environment variables, connects to PostgreSQL, and exposes a structured HTTP API. The project follows a clean architecture pattern with defined layers for handlers, services, and repositories.

## Features

- RESTful HTTP server
- Environment variable configuration
- PostgreSQL database integration
- Clean architecture (Handlers, Services, Repositories)
- Middleware support
- Centralized routing

## Project Structure

```text
money-manager-server/
├── config/        # Database and application configuration
├── handlers/      # HTTP request handlers (Controllers)
├── middleware/    # HTTP middleware (Logging, Auth, etc.)
├── models/        # Application data structures and DB schemas
├── repository/    # Database queries and data access layer
├── router/        # Centralized HTTP route definitions
├── services/      # Business logic layer
├── main.go        # Application entry point
├── go.mod         # Go module dependencies
└── README.md      # Project documentation
```

## Requirements

- Go 1.20 or newer
- PostgreSQL database

## Environment Setup

Create a `.env` file in the repository root containing your database credentials and connection string:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=money_manager

DATABASE_URL=postgresql://postgres:yourpassword@localhost:5432/money_manager
```

## Installation

Install all required dependencies:

```bash
go mod tidy
```

## Running the Application

To start the server, run:

```bash
go run main.go
```

The server will initialize the database connection and start listening for incoming requests.

## Architecture

This project strictly adheres to a layered architecture to ensure separation of concerns, scalability, and maintainability:

1. **Router Layer (`/router`)**: Intercepts incoming requests and routes them to the appropriate handlers.
2. **Handler Layer (`/handlers`)**: Handles HTTP requests and responses. Parses payloads and validates inputs before calling the service layer.
3. **Service Layer (`/services`)**: Contains the core business logic of the application.
4. **Repository Layer (`/repository`)**: Interacts directly with the database to perform CRUD operations.
5. **Models Layer (`/models`)**: Defines the data structures used across the application.
6. **Middleware Layer (`/middleware`)**: Handles cross-cutting concerns like authentication, CORS, and request logging.
7. **Config Layer (`/config`)**: Manages the application configuration and database connection pooling.

