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

	// API Test
	fmt.Println("=== API Test ===")
	testResponse, err := client.APITest.Retrieve(context.Background())
	if err != nil {
		log.Printf("API Test Error: %v", err)
	} else {
		fmt.Printf("API Test Status: %s\n", testResponse.Status)
		fmt.Printf("API Test Locale: %s\n", testResponse.Locale)
	}

	// Temel ödeme işlemi
	fmt.Println("\n=== Temel Ödeme ===")
	paymentRequest := &iyzipay.PaymentRequest{
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
			{
				ID:        "BI103",
				Name:      "USB",
				Category1: "Electronics",
				Category2: "USB / Cable",
				ItemType:  iyzipay.BasketItemTypePhysical,
				Price:     "0.2",
			},
		},
	}

	ctx := context.Background()
	response, err := client.Payment.Create(ctx, paymentRequest)
	if err != nil {
		log.Printf("Payment Error: %v", err)
		return
	}

	fmt.Printf("Payment Status: %s\n", response.Status)
	if response.Status == "success" {
		fmt.Printf("Payment ID: %s\n", response.PaymentID)
		fmt.Printf("Fraud Status: %d\n", response.FraudStatus)
		fmt.Printf("Card Association: %s\n", response.CardAssociation)
		fmt.Printf("Card Family: %s\n", response.CardFamily)
		fmt.Printf("Card Type: %s\n", response.CardType)
		fmt.Printf("BIN Number: %s\n", response.BinNumber)
		fmt.Printf("Last Four Digits: %s\n", response.LastFourDigits)
		fmt.Printf("Auth Code: %s\n", response.AuthCode)

		// İmza doğrulama
		if response.Signature != "" {
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
			} else {
				fmt.Println("✗ Signature verification failed")
			}
		}

		// Item transactions detayları
		if len(response.ItemTransactions) > 0 {
			fmt.Println("\nItem Transactions:")
			for i, item := range response.ItemTransactions {
				fmt.Printf("  Item %d:\n", i+1)
				fmt.Printf("    Item ID: %s\n", item.ItemID)
				fmt.Printf("    Payment Transaction ID: %s\n", item.PaymentTransactionID)
				fmt.Printf("    Transaction Status: %d\n", item.TransactionStatus)
				fmt.Printf("    Price: %s\n", item.Price)
				fmt.Printf("    Paid Price: %s\n", item.PaidPrice)
			}
		}
	} else {
		fmt.Printf("Payment Failed - Error Code: %s\n", response.ErrorCode)
		fmt.Printf("Error Message: %s\n", response.ErrorMessage)
		fmt.Printf("Error Group: %s\n", response.ErrorGroup)
	}

	// Basic Payment örneği
	fmt.Println("\n=== Basic Payment ===")
	basicRequest := &iyzipay.BasicPaymentRequest{
		Locale:         iyzipay.LocaleTR,
		ConversationID: "123456789",
		Price:          "1.0",
		PaidPrice:      "1.2",
		Installment:    1,
		BuyerEmail:     "email@email.com",
		BuyerID:        "BY789",
		BuyerIP:        "85.34.78.112",
		PosOrderID:     "AP789",
		PaymentCard: &iyzipay.PaymentCard{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
			CVC:            "123",
		},
		Currency: iyzipay.CurrencyTRY,
	}

	basicResponse, err := client.BasicPayment.Create(ctx, basicRequest)
	if err != nil {
		log.Printf("Basic Payment Error: %v", err)
	} else {
		fmt.Printf("Basic Payment Status: %s\n", basicResponse.Status)
		if basicResponse.Status == "success" {
			fmt.Printf("Basic Payment ID: %s\n", basicResponse.PaymentID)
		} else {
			fmt.Printf("Basic Payment Failed - Error: %s\n", basicResponse.ErrorMessage)
		}
	}
}
