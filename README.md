# Site Builder Backend - Clean Architecture with DDD

This project is a monolithic application built using Clean Architecture principles and Domain-Driven Design (DDD). It incorporates CQRS (Command Query Responsibility Segregation) for separating read and write operations.

## Architecture Overview

This application follows a Clean Architecture approach with concentric layers:

1. **Domain Layer** - Core business logic and domain entities
2. **Application Layer** - Use cases, CQRS handlers, and application services
3. **Infrastructure Layer** - External concerns like databases, message brokers, etc.
4. **Presentation Layer** - API controllers, CLI commands, etc.

### DDD Implementation

The project uses tactical DDD patterns:
- **Bounded Contexts** - Separate functional areas (Blog, Product, Order, etc.)
- **Aggregates** - Clusters of domain objects treated as a unit
- **Entities** - Objects with identity
- **Value Objects** - Objects defined by their attributes
- **Domain Events** - Events that domain experts care about
- **Repositories** - Abstraction for data access

### CQRS Implementation

- **Commands** - Modify state, return minimal response
- **Queries** - Read state, return DTOs, can be cached

## Directory Structure

```
├── cmd                         # Command-line applications
│   ├── api                     # API server
│   ├── consumer                # Message consumer
│   └── migration               # Database migrations
├── pkg                         # Core packages
│   ├── domain                  # Domain layer (DDD entities, aggregates, etc.)
│   │   ├── blog                # Blog bounded context
│   │   ├── order               # Order bounded context
│   │   ├── product             # Product bounded context
│   │   ├── site                # Site bounded context
│   │   ├── support             # Support bounded context
│   │   ├── user                # User bounded context
│   │   ├── drive               # Drive bounded context
│   │   ├── payment             # Payment bounded context
│   │   └── ai                  # AI bounded context
│   ├── application             # Application layer (use cases, services)
│   │   ├── command             # Command handlers (write operations)
│   │   ├── query               # Query handlers (read operations)
│   │   ├── dto                 # Data Transfer Objects
│   │   ├── service             # Domain services
│   │   └── event               # Domain event handlers
│   ├── infrastructure          # Infrastructure layer
│   │   ├── repository          # Repository implementations
│   │   ├── persistence         # Database connectors
│   │   │   ├── mysql           # MySQL adapter
│   │   │   ├── elastic         # Elasticsearch adapter
│   │   │   └── redis           # Redis adapter
│   │   ├── messaging           # Messaging infrastructure
│   │   │   ├── rabbitmq        # RabbitMQ implementation
│   │   │   └── outbox          # Outbox pattern implementation
│   │   └── cache               # Cache services
│   └── interface               # Interface layer
│       ├── api                 # API interfaces
│       │   └── rest            # REST API controllers
│       └── consumer            # Message consumers
├── internal                    # Internal packages
│   ├── config                  # Configuration
│   ├── middleware              # Middleware
│   └── validator               # Custom validators
├── scripts                     # Utility scripts
├── test                        # Integration and E2E tests
└── deployments                 # Deployment configurations
```

## Bounded Contexts and Aggregates

The application is divided into several bounded contexts:

1. **Blog Context**
   - Article (Aggregate Root)
   - Category
   - Media

2. **Order Context**
   - Order (Aggregate Root)
   - Basket
   - BasketItem
   - OrderItem
   - ReturnItem

3. **Product Context**
   - Product (Aggregate Root)
   - Category
   - ProductVariant
   - ProductAttribute
   - Discount
   - Coupon
   - ProductReview

4. **Site Context**
   - Site (Aggregate Root)
   - Page
   - HeaderFooter
   - DefaultTheme
   - Setting

5. **User Context**
   - User (Aggregate Root)
   - Customer
   - Role
   - Permission
   - Address
   - City
   - Province
   - Plan
   - UnitPrice

6. **Support Context**
   - Ticket (Aggregate Root)
   - Comment
   - CustomerTicket
   - CustomerComment

7. **Drive Context**
   - Storage (Aggregate Root)
   - FileItem

8. **Payment Context**
   - Payment (Aggregate Root)
   - Gateway
   - ParbadPayment
   - ParbadTransaction

9. **AI Context**
   - Credit (Aggregate Root)

## Context Mapping

The bounded contexts interact with each other through defined relationships:

1. **Blog → User**: Articles are created by Users (Customer Supplier)
2. **Product → Order**: Orders contain Products (Conformist)
3. **Order → Payment**: Payments are processed for Orders (Customer Supplier)
4. **Site → Blog/Product**: Sites display Blog and Product content (Open Host Service)
5. **User → Support**: Users create Support tickets (Partnership)
6. **Drive → All Contexts**: Drive provides file storage for all contexts (Shared Kernel)

## Technologies Used

- **Web Framework**: Gin 1.10.0
- **ORM**: GORM 1.25.12
- **Search Engine**: Elasticsearch v9
- **Cache**: Redis v9
- **Message Broker**: RabbitMQ
- **Validation**: GoValidator v11
- **CLI**: Cobra 1.9.1
- **Configuration**: Viper 1.20.1
- **Testing**: Testify 1.10.0
- **Logging**: Zap 1.27.0

## Getting Started

### Prerequisites

- Go 1.21+
- MySQL
- Elasticsearch
- Redis
- RabbitMQ

### Setup

1. Clone the repository
2. Create a `.env` file based on `.env.example`
3. Run:

```bash
go mod download
go run cmd/migration/main.go
go run cmd/api/main.go
```

## API Examples

### Create Article
```
POST /api/v1/blog/articles
```

### Get Article
```
GET /api/v1/blog/articles/{id}
```

## Testing

```bash
go test ./test/...
```

## License

MIT 