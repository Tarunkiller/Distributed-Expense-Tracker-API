# Distributed Expense Tracker API

A production-ready Expense Tracker backend built with **Golang**, **Gin Framework**, and **PostgreSQL**.

## Architecture & Tech Stack

- **Go 1.21**: Core language using Clean Architecture patterns
- **Gin**: High-performance HTTP web framework
- **PostgreSQL & GORM**: Relational database mapping, migrations, and pooling
- **JWT & Bcrypt**: Secure token-based authentication and password hashing
- **Docker & Docker Compose**: Containerization for seamless deployments
- **Swagger**: Auto-generated interactive API documentation

## Core Features

- **Authentication**: Signup, Login, and JWT generation with role-based middleware.
- **Expense CRUD**: Manage personal expenses with categories, descriptions, and dynamic filtering.
- **Team Management**: Create teams and share expenses among team members (Many-to-Many).
- **Analytics Engine**: GORM aggregations providing Monthly Summaries and Category Breakdowns.
- **Production-Ready Middleware**: Rate limiting, CORS, unified error handling, and logger middleware.

## Getting Started Locally

### Using Docker (Recommended)

1. Ensure Docker and Docker Compose are installed.
2. Spin up the entire stack (Go App + PostgreSQL):
   ```bash
   docker-compose up -d --build
   ```
3. The API will be available at `http://localhost:8080`.

### Without Docker

1. Ensure you have a running PostgreSQL instance.
2. Copy `.env.example` to `.env` and fill in your DB credentials.
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the app:
   ```bash
   go run cmd/api/main.go
   ```

## API Documentation (Swagger)

Once the application is running, you can access the interactive Swagger UI at:
`http://localhost:8080/swagger/index.html`

> **Note**: To regenerate swagger docs after modifying code, install the `swag` cli (`go install github.com/swaggo/swag/cmd/swag@latest`) and run `swag init -g cmd/api/main.go`.

## Deployment (Render/Railway)

1. Push this repository to GitHub.
2. Link the repository to your Render/Railway account.
3. Create a managed **PostgreSQL** database on the platform.
4. Set the environment variables exactly as in `.env.example` in the deployment dashboard, using the connection details from the managed Postgres database.
5. The `Dockerfile` provided will automatically build and deploy the Go binary perfectly.
