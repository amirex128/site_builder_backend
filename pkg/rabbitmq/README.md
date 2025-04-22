# RabbitMQ Builder Pattern

This package provides a fluent builder pattern interface for working with RabbitMQ in Go.

## Features

- Simplified API for declaring exchanges, queues, and bindings
- Fluent interface for better readability
- Automatic acknowledgment/rejection of messages based on handler return value
- Support for configuring exchange and queue properties
- Publisher builder for easy message publishing with advanced options

## Usage Examples

### Setting up a Consumer

```go
// Initialize repository, use case, and consumer
repo := user_repo.NewUserRepository(db.DB)
useCase := user_use_case.NewUserUseCase(repo)
consumer := user_consumer.NewUserConsumer(useCase)

// Register consumer
err := client.Exchange("user_exchange").
    Queue("sms_queue").
    Type("direct").
    RoutingKey("user.sms").
    Config(true, false, false, false). // durable, autoDelete, exclusive, noWait
    Consume(consumer.SendSmsConsume)

if err != nil {
    // Handle error
}
```

### Using the Publisher Builder

```go
// Basic publishing
message := []byte(`{"phone": "+123456789", "message": "Hello!"}`)

err := client.Publisher("user_exchange").
    RoutingKey("user.sms").
    Type("direct").
    Config(true, false, false, false). // durable, autoDelete, internal, noWait
    Publish(context.Background(), message)

if err != nil {
    // Handle error
}
```

### Publishing JSON Messages

```go
// Create a message struct
smsRequest := struct {
    Phone   string `json:"phone"`
    Message string `json:"message"`
}{
    Phone:   "+123456789",
    Message: "Hello!",
}

// Publish JSON directly
err := client.Publisher("user_exchange").
    RoutingKey("user.sms").
    PublishJSON(context.Background(), smsRequest)

if err != nil {
    // Handle error
}
```

### Advanced Publishing Options

```go
// Advanced publishing with more options
err := client.Publisher("user_exchange").
    RoutingKey("user.sms").
    Type("direct").
    Config(true, false, false, false). // durable, autoDelete, internal, noWait
    DeliveryMode(true).                // persistent
    Priority(5).                       // priority 0-9
    Expiration("60000").              // TTL in milliseconds
    WithConfirms().                   // enable publisher confirms
    Publish(context.Background(), message)

if err != nil {
    // Handle error
}
```

### Batch Publishing

```go
// Create multiple messages
messages := [][]byte{
    []byte(`{"id": 1, "data": "First message"}`),
    []byte(`{"id": 2, "data": "Second message"}`),
    []byte(`{"id": 3, "data": "Third message"}`),
}

// Publish messages in batch
err := client.Publisher("user_exchange").
    RoutingKey("user.batch").
    PublishBatch(context.Background(), messages)

if err != nil {
    // Handle error
}
```

## API Reference

### Client Methods

- `Exchange(name string) *Builder`: Creates a new consumer builder with the specified exchange
- `Publisher(name string) *Publisher`: Creates a new publisher builder with the specified exchange
- `Close() error`: Closes the RabbitMQ connection

### Consumer Builder Methods

- `Queue(name string) *Builder`: Sets the queue name
- `Type(exchangeType string) *Builder`: Sets the exchange type (direct, fanout, topic, etc.)
- `RoutingKey(key string) *Builder`: Sets the routing key (defaults to queue name)
- `Config(durable, autoDelete, exclusive, noWait bool) *Builder`: Sets configuration options
- `PrefetchCount(count int) *Builder`: Sets the prefetch count for QoS
- `Args(args amqp.Table) *Builder`: Sets additional arguments for queue/exchange
- `Consume(handler ConsumerHandlerFunc) error`: Begins consuming messages

### Publisher Builder Methods

- `Type(exchangeType string) *Publisher`: Sets the exchange type
- `RoutingKey(key string) *Publisher`: Sets the routing key for publishing
- `Config(durable, autoDelete, internal, noWait bool) *Publisher`: Sets exchange configuration
- `ContentType(contentType string) *Publisher`: Sets message content type
- `DeliveryMode(persistent bool) *Publisher`: Sets persistence mode
- `Expiration(expiration string) *Publisher`: Sets message TTL
- `Priority(priority uint8) *Publisher`: Sets message priority (0-9)
- `WithConfirms() *Publisher`: Enables publisher confirms
- `Args(args amqp.Table) *Publisher`: Sets additional arguments for exchange
- `Publish(ctx context.Context, message []byte) error`: Publishes a message
- `PublishJSON(ctx context.Context, payload interface{}) error`: Marshals and publishes JSON
- `PublishBatch(ctx context.Context, messages [][]byte) error`: Publishes multiple messages

## Handler Function

The consumer handler function has the following signature:

```go
func(ctx context.Context, msg amqp.Delivery) error
```

- If the handler returns `nil`, the message is acknowledged (ack)
- If the handler returns an error, the message is rejected and requeued (nack) 