# Contributing to iyzipay-go

Thank you for your interest in contributing to iyzipay-go! This document provides guidelines and information for contributors.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)
- [Code Style](#code-style)
- [Security](#security)

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Set up the development environment
4. Make your changes
5. Test your changes
6. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git
- Make (optional, but recommended)

### Setup

```bash
# Clone your fork
git clone https://github.com/your-username/iyzipay-go.git
cd iyzipay-go

# Add the original repository as upstream
git remote add upstream https://github.com/parevo-lab/iyzipay-go.git

# Install development tools
make install-tools

# Download dependencies
make deps

# Run tests to ensure everything works
make test
```

## Making Changes

### Before You Start

1. Check if there's already an issue for your change
2. If not, create an issue to discuss the change
3. Fork the repository and create a feature branch

### Branch Naming

Use descriptive branch names:
- `feature/add-webhook-support`
- `bugfix/fix-timeout-handling`
- `docs/update-readme`

### Commit Messages

Follow conventional commit format:
```
type(scope): description

body (optional)

footer (optional)
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:
```
feat(payment): add webhook signature verification

fix(client): handle timeout errors properly

docs: update installation instructions
```

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with race detection
make test-race

# Run tests with coverage
make test-cover

# Run benchmarks
make bench
```

### Writing Tests

- Write unit tests for new functionality
- Ensure test coverage remains above 80%
- Use table-driven tests where appropriate
- Mock external dependencies

Example test:
```go
func TestPaymentCreate(t *testing.T) {
    tests := []struct {
        name    string
        request *PaymentRequest
        want    *PaymentResponse
        wantErr bool
    }{
        {
            name: "successful payment",
            request: &PaymentRequest{
                // test data
            },
            want: &PaymentResponse{
                Status: "success",
            },
            wantErr: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

## Pull Request Process

1. **Update Documentation**: Ensure the README.md and any relevant documentation is updated
2. **Add Tests**: Include tests for new functionality
3. **Update Changelog**: Add your changes to CHANGELOG.md
4. **Check CI**: Ensure all CI checks pass
5. **Request Review**: Request review from maintainers

### Pull Request Checklist

- [ ] I have read the CONTRIBUTING document
- [ ] My code follows the code style of this project
- [ ] My change requires a change to the documentation
- [ ] I have updated the documentation accordingly
- [ ] I have added tests to cover my changes
- [ ] All new and existing tests passed
- [ ] My changes generate no new warnings

## Code Style

### General Guidelines

- Follow standard Go conventions
- Use `gofmt` and `goimports` for formatting
- Run `golangci-lint` before submitting
- Keep functions small and focused
- Use meaningful variable and function names

### Specific Guidelines

1. **Error Handling**:
   ```go
   // Good
   if err != nil {
       return fmt.Errorf("failed to create payment: %w", err)
   }

   // Bad
   if err != nil {
       return err
   }
   ```

2. **Context Usage**:
   ```go
   // Good
   func (s *Service) Create(ctx context.Context, req *Request) (*Response, error) {
       // implementation
   }

   // Bad
   func (s *Service) Create(req *Request) (*Response, error) {
       // implementation
   }
   ```

3. **Struct Tags**:
   ```go
   type PaymentRequest struct {
       Amount   string `json:"amount"`
       Currency string `json:"currency"`
   }
   ```

### Code Quality Tools

Run these before submitting:

```bash
# Format code
make fmt

# Lint code
make lint

# Vet code
make vet

# Security scan
make security

# All quality checks
make all
```

## Security

### Reporting Security Issues

Please do not report security vulnerabilities through public GitHub issues. Instead, send an email to [security@parevo.com](mailto:security@parevo.com).

### Security Guidelines

- Never commit API keys, tokens, or sensitive data
- Use context timeouts for all HTTP requests
- Validate all input parameters
- Follow secure coding practices

## API Design Guidelines

### Request/Response Structures

```go
// Good: Clear, typed structures
type PaymentRequest struct {
    Locale         string       `json:"locale"`
    ConversationID string       `json:"conversationId"`
    Amount         string       `json:"amount"`
    Currency       string       `json:"currency"`
    PaymentCard    *PaymentCard `json:"paymentCard"`
}

// Bad: Using interface{} or maps
type PaymentRequest map[string]interface{}
```

### Service Methods

```go
// Good: Context-aware, error handling
func (s *PaymentService) Create(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
    // implementation
}

// Bad: No context, unclear return types
func (s *PaymentService) Create(req interface{}) interface{} {
    // implementation
}
```

## Documentation

### Code Documentation

- Add godoc comments for all public functions and types
- Include examples in documentation
- Keep documentation up to date with code changes

Example:
```go
// PaymentService handles payment operations.
type PaymentService struct {
    client *Client
}

// Create processes a payment request.
//
// Example:
//   request := &iyzipay.PaymentRequest{
//       Locale: iyzipay.LocaleTR,
//       Amount: "10.00",
//   }
//   response, err := client.Payment.Create(ctx, request)
func (s *PaymentService) Create(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
    // implementation
}
```

### README Updates

When adding new features, update:
- Installation instructions (if needed)
- Usage examples
- API documentation
- Feature list

## Release Process

Releases are automated through GitHub Actions when tags are pushed:

1. Update CHANGELOG.md
2. Create and push a tag:
   ```bash
   git tag -a v1.1.0 -m "Release v1.1.0"
   git push origin v1.1.0
   ```
3. GitHub Actions will automatically:
   - Run all tests
   - Build release artifacts
   - Create a GitHub release
   - Notify about the release

## Questions?

If you have questions about contributing, feel free to:
- Open an issue for discussion
- Email us at [support@parevo.com](mailto:support@parevo.com)
- Check existing issues and discussions

Thank you for contributing to iyzipay-go! 🎉