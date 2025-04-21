# Internal

This directory contains application-specific internal packages that are not part of the core domain logic but are needed for the application to function correctly.

## Structure

The internal directory is organized by functionality:

- `config` - Configuration management
- `middleware` - HTTP middleware components
- `validator` - Custom validators

## Components

### Config

Configuration management components:
- Environment variable loading
- Configuration structures
- Defaults and validation
- Uses Viper for configuration loading

### Middleware

HTTP middleware components for Gin:
- Authentication
- Authorization
- Logging
- Request ID
- CORS handling
- Rate limiting
- Error handling

### Validator

Custom validators for request validation:
- Field validation rules
- Custom validation logic
- Validation error formatting
- Extensions for GoValidator

## Design Principles

- **Application-Specific**: Contains code specific to this application
- **Not Exported**: Not meant to be imported by other projects
- **Support Code**: Provides support for the main application code
- **Cross-Cutting Concerns**: Handles aspects that cut across layers

## Usage

The internal packages should be small, focused, and serve specific purposes for the application. They can be used by any layer but should not contain core business logic. 