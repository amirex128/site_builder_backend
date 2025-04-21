# Command-Line Applications

This directory contains the main applications that serve as entry points to the system. Each subdirectory is a standalone application built using Cobra CLI.

## Structure

The cmd directory contains separate applications:

- `api` - REST API server
- `consumer` - Message consumer service
- `migration` - Database migration tool

## Components

### API Server

The API server provides REST endpoints for clients:
- HTTP server configuration
- Route registration using Gin
- Middleware configuration
- Dependency injection setup
- Application bootstrapping

### Consumer Service

The consumer service processes messages from RabbitMQ:
- Consumer configuration
- Message processing
- Error handling
- Retry logic

### Migration Tool

The migration tool manages database schema changes:
- Schema migration
- Data seeding
- Database initialization
- Version management

## Design Principles

- **Single Responsibility**: Each application has a clear, single purpose
- **Configuration**: Applications load their configuration from environment variables
- **Dependency Injection**: Services are wired together at startup
- **Graceful Shutdown**: Applications handle signals and shut down cleanly
- **CLI Interface**: Use Cobra for consistent command-line interfaces

## Usage

The applications in this directory are built using the Go standard build tools:

```bash
# Build API server
go build -o build/api ./cmd/api

# Build consumer service
go build -o build/consumer ./cmd/consumer

# Build migration tool
go build -o build/migration ./cmd/migration
``` 