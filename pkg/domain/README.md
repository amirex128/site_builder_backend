# Domain Layer

This directory contains the core business logic and domain models following Domain-Driven Design (DDD) principles. It represents the innermost layer of Clean Architecture.

## Structure

The domain layer is organized by bounded contexts, each representing a distinct business capability:

- `blog` - Blog articles, categories, and related entities
- `order` - Order processing, baskets, and order items
- `product` - Products, categories, variants, and pricing
- `site` - Website configuration, pages, headers/footers
- `support` - Support tickets and communication
- `user` - User accounts, roles, and permissions
- `drive` - File storage and management
- `payment` - Payment processing and gateways
- `ai` - AI-related functionalities and credits

## Domain Components

Each bounded context contains:

- **Entities**: Domain objects with identity (e.g., Article, Order)
- **Value Objects**: Immutable objects without identity (e.g., Address, Money)
- **Aggregates**: Clusters of entities and value objects with a root entity
- **Domain Events**: Events that domain experts care about
- **Repository Interfaces**: Contracts for data access
- **Domain Services**: Domain logic that doesn't belong to any entity

## Design Principles

- **Rich Domain Model**: Business logic lives in the domain objects
- **Immutability**: Value objects are immutable
- **Encapsulation**: Internal state is protected
- **Domain Events**: For cross-aggregate communication
- **No Infrastructure Dependencies**: This layer should have no external dependencies

## Usage

The domain layer should not depend on any other layer of the application. It contains the core business rules and is independent of any external concerns like databases or UI. 