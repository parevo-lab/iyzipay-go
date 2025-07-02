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

	// 3D Secure Initialize
	fmt.Println("=== 3D Secure Initialize ===")

	threedsRequest := &iyzipay.PaymentRequest{
		Locale:         iyzipay.LocaleTR,
		ConversationID: "123456789",
		Price:          "1.0",
		PaidPrice:      "1.2",
		Currency:       iyzipay.CurrencyTRY,
		Installment:    1,
		BasketID:       "B67832",
		PaymentChannel: iyzipay.PaymentChannelWeb,
		PaymentGroup:   iyzipay.PaymentGroupProduct,
		CallbackURL:    "https://www.merchant.com/callback",
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
	initResponse, err := client.ThreedsInitialize.Create(ctx, threedsRequest)
	if err != nil {
		log.Printf("3DS Initialize Error: %v", err)
		return
	}

	fmt.Printf("3DS Initialize Status: %s\n", initResponse.Status)
	if initResponse.Status == "success" {
		fmt.Printf("Payment ID: %s\n", initResponse.PaymentID)
		fmt.Printf("MD Status: %d\n", initResponse.MdStatus)

		// İmza doğrulama
		if initResponse.Signature != "" {
			params := []string{
				initResponse.PaymentID,
				initResponse.ConversationID,
			}
			expectedSignature := iyzipay.CalculateHMACSignature(params, client.GetConfig().SecretKey)
			if initResponse.Signature == expectedSignature {
				fmt.Println("✓ 3DS Initialize Signature verified successfully")
			} else {
				fmt.Println("✗ 3DS Initialize Signature verification failed")
			}
		}

		// 3D Secure form HTML içeriği konsola yazdırılır
		// Gerçek uygulamada bu HTML kullanıcıya gösterilir
		if initResponse.ThreedsFormData != "" {
			fmt.Println("\n3D Secure Form HTML content received (length:", len(initResponse.ThreedsFormData), ")")
			fmt.Println("In real application, this HTML should be displayed to user for 3DS authentication")
		}

		// Simülasyon: Kullanıcı 3DS doğrulamasını tamamladı varsayalım
		// Gerçek uygulamada bu bilgiler callback URL'den gelir
		fmt.Println("\n=== 3D Secure Payment Completion ===")

		// 3DS Payment completion request
		paymentRequest := &iyzipay.ThreedsPaymentRequest{
			Locale:           iyzipay.LocaleTR,
			ConversationID:   "123456789",
			PaymentID:        initResponse.PaymentID,
			ConversationData: "conversation data", // Bu gerçek uygulamada callback'den gelir
		}

		paymentResponse, err := client.ThreedsPayment.Create(ctx, paymentRequest)
		if err != nil {
			log.Printf("3DS Payment Error: %v", err)
		} else {
			fmt.Printf("3DS Payment Status: %s\n", paymentResponse.Status)
			if paymentResponse.Status == "success" {
				fmt.Printf("Final Payment ID: %s\n", paymentResponse.PaymentID)
				fmt.Printf("Fraud Status: %d\n", paymentResponse.FraudStatus)
				fmt.Printf("Card Association: %s\n", paymentResponse.CardAssociation)
				fmt.Printf("Auth Code: %s\n", paymentResponse.AuthCode)

				// İmza doğrulama
				if paymentResponse.Signature != "" {
					params := []string{
						paymentResponse.PaymentID,
						paymentResponse.Currency,
						paymentResponse.BasketID,
						paymentResponse.ConversationID,
						paymentResponse.PaidPrice,
						paymentResponse.Price,
					}
					expectedSignature := iyzipay.CalculateHMACSignature(params, client.GetConfig().SecretKey)
					if paymentResponse.Signature == expectedSignature {
						fmt.Println("✓ 3DS Payment Signature verified successfully")
					} else {
						fmt.Println("✗ 3DS Payment Signature verification failed")
					}
				}
			} else {
				fmt.Printf("3DS Payment Failed - Error: %s\n", paymentResponse.ErrorMessage)
			}
		}
	} else {
		fmt.Printf("3DS Initialize Failed - Error Code: %s\n", initResponse.ErrorCode)
		fmt.Printf("Error Message: %s\n", initResponse.ErrorMessage)
	}

	// Basic 3DS Initialize örneği
	fmt.Println("\n=== Basic 3DS Initialize ===")
	basicThreedsRequest := &iyzipay.BasicPaymentRequest{
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
		Currency:    iyzipay.CurrencyTRY,
		CallbackURL: "https://www.merchant.com/callback",
	}

	basicInitResponse, err := client.ThreedsInitialize.CreateBasic(ctx, basicThreedsRequest)
	if err != nil {
		log.Printf("Basic 3DS Initialize Error: %v", err)
	} else {
		fmt.Printf("Basic 3DS Initialize Status: %s\n", basicInitResponse.Status)
		if basicInitResponse.Status == "success" {
			fmt.Printf("Basic 3DS Payment ID: %s\n", basicInitResponse.PaymentID)
			fmt.Println("Basic 3DS Form HTML content received")
		} else {
			fmt.Printf("Basic 3DS Initialize Failed - Error: %s\n", basicInitResponse.ErrorMessage)
		}
	}
}
