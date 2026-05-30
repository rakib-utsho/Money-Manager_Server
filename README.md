# Money Manager Server — CV-ready Project

Money Manager Server is a concise, production-oriented Go backend for a personal finance app. It's designed to showcase backend engineering skills that recruiters and hiring managers look for: building RESTful APIs, secure authentication, relational data modeling, and working with production-grade tools (Go, PostgreSQL, environment configuration).

## Why this project helps you stand out

- Demonstrates full-stack backend skills using a statically-typed language (Go).
- Shows experience designing and implementing REST APIs and DB schemas.
- Highlights knowledge of authentication, environment configuration, and project structure.
- Easy to extend for real-world features (transactions, investments, reporting).

## Key features

- RESTful HTTP server with clear routing and handlers.
- PostgreSQL integration with centralized DB connection management.
- Authentication scaffolding (see `handlers/auth.go` and `services/auth.service.go`).
- Layered project structure: `handlers`, `services`, `repository`, `models`, and `router`.

## Suggested CV / resume bullets (copy-paste and customize)

- Built a production-ready backend in Go for a personal finance app: implemented RESTful APIs, secure authentication, and PostgreSQL persistence.
- Designed and implemented relational data models for users, transactions, and investments; implemented repository patterns for maintainable DB access.
- Integrated environment-based configuration and connection pooling for PostgreSQL; automated dependency management with `go mod`.
- Architected layered code separation (`handlers` → `services` → `repository`) to improve testability and maintainability.

## Tech stack / skills demonstrated

- Language: Go
- Database: PostgreSQL
- Libraries & concepts: godotenv, database/sql (or preferred driver), REST API design, middleware, authentication patterns
- Development skills: modular project layout, environment config, dependency management

## Quickstart (run locally)

1. Create a `.env` file in the project root with your DB credentials:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=money_manager
```

2. Install dependencies and run:

```bash
go mod tidy
go run main.go
```

Server runs on port `8080` by default.

## How to make it even stronger for applications

- Add a short demo or GIF showing the API in action (Postman or curl requests). Recruiters love visible evidence.
- Add unit and integration tests (examples: handler tests, repository tests using a test DB or Docker). Include test coverage in the README.
- Add CI (GitHub Actions) to run tests and linting on push/PR.
- Deploy a live demo (Heroku, Fly, Render, or Docker + VPS) and include the link in the README.

## Architecture & where to look in the code

- `main.go` — app bootstrap and server start
- `config/db.go` — DB connection setup
- `router/router.go` — route wiring
- `handlers/` — HTTP handler implementations (see `handlers/auth.go`)
- `services/` — business logic (see `services/auth.service.go`)
- `repository/` — DB access (see `repository/user.repo.go`)

## Next steps I can help with

- Add a `LICENSE` and a small `.env.example` file.
- Create copy-ready screenshots or a demo script you can link in your CV.
- Draft personalized resume bullets tailored to a specific job description.

If you want, I can now add a `.env.example` and a `CONTRIBUTING.md`, or draft 3 tailored CV bullets for a role you're applying to.

# Money Manager Server

Money Manager Server is a lightweight Go backend for a personal finance application. It starts an HTTP server, loads configuration from a `.env` file, connects to a PostgreSQL database, and exposes basic endpoints for verifying the API and handling user/auth operations.

## Quickstart

1. Copy or clone the repository and change into the project directory.
2. Create a `.env` file in the project root (example below).
3. Install dependencies and build/run:

```bash
go mod tidy
go run main.go
```

The server listens on port `8080` by default.

## Prerequisites

- Go 1.26.2 or newer
- PostgreSQL database

## Environment Variables

Create a `.env` file at the project root with your DB credentials. The application reads these values in `config/db.go`.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=money_manager
```

## What this project does

- Loads environment variables with `godotenv`.
- Connects to PostgreSQL using `config.ConnectDB()` and exposes a shared `config.DB` handle.
- Serves a health/root endpoint at `GET /` which returns `Money Manager API is running!`.
- Contains scaffolding for models, repositories, handlers, services, router and middleware.

## Available files (overview)

- `main.go` - application entrypoint and server bootstrap.
- `config/` - database connection and shared config (see [config/db.go](config/db.go)).
- `handlers/` - HTTP handlers (see [handlers/auth.go](handlers/auth.go)).
- `models/` - data models for `User`, `Transaction`, and `Investment`.
- `repository/` - database access helpers (see [repository/user.repo.go](repository/user.repo.go)).
- `router/` - HTTP route wiring (see [router/router.go](router/router.go)).
- `services/` - business logic layer (see [services/auth.service.go](services/auth.service.go)).

## Development notes

- To add routes, implement handlers under `handlers/` and wire them in `router/router.go`.
- Database access should go through the `repository/` layer to keep DB logic centralized.

## Contributing

Contributions are welcome. Open an issue or submit a pull request with a clear description of changes.

## License

This project does not include a license file. Add a `LICENSE` if you intend to publish with an open-source license.


