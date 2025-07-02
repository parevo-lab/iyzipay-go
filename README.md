# iyzipay-go

[![Go Reference](https://pkg.go.dev/badge/github.com/parevo/iyzipay-go.svg)](https://pkg.go.dev/github.com/parevo/iyzipay-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/parevo/iyzipay-go)](https://goreportcard.com/report/github.com/parevo/iyzipay-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

İyzipay Payment Gateway for Go - The official İyzipay API client for Go programming language.

İyzipay hesabı için [https://iyzico.com](https://iyzico.com) adresinden kayıt olabilirsiniz.

## Gereksinimler

Go 1.21 veya üzeri

## Kurulum

```bash
go get github.com/parevo/iyzipay-go
```

## Kullanım

### Başlatma

```go
package main

import (
    "github.com/parevo/iyzipay-go"
)

func main() {
    // Client oluşturma
    client := iyzipay.NewClient(&iyzipay.Config{
        APIKey:    "your-api-key",
        SecretKey: "your-secret-key",
        BaseURL:   "https://sandbox-api.iyzipay.com", // Test ortamı
    })
    
    // Ortam değişkenlerinden de kullanılabilir
    // IYZIPAY_API_KEY, IYZIPAY_SECRET_KEY, IYZIPAY_BASE_URL
    client := iyzipay.NewClientFromEnv()
}
```

### Örnek Ödeme

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/parevo/iyzipay-go"
)

func main() {
    client := iyzipay.NewClient(&iyzipay.Config{
        APIKey:    "sandbox-afXhZPW0MQlE4dCUUlHcEopnMBgXnAZI",
        SecretKey: "sandbox-wbwpzKIiplZxI3hh5ALI4FJyAcZKL6kq",
        BaseURL:   "https://sandbox-api.iyzipay.com",
    })

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

    ctx := context.Background()
    response, err := client.Payment.Create(ctx, request)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Payment Status: %s\n", response.Status)
    fmt.Printf("Payment ID: %s\n", response.PaymentID)
}
```

## API Servisleri

- **Payment**: Temel ödeme işlemleri
- **ThreedsInitialize**: 3D Secure ödeme başlatma
- **ThreedsPayment**: 3D Secure ödeme tamamlama
- **CheckoutForm**: Checkout form işlemleri
- **Card**: Kart saklama ve yönetimi
- **Refund**: İade işlemleri
- **Cancel**: İptal işlemleri
- **SubMerchant**: Alt bayi yönetimi
- **Subscription**: Abonelik işlemleri
- **BKM**: BKM Express ödeme
- **APM**: Alternatif ödeme yöntemleri

## Test Kartları

Başarılı ödeme simülasyonu için kullanılabilecek test kartları:

| Kart Numarası    | Banka        | Kart Tipi          |
|------------------|--------------|-------------------|
| 5528790000000008 | Halkbank     | Master Card       |
| 4766620000000001 | Denizbank    | Visa (Debit)      |
| 4603450000000000 | Denizbank    | Visa (Credit)     |
| 5400360000000003 | Garanti      | Master Card       |

## Lisans

MIT Lisansı - Detaylar için [LICENSE](LICENSE) dosyasına bakın.

## Katkıda Bulunma

Pull request'ler memnuniyetle karşılanır. Büyük değişiklikler için lütfen önce bir issue açın.

## Destek

Herhangi bir sorun için [GitHub Issues](https://github.com/parevo/iyzipay-go/issues) sayfasını kullanın.