
# Money Manager Server

Money Manager Server is a lightweight Go backend I built to learn backend engineering and to demonstrate a production-oriented project on my CV. It powers a simple personal finance API (users, transactions, investments) and showcases design choices I made while learning Go, HTTP APIs, and working with PostgreSQL.

**Project snapshot**
- **Purpose:** Showcase backend skills and practical experience building a layered Go service.
- **Tech:** Go, PostgreSQL, environment configuration, simple auth.

**Why I built this (my learning journey)**
- I wanted a compact, real-world project to practice idiomatic Go, separation of concerns, and database design. Building this project taught me how to structure a service into handlers, services, and repositories so logic stays testable and maintainable.
- I iterated on authentication, DB connection management, and routing while focusing on clear, resume-ready deliverables.

**What I learned**
- Go project layout and dependency management with `go mod`.
- Designing RESTful APIs and wiring routes in [router/router.go](router/router.go).
- Database connection patterns and migrations (see [config/db.go](config/db.go)).
- Implementing a service layer and repository pattern for testability.
- Handling environment configuration and secure secrets with `.env` files.

**Key features**
- RESTful endpoints for user auth and basic finance models.
- Layered architecture: handlers → services → repository.
- Centralized DB connection in [config/db.go](config/db.go).
- Authentication scaffolding in [handlers/auth.go](handlers/auth.go) and [services/auth.service.go](services/auth.service.go).

**How to run (quickstart)**
1. Create a `.env` file in the project root with your DB credentials (example below).

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=money_manager
```

2. Install deps and run the server:

```bash
go mod tidy
go run main.go
```

The server listens on port 8080 by default.

**Architecture & where to look**
- `main.go` — app bootstrap and server start
- [config/db.go](config/db.go) — DB connection setup
- [router/router.go](router/router.go) — route wiring
- [handlers/](handlers/) — HTTP handlers (see [handlers/auth.go](handlers/auth.go))
- [services/](services/) — business logic (see [services/auth.service.go](services/auth.service.go))
- [repository/](repository/) — DB access (see [repository/user.repo.go](repository/user.repo.go))

**Resume-ready bullets**
- Built a production-oriented backend in Go: designed REST APIs, implemented authentication, and integrated PostgreSQL persistence.
- Implemented layered architecture (handlers → services → repository) to improve maintainability and testability.
- Managed environment-based configuration and DB connections for a reliable local development setup.

**How I built it (short timeline)**
- Prototype: sketched endpoints and models.
- Iteration: implemented handlers and repository layer, then refactored into services for clearer responsibilities.
- Polish: added dotenv support, README, and resume-ready notes.

**Next steps / future improvements**
- Add unit and integration tests and include CI (GitHub Actions).
- Add a `.env.example`, database migrations, and deployment instructions.
- Add a small Postman collection or example curl commands and a demo screenshot/gif.

If you want, I can:
- add a `.env.example` and `CONTRIBUTING.md`;
- draft 3 tailored resume bullets for a specific job posting;
- add example curl requests or a Postman collection.

---

If you'd like a more tailored README (shorter for GitHub profile, longer for portfolio), tell me which tone you prefer and I will adjust it.


