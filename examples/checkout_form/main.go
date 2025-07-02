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

	// Checkout Form Initialize
	fmt.Println("=== Checkout Form Initialize ===")

	checkoutRequest := &iyzipay.CheckoutFormInitializeRequest{
		Locale:              iyzipay.LocaleTR,
		ConversationID:      "123456789",
		Price:               "1.0",
		PaidPrice:           "1.2",
		Currency:            iyzipay.CurrencyTRY,
		BasketID:            "B67832",
		PaymentGroup:        iyzipay.PaymentGroupProduct,
		CallbackURL:         "https://www.merchant.com/callback",
		EnabledInstallments: []int{2, 3, 6, 9},
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
		PaymentSource: "API",
		PosOrderID:    "POS789",
		CardUserKey:   "", // Stored card user key (optional)
	}

	ctx := context.Background()
	initResponse, err := client.CheckoutForm.Initialize(ctx, checkoutRequest)
	if err != nil {
		log.Printf("Checkout Form Initialize Error: %v", err)
		return
	}

	fmt.Printf("Checkout Form Initialize Status: %s\n", initResponse.Status)
	if initResponse.Status == "success" {
		fmt.Printf("Token: %s\n", initResponse.Token)
		fmt.Printf("Token Expire Time: %d\n", initResponse.TokenExpireTime)

		// İmza doğrulama
		if initResponse.Signature != "" {
			params := []string{
				initResponse.ConversationID,
				initResponse.Token,
			}
			expectedSignature := iyzipay.CalculateHMACSignature(params, client.GetConfig().SecretKey)
			if initResponse.Signature == expectedSignature {
				fmt.Println("✓ Checkout Form Initialize Signature verified successfully")
			} else {
				fmt.Println("✗ Checkout Form Initialize Signature verification failed")
			}
		}

		// Checkout form HTML içeriği konsola yazdırılır
		if initResponse.CheckoutFormURL != "" {
			fmt.Println("\nCheckout Form HTML content received (length:", len(initResponse.CheckoutFormURL), ")")
			fmt.Println("In real application, this HTML should be displayed to user")
		}

		// Payment page URL
		if initResponse.PaymentPageURL != "" {
			fmt.Printf("Payment Page URL: %s\n", initResponse.PaymentPageURL)
			fmt.Println("User should be redirected to this URL for payment")
		}

		// Simülasyon: Kullanıcı ödemeyi tamamladı varsayalım
		// Gerçek uygulamada bu token callback URL'den gelir
		fmt.Println("\n=== Checkout Form Result Retrieval ===")

		// Checkout form result retrieval
		retrieveRequest := &iyzipay.RetrieveCheckoutFormRequest{
			Locale:         iyzipay.LocaleTR,
			ConversationID: "123456789",
			Token:          initResponse.Token,
		}

		resultResponse, err := client.CheckoutForm.Retrieve(ctx, retrieveRequest)
		if err != nil {
			log.Printf("Checkout Form Retrieve Error: %v", err)
		} else {
			fmt.Printf("Checkout Form Result Status: %s\n", resultResponse.Status)
			if resultResponse.Status == "success" {
				fmt.Printf("Payment Status: %s\n", resultResponse.PaymentStatus)
				fmt.Printf("Payment ID: %s\n", resultResponse.PaymentID)
				fmt.Printf("Price: %s\n", resultResponse.Price)
				fmt.Printf("Paid Price: %s\n", resultResponse.PaidPrice)
				fmt.Printf("Currency: %s\n", resultResponse.Currency)
				fmt.Printf("Installment: %d\n", resultResponse.Installment)
				fmt.Printf("Basket ID: %s\n", resultResponse.BasketID)
				fmt.Printf("MD Status: %d\n", resultResponse.MdStatus)

				// İmza doğrulama
				if resultResponse.Signature != "" {
					params := []string{
						resultResponse.PaymentID,
						resultResponse.Currency,
						resultResponse.BasketID,
						resultResponse.ConversationID,
						resultResponse.PaidPrice,
						resultResponse.Price,
					}
					expectedSignature := iyzipay.CalculateHMACSignature(params, client.GetConfig().SecretKey)
					if resultResponse.Signature == expectedSignature {
						fmt.Println("✓ Checkout Form Result Signature verified successfully")
					} else {
						fmt.Println("✗ Checkout Form Result Signature verification failed")
					}
				}

				// Item transactions detayları
				if len(resultResponse.ItemTransactions) > 0 {
					fmt.Println("\nItem Transactions:")
					for i, item := range resultResponse.ItemTransactions {
						fmt.Printf("  Item %d:\n", i+1)
						fmt.Printf("    Item ID: %s\n", item.ItemID)
						fmt.Printf("    Payment Transaction ID: %s\n", item.PaymentTransactionID)
						fmt.Printf("    Transaction Status: %d\n", item.TransactionStatus)
						fmt.Printf("    Price: %s\n", item.Price)
						fmt.Printf("    Paid Price: %s\n", item.PaidPrice)
					}
				}
			} else {
				fmt.Printf("Checkout Form Failed - Error: %s\n", resultResponse.ErrorMessage)
			}
		}

	} else {
		fmt.Printf("Checkout Form Initialize Failed - Error Code: %s\n", initResponse.ErrorCode)
		fmt.Printf("Error Message: %s\n", initResponse.ErrorMessage)
		fmt.Printf("Error Group: %s\n", initResponse.ErrorGroup)
	}

	// Force 3DS örneği
	fmt.Println("\n=== Checkout Form with Force 3DS ===")

	forceThreeDS := true
	checkoutRequestWith3DS := &iyzipay.CheckoutFormInitializeRequest{
		Locale:              iyzipay.LocaleTR,
		ConversationID:      "123456790",
		Price:               "10.0",
		PaidPrice:           "12.0",
		Currency:            iyzipay.CurrencyTRY,
		BasketID:            "B67833",
		PaymentGroup:        iyzipay.PaymentGroupProduct,
		CallbackURL:         "https://www.merchant.com/callback",
		EnabledInstallments: []int{1},
		ForceThreeDS:        &forceThreeDS,
		Buyer: &iyzipay.Buyer{
			ID:                  "BY790",
			Name:                "Jane",
			Surname:             "Smith",
			GsmNumber:           "+905350000001",
			Email:               "jane@email.com",
			IdentityNumber:      "11111111111",
			LastLoginDate:       "2015-10-05 12:43:35",
			RegistrationDate:    "2013-04-21 15:12:09",
			RegistrationAddress: "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			IP:                  "85.34.78.112",
			City:                "Istanbul",
			Country:             "Turkey",
			ZipCode:             "34732",
		},
		ShippingAddress: &iyzipay.Address{
			ContactName: "Jane Smith",
			City:        "Istanbul",
			Country:     "Turkey",
			Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			ZipCode:     "34742",
		},
		BillingAddress: &iyzipay.Address{
			ContactName: "Jane Smith",
			City:        "Istanbul",
			Country:     "Turkey",
			Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			ZipCode:     "34742",
		},
		BasketItems: []iyzipay.BasketItem{
			{
				ID:        "BI201",
				Name:      "Premium Product",
				Category1: "Electronics",
				Category2: "Phone",
				ItemType:  iyzipay.BasketItemTypePhysical,
				Price:     "10.0",
			},
		},
		PaymentSource: "API",
		PosOrderID:    "POS790",
	}

	force3DSResponse, err := client.CheckoutForm.Initialize(ctx, checkoutRequestWith3DS)
	if err != nil {
		log.Printf("Force 3DS Checkout Form Error: %v", err)
	} else {
		fmt.Printf("Force 3DS Checkout Form Status: %s\n", force3DSResponse.Status)
		if force3DSResponse.Status == "success" {
			fmt.Printf("Force 3DS Token: %s\n", force3DSResponse.Token)
			fmt.Println("This payment will require 3D Secure authentication")
		} else {
			fmt.Printf("Force 3DS Failed - Error: %s\n", force3DSResponse.ErrorMessage)
		}
	}
}
