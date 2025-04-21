# Interface Layer

This directory contains components that handle interactions with external systems and users, forming the outermost layer of Clean Architecture.

## Structure

The interface layer is organized by entry point type:

- `api` - API interfaces
  - `rest` - REST API controllers using Gin
- `consumer` - Message consumers for handling asynchronous requests

## Components

### REST API Controllers

REST API controllers handle HTTP requests and responses:
- Route configuration using Gin
- Request parsing and validation
- Response formatting
- Error handling
- Authentication and authorization
- Input validation using GoValidator

### Message Consumers

Message consumers process messages from message brokers:
- Message deserialization
- Consumer logic
- Error handling and retry
- Dead letter handling

## Design Principles

- **Controllers**: Handle HTTP request/response concerns
- **Consumers**: Process asynchronous messages
- **Adapters**: Adapt external requests to application use cases
- **Thin Controllers**: Minimal logic, delegate to application services
- **Input Validation**: Validate and sanitize all external input
- **Error Handling**: Consistent error response formats

## Usage

The interface layer depends on the application layer but not on the infrastructure layer directly. It should be kept thin, with minimal business logic, delegating to application services for processing. 