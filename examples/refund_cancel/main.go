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

	// Önce bir ödeme yapalım ki iade/iptal edebilsin
	fmt.Println("=== Ödeme Oluşturma (İade/İptal için) ===")

	paymentRequest := &iyzipay.PaymentRequest{
		Locale:         iyzipay.LocaleTR,
		ConversationID: "123456789",
		Price:          "2.0",
		PaidPrice:      "2.4",
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
				Price:     "1.0",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  iyzipay.BasketItemTypeVirtual,
				Price:     "1.0",
			},
		},
	}

	paymentResponse, err := client.Payment.Create(ctx, paymentRequest)
	if err != nil {
		log.Printf("Payment Error: %v", err)
		return
	}

	fmt.Printf("Payment Status: %s\n", paymentResponse.Status)
	if paymentResponse.Status != "success" {
		fmt.Printf("Payment Failed - Error: %s\n", paymentResponse.ErrorMessage)
		return
	}

	fmt.Printf("Payment ID: %s\n", paymentResponse.PaymentID)

	// İade işlemi (Kısmi İade)
	fmt.Println("\n=== Kısmi İade İşlemi ===")

	// İlk item transaction ID'sini al
	if len(paymentResponse.ItemTransactions) == 0 {
		fmt.Println("No item transactions found for refund")
		return
	}

	firstItemTransactionID := paymentResponse.ItemTransactions[0].PaymentTransactionID

	refundRequest := &iyzipay.RefundRequest{
		Locale:               iyzipay.LocaleTR,
		ConversationID:       "123456789",
		PaymentTransactionID: firstItemTransactionID,
		Price:                "0.5", // Kısmi iade
		Currency:             iyzipay.CurrencyTRY,
		IP:                   "85.34.78.112",
		Reason:               iyzipay.RefundReasonBuyerRequest,
		Description:          "Customer requested refund for item",
	}

	refundResponse, err := client.Refund.Create(ctx, refundRequest)
	if err != nil {
		log.Printf("Refund Error: %v", err)
	} else {
		fmt.Printf("Refund Status: %s\n", refundResponse.Status)
		if refundResponse.Status == "success" {
			fmt.Printf("Refund Payment ID: %s\n", refundResponse.PaymentID)
			fmt.Printf("Refund Transaction ID: %s\n", refundResponse.PaymentTransactionID)
			fmt.Printf("Refund Price: %s\n", refundResponse.Price)
			fmt.Printf("Refund Currency: %s\n", refundResponse.Currency)
			fmt.Printf("Connector Name: %s\n", refundResponse.ConnectorName)
			fmt.Printf("Auth Code: %s\n", refundResponse.AuthCode)
			fmt.Printf("Host Reference: %s\n", refundResponse.HostReference)
			fmt.Println("✓ Partial refund completed successfully")
		} else {
			fmt.Printf("Refund Failed - Error Code: %s\n", refundResponse.ErrorCode)
			fmt.Printf("Error Message: %s\n", refundResponse.ErrorMessage)
		}
	}

	// Tam İade işlemi (Diğer item için)
	fmt.Println("\n=== Tam İade İşlemi ===")

	if len(paymentResponse.ItemTransactions) > 1 {
		secondItemTransactionID := paymentResponse.ItemTransactions[1].PaymentTransactionID

		fullRefundRequest := &iyzipay.RefundRequest{
			Locale:               iyzipay.LocaleTR,
			ConversationID:       "123456790",
			PaymentTransactionID: secondItemTransactionID,
			Price:                "1.0", // Tam iade
			Currency:             iyzipay.CurrencyTRY,
			IP:                   "85.34.78.112",
			Reason:               iyzipay.RefundReasonOther,
			Description:          "Full refund requested by customer",
		}

		fullRefundResponse, err := client.Refund.Create(ctx, fullRefundRequest)
		if err != nil {
			log.Printf("Full Refund Error: %v", err)
		} else {
			fmt.Printf("Full Refund Status: %s\n", fullRefundResponse.Status)
			if fullRefundResponse.Status == "success" {
				fmt.Printf("Full Refund Transaction ID: %s\n", fullRefundResponse.PaymentTransactionID)
				fmt.Printf("Full Refund Price: %s\n", fullRefundResponse.Price)
				fmt.Println("✓ Full refund completed successfully")
			} else {
				fmt.Printf("Full Refund Failed - Error: %s\n", fullRefundResponse.ErrorMessage)
			}
		}
	}

	// Yeni bir ödeme yap (iptal için)
	fmt.Println("\n=== İptal için Yeni Ödeme ===")

	cancelPaymentRequest := &iyzipay.PaymentRequest{
		Locale:         iyzipay.LocaleTR,
		ConversationID: "123456791",
		Price:          "1.0",
		PaidPrice:      "1.2",
		Currency:       iyzipay.CurrencyTRY,
		Installment:    1,
		BasketID:       "B67833",
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
			ID:                  "BY790",
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
				ID:        "BI201",
				Name:      "Test Product",
				Category1: "Test",
				Category2: "Product",
				ItemType:  iyzipay.BasketItemTypeVirtual,
				Price:     "1.0",
			},
		},
	}

	cancelPaymentResponse, err := client.Payment.Create(ctx, cancelPaymentRequest)
	if err != nil {
		log.Printf("Cancel Payment Error: %v", err)
		return
	}

	fmt.Printf("Cancel Payment Status: %s\n", cancelPaymentResponse.Status)
	if cancelPaymentResponse.Status != "success" {
		fmt.Printf("Cancel Payment Failed - Error: %s\n", cancelPaymentResponse.ErrorMessage)
		return
	}

	fmt.Printf("Cancel Payment ID: %s\n", cancelPaymentResponse.PaymentID)

	// İptal işlemi
	fmt.Println("\n=== İptal İşlemi ===")

	cancelRequest := &iyzipay.CancelRequest{
		Locale:         iyzipay.LocaleTR,
		ConversationID: "123456791",
		PaymentID:      cancelPaymentResponse.PaymentID,
		IP:             "85.34.78.112",
		Reason:         iyzipay.RefundReasonBuyerRequest,
		Description:    "Customer requested cancellation",
	}

	cancelResponse, err := client.Cancel.Create(ctx, cancelRequest)
	if err != nil {
		log.Printf("Cancel Error: %v", err)
	} else {
		fmt.Printf("Cancel Status: %s\n", cancelResponse.Status)
		if cancelResponse.Status == "success" {
			fmt.Printf("Cancelled Payment ID: %s\n", cancelResponse.PaymentID)
			fmt.Printf("Cancel Price: %s\n", cancelResponse.Price)
			fmt.Printf("Cancel Currency: %s\n", cancelResponse.Currency)
			fmt.Printf("Connector Name: %s\n", cancelResponse.ConnectorName)
			fmt.Printf("Auth Code: %s\n", cancelResponse.AuthCode)
			fmt.Printf("Host Reference: %s\n", cancelResponse.HostReference)
			fmt.Println("✓ Payment cancelled successfully")
		} else {
			fmt.Printf("Cancel Failed - Error Code: %s\n", cancelResponse.ErrorCode)
			fmt.Printf("Error Message: %s\n", cancelResponse.ErrorMessage)
		}
	}

	// Ödeme detaylarını alma
	fmt.Println("\n=== Ödeme Detaylarını Alma ===")

	retrieveRequest := &iyzipay.RetrievePaymentRequest{
		Locale:         iyzipay.LocaleTR,
		ConversationID: "123456789",
		PaymentID:      paymentResponse.PaymentID,
		IP:             "85.34.78.112",
	}

	retrieveResponse, err := client.Payment.Retrieve(ctx, retrieveRequest)
	if err != nil {
		log.Printf("Retrieve Payment Error: %v", err)
	} else {
		fmt.Printf("Retrieve Payment Status: %s\n", retrieveResponse.Status)
		if retrieveResponse.Status == "success" {
			fmt.Printf("Retrieved Payment ID: %s\n", retrieveResponse.PaymentID)
			fmt.Printf("Payment Status: %s\n", retrieveResponse.Phase)
			fmt.Printf("Fraud Status: %d\n", retrieveResponse.FraudStatus)
			fmt.Printf("Total Price: %s\n", retrieveResponse.Price)
			fmt.Printf("Paid Price: %s\n", retrieveResponse.PaidPrice)
			fmt.Printf("Currency: %s\n", retrieveResponse.Currency)
			fmt.Printf("Installment: %d\n", retrieveResponse.Installment)
			fmt.Printf("Card Type: %s\n", retrieveResponse.CardType)
			fmt.Printf("Card Association: %s\n", retrieveResponse.CardAssociation)
			fmt.Printf("Last Four Digits: %s\n", retrieveResponse.LastFourDigits)

			if len(retrieveResponse.ItemTransactions) > 0 {
				fmt.Println("\nItem Transaction Details:")
				for i, item := range retrieveResponse.ItemTransactions {
					fmt.Printf("  Item %d:\n", i+1)
					fmt.Printf("    Item ID: %s\n", item.ItemID)
					fmt.Printf("    Transaction ID: %s\n", item.PaymentTransactionID)
					fmt.Printf("    Transaction Status: %d\n", item.TransactionStatus)
					fmt.Printf("    Price: %s\n", item.Price)
					fmt.Printf("    Paid Price: %s\n", item.PaidPrice)
					fmt.Printf("    Merchant Commission: %s\n", item.MerchantCommissionRateAmount)
					fmt.Printf("    İyzico Commission: %s\n", item.IyziCommissionRateAmount)
				}
			}
		} else {
			fmt.Printf("Retrieve Payment Failed - Error: %s\n", retrieveResponse.ErrorMessage)
		}
	}
}
