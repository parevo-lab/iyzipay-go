# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-01-XX

### Added
- 🎉 Initial release of iyzipay-go
- Complete İyzico Payment Gateway API coverage
- Support for all major payment operations:
  - Basic payments
  - 3D Secure payments
  - Checkout form integration
  - Card storage and management
  - Payment refunds and cancellations
  - Sub merchant operations
  - BKM Express payments
  - Alternative payment methods (APM)
  - Subscription management
  - Cross booking operations
  - Settlement and balance operations
- Dual authentication support (IYZWSv1 and IYZWSv2)
- Automatic PKI string generation
- HMAC-SHA256 signature verification
- Context-aware HTTP client with proper timeout handling
- Comprehensive type-safe request/response structures
- Zero external dependencies - uses only Go standard library
- Extensive unit and integration tests
- Detailed examples for all major use cases
- Support for both sandbox and production environments
- Environment variable configuration support
- BIN number lookup functionality
- Installment information retrieval
- Comprehensive error handling
- Production-ready logging and monitoring capabilities

### Security
- Implemented secure PKI string generation
- Added signature verification for all responses
- Secure credential handling with environment variable support
- Proper handling of sensitive card data

### Documentation
- Comprehensive README with detailed examples
- Complete API documentation
- Code examples for all major operations
- Test card information for development
- Security best practices guide
- Migration guide from other SDKs

### Testing
- Unit tests with >90% coverage
- Integration tests with mock servers
- Example applications demonstrating real-world usage
- Continuous integration setup

## [Unreleased]

### Planned Features
- Webhook signature verification helpers
- Rate limiting support
- Enhanced logging and metrics
- Advanced retry mechanisms
- Bulk operations support
- Additional payment method integrations