# Golang API sample

This repository contains a sample Golang API project with structured internal layers, configuration management, Swagger documentation, and Docker support.

## Project Structure

```
├── cmd/                   # Entry points for API and utility commands
│   └── api_sample
├── config/                # Configuration files
│   ├── .env
│   └── config.yaml
├── docs/                  # Swagger documentation (generated with `make swag-init`)
├── internal/              # Internal application logic
│   ├── app/               # Application initialization, config loading, dependency injection, server setup
│   ├── config/            # Configuration reading
│   ├── dto/               # Request/response structures and error DTOs
│   ├── model/             # Business entities
│   ├── gateway/           # Lower-level API clients for external services
│   ├── repository/        # Database access layer
│   ├── service/           # Business logic layer
│   ├── transport/         # HTTP and gRPC endpoints with versioning and templates
│   ├── server/            # HTTP server setup
│   └── middleware/        # Logging, CORS, error handling, security headers, API-KEY validation
├── pkg/                   # Shared packages
│   ├── logger/            # Logger interfaces and implementations (logrus, slog)
│   ├── errors/            # Application-specific errors
│   └── database/          # Database connections
├── Makefile               # Build and run automation
├── Dockerfile             # Docker image definition
└── docker-compose.yaml    # Docker Compose template
```

## Getting Started

### Prerequisites

- Golang 1.21.6
- Docker & Docker Compose (optional for containerized setup)

### Configuration

1. `.env` file for environment variables
2. `config.yaml` for structured application configuration

### Building & Running

You can run the application using either `make` commands or directly with `go`.

#### Using Makefile

```bash
# Build the project
make all

# Launch the compiled binary
make launch

# Run the project without building
make run

# Run tests
make test

# Initialize Swagger documentation
make swag-init

# Docker operations
make build       # Build Docker image
make up          # Run containers
make down        # Stop containers
make clean       # Clean binaries
```

#### Without Makefile

```bash
# Build the binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/api_sample cmd/api_sample/main.go

# Run the binary
./bin/api_sample

# Run directly with Go
go run cmd/api_sample/main.go

# Run tests
go test -cover -count=1 -v ./...

# Initialize Swagger docs
swag init -g internal/app/app.go
```

## Swagger Documentation

After running:

```bash
make swag-init
```

Swagger documentation will be generated in the `docs/` directory.

## Application Layers Overview

- **Transport** – HTTP/gRPC endpoints and versioning.
- **Service** – Internal business logic layer.
- **Repository / Gateway** – Lower-level database and external service interactions.

- **App** – Entry point that initializes configuration, injects dependencies, and starts the server with graceful shutdown.
- **Config** – Reads and parses configuration files.
- **DTO** – Request, response, and error data structures.
- **Model** – Business entities.
- **Middleware** – Logging, CORS, error handling, security headers, API-KEY validation.
- **Server** – HTTP server setup.

- **Pkg Layer** – Reusable components:
  - `logger/` – Logger interface and implementations (logrus, slog)
  - `errors/` – Application-specific errors
  - `database/` – DB connections

## Docker Support

Use the provided `Dockerfile` and `docker-compose.yaml` to run the service in containers:

```bash
# Build Docker image
docker-compose build

# Start service
docker-compose up

# Stop service
docker-compose down
```
