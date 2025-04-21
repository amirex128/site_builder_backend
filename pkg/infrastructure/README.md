# Infrastructure Layer

This directory contains the infrastructure components that provide concrete implementations of interfaces defined in the domain and application layers, following Clean Architecture principles.

## Structure

The infrastructure layer is organized by technical concern:

- `repository` - Repository implementations
- `persistence` - Database adapters
  - `mysql` - MySQL database adapter using GORM
  - `elastic` - Elasticsearch adapter
  - `redis` - Redis adapter
- `messaging` - Messaging infrastructure
  - `rabbitmq` - RabbitMQ implementation
  - `outbox` - Outbox pattern implementation
- `cache` - Cache services

## Components

### Repositories

Repository implementations provide concrete data access for domain entities:
- Implement repository interfaces defined in the domain layer
- Abstract database operations
- Handle ORM mapping
- Manage transactions

### Persistence

Database adapters manage connections and operations for various storage technologies:
- Connection setup and configuration
- Query execution
- Transaction management
- Mapping between domain entities and database models

### Messaging

Messaging components handle communication with message brokers:
- Message publishing
- Consumer implementations
- Message serialization/deserialization
- Outbox pattern for reliable messaging

### Cache

Cache services provide data caching functionality:
- Cache management
- Cache invalidation strategies
- Distributed caching support
- Cache serialization

## Design Principles

- **Adapters**: Implements interfaces defined in inner layers
- **Infrastructure Concerns**: Deals with technical implementation details
- **External Services**: Interfaces with external systems
- **Dependency Injection**: Provides concrete implementations for higher layers
- **Transaction Management**: Handles database transactions

## Usage

The infrastructure layer depends on the domain and application layers but should not be depended upon by them. It adapts external concerns to the needs of the inner layers. 