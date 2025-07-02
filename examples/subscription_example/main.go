package main

import (
	"context"
	"fmt"
	"log"

	"github.com/parevo-lab/iyzipay-go"
)

func main() {
	// İyzipay client oluşturma
	client := iyzipay.NewClient(&iyzipay.Config{
		APIKey:    "sandbox-afXhZPW0MQlE4dCUUlHcEopnMBgXnAZI",
		SecretKey: "sandbox-wbwpzKIiplZxI3hh5ALI4FJyAcZKL6kq",
		BaseURL:   "https://sandbox-api.iyzipay.com",
	})

	ctx := context.Background()

	// Abonelik başlatmak için gerekli bilgiler
	request := &iyzipay.CreateSubscriptionInitRequest{
		Locale:                    iyzipay.LocaleTR,
		ConversationID:            "123456789",
		PricingPlanReferenceCode:  "plan_ref_code", // Buraya İyzipay panelinden aldığınız plan kodunu girin
		SubscriptionInitialStatus: iyzipay.SubscriptionInitialStatusActive,
		Customer: &iyzipay.SubscriptionCustomer{
			Name:           "Ahmet",
			Surname:        "Can",
			IdentityNumber: "12345678901",
			Email:          "ahmet@email.com",
			GsmNumber:      "+905555555555",
			BillingAddress: &iyzipay.SubscriptionAddress{
				Address:     "Adres",
				ZipCode:     "34000",
				ContactName: "Ahmet Can",
				City:        "İstanbul",
				Country:     "Türkiye",
				District:    "Kadıköy",
			},
			ShippingAddress: &iyzipay.SubscriptionAddress{
				Address:     "Adres",
				ZipCode:     "34000",
				ContactName: "Ahmet Can",
				City:        "İstanbul",
				Country:     "Türkiye",
				District:    "Kadıköy",
			},
		},
		PaymentCard: &iyzipay.SubscriptionCard{
			CardHolderName: "Ahmet Can",
			CardNumber:     "5528790000000008",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
			CVC:            "123",
		},
	}

	response, err := client.Subscription.Initialize(ctx, request)
	if err != nil {
		log.Fatalf("Abonelik başlatılamadı: %v", err)
	}

	fmt.Println("Abonelik Başlatıldı!")
	fmt.Printf("SubscriptionReferenceCode: %s\n", response.SubscriptionReferenceCode)
	fmt.Printf("Durum: %s\n", response.SubscriptionStatus)
	fmt.Printf("Başlangıç Tarihi: %s\n", response.StartDate)
	fmt.Printf("Bitiş Tarihi: %s\n", response.EndDate)
}
