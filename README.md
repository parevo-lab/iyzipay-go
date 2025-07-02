# iyzipay-go

[![Go Reference](https://pkg.go.dev/badge/github.com/parevo-lab/iyzipay-go.svg)](https://pkg.go.dev/github.com/parevo-lab/iyzipay-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/parevo-lab/iyzipay-go)](https://goreportcard.com/report/github.com/parevo-lab/iyzipay-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-blue)](https://golang.org/dl/)
[![CI](https://github.com/parevo-lab/iyzipay-go/workflows/CI/badge.svg)](https://github.com/parevo-lab/iyzipay-go/actions/workflows/ci.yml)
[![Release](https://github.com/parevo-lab/iyzipay-go/workflows/Release/badge.svg)](https://github.com/parevo-lab/iyzipay-go/actions/workflows/release.yml)
[![CodeQL](https://github.com/parevo-lab/iyzipay-go/workflows/CodeQL/badge.svg)](https://github.com/parevo-lab/iyzipay-go/actions/workflows/codeql.yml)
[![Codecov](https://codecov.io/gh/parevo-lab/iyzipay-go/branch/main/graph/badge.svg)](https://codecov.io/gh/parevo-lab/iyzipay-go)

**iyzipay-go** is the official Go client library for İyzico Payment Gateway. This library provides a comprehensive, type-safe, and easy-to-use interface for integrating İyzico's payment processing capabilities into your Go applications.

## 🚀 Features

- **Complete API Coverage**: Support for all İyzico API endpoints including payments, 3D Secure, checkout forms, subscriptions, and more
- **Type-Safe**: Strongly typed request/response structures with comprehensive validation
- **Context Support**: Built with Go's context package for proper timeout and cancellation handling
- **Zero Dependencies**: Uses only Go's standard library - no external dependencies
- **Dual Authentication**: Supports both IYZWSv1 and IYZWSv2 authentication methods
- **PKI String Generation**: Automatic PKI string generation for secure API communication
- **Signature Verification**: Built-in HMAC-SHA256 signature verification for response validation
- **Comprehensive Examples**: Detailed examples for all major use cases
- **Production Ready**: Thoroughly tested with comprehensive unit and integration tests

## 📋 Requirements

- Go 1.21 or higher
- İyzico merchant account ([Sign up here](https://iyzico.com))

## 📦 Installation

```bash
go get github.com/parevo-lab/iyzipay-go
```

## 🏁 Quick Start

### Basic Configuration

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/parevo-lab/iyzipay-go"
)

func main() {
    // Create client with explicit configuration
    client := iyzipay.NewClient(&iyzipay.Config{
        APIKey:    "your-api-key",
        SecretKey: "your-secret-key",
        BaseURL:   "https://sandbox-api.iyzipay.com", // Use https://api.iyzipay.com for production
    })
    
    // Or create client from environment variables
    // Set IYZIPAY_API_KEY, IYZIPAY_SECRET_KEY, IYZIPAY_BASE_URL
    client := iyzipay.NewClientFromEnv()
}
```

### API Test

```go
ctx := context.Background()
response, err := client.APITest.Retrieve(ctx)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("API Status: %s\n", response.Status)
```

## 💳 Payment Operations

### Basic Payment

```go
request := &iyzipay.PaymentRequest{
    Locale:         iyzipay.LocaleTR,
    ConversationID: "123456789",
    Price:          "1.0",
    PaidPrice:      "1.2",
    Currency:       iyzipay.CurrencyTRY,
    Installment:    1,
    BasketID:       "B67832",
    PaymentChannel: iyzipay.PaymentChannelWeb,
    PaymentGroup:   iyzipay.PaymentGroupProduct,
    PaymentCard: &iyzipay.PaymentCard{
        CardHolderName: "John Doe",
        CardNumber:     "5528790000000008",
        ExpireMonth:    "12",
        ExpireYear:     "2030",
        CVC:            "123",
    },
    Buyer: &iyzipay.Buyer{
        ID:                  "BY789",
        Name:                "John",
        Surname:             "Doe",
        GsmNumber:           "+905350000000",
        Email:               "email@email.com",
        IdentityNumber:      "74300864791",
        LastLoginDate:       "2015-10-05 12:43:35",
        RegistrationDate:    "2013-04-21 15:12:09",
        RegistrationAddress: "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
        IP:                  "85.34.78.112",
        City:                "Istanbul",
        Country:             "Turkey",
        ZipCode:             "34732",
    },
    ShippingAddress: &iyzipay.Address{
        ContactName: "Jane Doe",
        City:        "Istanbul",
        Country:     "Turkey",
        Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
        ZipCode:     "34742",
    },
    BillingAddress: &iyzipay.Address{
        ContactName: "Jane Doe",
        City:        "Istanbul",
        Country:     "Turkey",
        Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
        ZipCode:     "34742",
    },
    BasketItems: []iyzipay.BasketItem{
        {
            ID:        "BI101",
            Name:      "Binocular",
            Category1: "Collectibles",
            Category2: "Accessories",
            ItemType:  iyzipay.BasketItemTypePhysical,
            Price:     "0.3",
        },
        {
            ID:        "BI102",
            Name:      "Game code",
            Category1: "Game",
            Category2: "Online Game Items",
            ItemType:  iyzipay.BasketItemTypeVirtual,
            Price:     "0.5",
        },
    },
}

response, err := client.Payment.Create(ctx, request)
if err != nil {
    log.Fatal(err)
}

if response.Status == "success" {
    fmt.Printf("Payment ID: %s\n", response.PaymentID)
    
    // Verify signature
    params := []string{
        response.PaymentID,
        response.Currency,
        response.BasketID,
        response.ConversationID,
        response.PaidPrice,
        response.Price,
    }
    expectedSignature := iyzipay.CalculateHMACSignature(params, client.GetConfig().SecretKey)
    if response.Signature == expectedSignature {
        fmt.Println("✓ Signature verified successfully")
    }
}
```

### 3D Secure Payment

```go
// Step 1: Initialize 3D Secure
initRequest := &iyzipay.PaymentRequest{
    // ... payment details (same as basic payment)
    CallbackURL: "https://www.yoursite.com/callback",
}

initResponse, err := client.ThreedsInitialize.Create(ctx, initRequest)
if err != nil {
    log.Fatal(err)
}

if initResponse.Status == "success" {
    // Display initResponse.ThreedsFormData to user for 3DS authentication
    fmt.Printf("3DS HTML Form: %s\n", initResponse.ThreedsFormData)
    
    // Step 2: Complete payment after 3DS authentication
    // (This would be called from your callback URL)
    paymentRequest := &iyzipay.ThreedsPaymentRequest{
        Locale:           iyzipay.LocaleTR,
        ConversationID:   "123456789",
        PaymentID:        initResponse.PaymentID,
        ConversationData: "conversation_data_from_callback",
    }
    
    paymentResponse, err := client.ThreedsPayment.Create(ctx, paymentRequest)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Final Payment Status: %s\n", paymentResponse.Status)
}
```

## 🛒 Checkout Form

The Checkout Form provides a hosted payment page that handles the entire payment flow:

```go
request := &iyzipay.CheckoutFormInitializeRequest{
    Locale:             iyzipay.LocaleTR,
    ConversationID:     "123456789",
    Price:              "1.0",
    PaidPrice:          "1.2",
    Currency:           iyzipay.CurrencyTRY,
    BasketID:           "B67832",
    PaymentGroup:       iyzipay.PaymentGroupProduct,
    CallbackURL:        "https://www.yoursite.com/callback",
    EnabledInstallments: []int{2, 3, 6, 9},
    Buyer:              buyer,         // Same buyer object as above
    ShippingAddress:    address,       // Same address object as above
    BillingAddress:     address,       // Same address object as above
    BasketItems:        basketItems,   // Same basket items as above
}

response, err := client.CheckoutForm.Initialize(ctx, request)
if err != nil {
    log.Fatal(err)
}

if response.Status == "success" {
    // Redirect user to response.PaymentPageURL
    // Or display response.CheckoutFormURL content
    fmt.Printf("Checkout Form Token: %s\n", response.Token)
    fmt.Printf("Payment Page URL: %s\n", response.PaymentPageURL)
}
```

## 💳 Card Management

### Store Card

```go
request := &iyzipay.CreateCardRequest{
    Locale:         iyzipay.LocaleTR,
    ConversationID: "123456789",
    Email:          "customer@email.com",
    ExternalID:     "customer_external_id",
    Card: &iyzipay.CardInformation{
        CardAlias:      "My Personal Card",
        CardHolderName: "John Doe",
        CardNumber:     "5528790000000008",
        ExpireMonth:    "12",
        ExpireYear:     "2030",
    },
}

response, err := client.Card.Create(ctx, request)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Card Token: %s\n", response.CardToken)
fmt.Printf("Card User Key: %s\n", response.CardUserKey)
```

### List Stored Cards

```go
request := &iyzipay.RetrieveCardListRequest{
    Locale:         iyzipay.LocaleTR,
    ConversationID: "123456789",
    CardUserKey:    "card_user_key_from_create_response",
}

response, err := client.Card.List(ctx, request)
if err != nil {
    log.Fatal(err)
}

for _, card := range response.CardDetails {
    fmt.Printf("Card: %s ending with %s\n", card.CardAlias, card.LastFourDigits)
}
```

### Payment with Stored Card

```go
request := &iyzipay.PaymentRequest{
    // ... other payment details
    PaymentCard: &iyzipay.PaymentCard{
        CardUserKey: "stored_card_user_key",
        CardToken:   "stored_card_token",
        // No need to provide card number, expiry, etc.
    },
}

response, err := client.Payment.Create(ctx, request)
```

## 🔄 Refunds and Cancellations

### Refund Payment

```go
request := &iyzipay.RefundRequest{
    Locale:               iyzipay.LocaleTR,
    ConversationID:       "123456789",
    PaymentTransactionID: "transaction_id_from_payment_response",
    Price:                "0.5", // Partial refund amount
    Currency:             iyzipay.CurrencyTRY,
    IP:                   "127.0.0.1",
    Reason:               iyzipay.RefundReasonBuyerRequest,
    Description:          "Customer requested refund",
}

response, err := client.Refund.Create(ctx, request)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Refund Status: %s\n", response.Status)
```

### Cancel Payment

```go
request := &iyzipay.CancelRequest{
    Locale:         iyzipay.LocaleTR,
    ConversationID: "123456789",
    PaymentID:      "payment_id_from_payment_response",
    IP:             "127.0.0.1",
    Reason:         iyzipay.RefundReasonBuyerRequest,
    Description:    "Customer requested cancellation",
}

response, err := client.Cancel.Create(ctx, request)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Cancel Status: %s\n", response.Status)
```

## 🏦 Sub Merchant Operations

### Create Sub Merchant

```go
request := &iyzipay.CreateSubMerchantRequest{
    Locale:                iyzipay.LocaleTR,
    ConversationID:        "123456789",
    SubMerchantExternalID: "SM001",
    SubMerchantType:       iyzipay.SubMerchantTypePersonal,
    Address:               "Business Address",
    ContactName:           "John",
    ContactSurname:        "Doe",
    Email:                 "merchant@email.com",
    GsmNumber:             "+905350000000",
    Name:                  "John's Store",
    IBAN:                  "TR180006200119000006672315",
    IdentityNumber:        "31300864726",
    Currency:              iyzipay.CurrencyTRY,
}

response, err := client.SubMerchant.Create(ctx, request)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Sub Merchant Key: %s\n", response.SubMerchantKey)
```

## 🔍 Utility Operations

### BIN Number Lookup

```go
request := &iyzipay.RetrieveBinNumberRequest{
    Locale:         iyzipay.LocaleTR,
    ConversationID: "123456789",
    BinNumber:      "552879", // First 6 digits of card
}

response, err := client.BinNumber.Retrieve(ctx, request)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Bank: %s\n", response.BankName)
fmt.Printf("Card Type: %s\n", response.CardType)
fmt.Printf("Card Association: %s\n", response.CardAssociation)
```

### Installment Information

```go
request := &iyzipay.RetrieveInstallmentInfoRequest{
    Locale:         iyzipay.LocaleTR,
    ConversationID: "123456789",
    BinNumber:      "552879",
    Price:          "100.0",
}

response, err := client.InstallmentInfo.Retrieve(ctx, request)
if err != nil {
    log.Fatal(err)
}

for _, detail := range response.InstallmentDetails {
    fmt.Printf("Bank: %s\n", detail.BankName)
    for _, price := range detail.InstallmentPrices {
        fmt.Printf("  %d installments: %s TL/month (Total: %s TL)\n",
            price.InstallmentNumber, price.InstallmentPrice, price.TotalPrice)
    }
}
```

## 🏗️ API Services

The library provides the following services:

| Service | Description |
|---------|-------------|
| `APITest` | API connectivity testing |
| `Payment` | Direct payment processing |
| `BasicPayment` | Simplified payment processing |
| `ThreedsInitialize` | 3D Secure initialization |
| `ThreedsPayment` | 3D Secure payment completion |
| `CheckoutForm` | Hosted checkout form |
| `Card` | Card storage and management |
| `Refund` | Payment refunds |
| `Cancel` | Payment cancellations |
| `SubMerchant` | Sub merchant management |
| `BKM` | BKM Express payments |
| `APM` | Alternative payment methods |
| `Subscription` | Subscription management |
| `InstallmentInfo` | Installment information |
| `BinNumber` | BIN number lookup |
| `PaymentItem` | Payment item management |
| `CrossBooking` | Cross booking operations |
| `RefundToBalance` | Refund to balance |
| `SettlementToBalance` | Settlement to balance |
| `UniversalCardStorage` | Universal card storage |

## 🌍 Constants and Enums

The library provides comprehensive constants for all enum values:

### Locales
```go
iyzipay.LocaleTR    // Turkish
iyzipay.LocaleEN    // English
```

### Currencies
```go
iyzipay.CurrencyTRY    // Turkish Lira
iyzipay.CurrencyUSD    // US Dollar
iyzipay.CurrencyEUR    // Euro
iyzipay.CurrencyGBP    // British Pound
// ... and more
```

### Payment Channels
```go
iyzipay.PaymentChannelWeb
iyzipay.PaymentChannelMobile
iyzipay.PaymentChannelMobileWeb
// ... and more
```

### Basket Item Types
```go
iyzipay.BasketItemTypePhysical
iyzipay.BasketItemTypeVirtual
```

### Refund Reasons
```go
iyzipay.RefundReasonDoublePayment
iyzipay.RefundReasonBuyerRequest
iyzipay.RefundReasonFraud
iyzipay.RefundReasonOther
```

## 🧪 Test Cards

Use these test cards in sandbox environment:

| Card Number | Bank | Type |
|-------------|------|------|
| 5528790000000008 | Halkbank | Master Card (Credit) |
| 4766620000000001 | Denizbank | Visa (Debit) |
| 4603450000000000 | Denizbank | Visa (Credit) |
| 5400360000000003 | Garanti | Master Card (Credit) |
| 4054180000000007 | Non-Turkish | Visa (Debit) |
| 5400010000000004 | Non-Turkish | Master Card (Credit) |

### Error Test Cards

| Card Number | Description |
|-------------|-------------|
| 4111111111111129 | Insufficient funds |
| 4129111111111111 | Do not honour |
| 4128111111111112 | Invalid transaction |
| 4127111111111113 | Lost card |
| 4126111111111114 | Stolen card |

## 🔒 Security

### Signature Verification

Always verify response signatures to ensure data integrity:

```go
if response.Signature != "" {
    params := []string{
        response.PaymentID,
        response.Currency,
        response.BasketID,
        response.ConversationID,
        response.PaidPrice,
        response.Price,
    }
    expectedSignature := iyzipay.CalculateHMACSignature(params, secretKey)
    if response.Signature == expectedSignature {
        // Signature is valid
    } else {
        // Invalid signature - possible tampering
    }
}
```

### PKI String Generation

The library automatically generates PKI strings for authentication. You can also generate them manually:

```go
pkiString := iyzipay.PKIString(requestObject)
```

## 🌟 Examples

Comprehensive examples are available in the `examples/` directory:

- **Basic Payment**: Simple payment processing
- **3D Secure Payment**: Secure payment with 3D authentication
- **Checkout Form**: Hosted payment form integration
- **Card Management**: Card storage and management
- **Refund/Cancel**: Payment refunds and cancellations

Run examples:
```bash
go run examples/basic_payment/main.go
go run examples/threeds_payment/main.go
go run examples/checkout_form/main.go
```

## 🧪 Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestPaymentCreate ./...
```

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Setup

```bash
git clone https://github.com/parevo-lab/iyzipay-go.git
cd iyzipay-go
go mod download
go test ./...
```

## 📝 Changelog

See [CHANGELOG.md](CHANGELOG.md) for a detailed list of changes.

## 🐛 Issues

If you encounter any issues, please [open an issue](https://github.com/parevo-lab/iyzipay-go/issues) with:

- Go version
- Library version
- Detailed description of the problem
- Code snippet to reproduce the issue

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Original İyzico Node.js SDK for API reference
- İyzico team for comprehensive API documentation
- Go community for excellent tooling and libraries

## 📞 Support

- 📧 Email: [info@parevo.com](mailto:support@parevo.com)
- 🌐 Website: [https://parevo.com](https://parevo.com)
- 📚 Documentation: [https://pkg.go.dev/github.com/parevo-lab/iyzipay-go](https://pkg.go.dev/github.com/parevo-lab/iyzipay-go)

---

Made with ❤️ by [Parevo Lab](https://github.com/parevo-lab)