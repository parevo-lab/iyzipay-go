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

	// Kart oluşturma/saklama
	fmt.Println("=== Kart Saklama ===")

	createCardRequest := &iyzipay.CreateCardRequest{
		Locale:         iyzipay.LocaleTR,
		ConversationID: "123456789",
		Email:          "email@email.com",
		ExternalID:     "external id",
		Card: &iyzipay.CardInformation{
			CardAlias:      "My card",
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
		},
	}

	cardResponse, err := client.Card.Create(ctx, createCardRequest)
	if err != nil {
		log.Printf("Card Create Error: %v", err)
		return
	}

	fmt.Printf("Card Create Status: %s\n", cardResponse.Status)
	if cardResponse.Status == "success" {
		fmt.Printf("Card Token: %s\n", cardResponse.CardToken)
		fmt.Printf("Card User Key: %s\n", cardResponse.CardUserKey)
		fmt.Printf("External ID: %s\n", cardResponse.ExternalID)
		fmt.Printf("Email: %s\n", cardResponse.Email)
		fmt.Printf("BIN Number: %s\n", cardResponse.BinNumber)
		fmt.Printf("Card Type: %s\n", cardResponse.CardType)
		fmt.Printf("Card Association: %s\n", cardResponse.CardAssociation)
		fmt.Printf("Card Family: %s\n", cardResponse.CardFamily)
		fmt.Printf("Card Bank Code: %s\n", cardResponse.CardBankCode)
		fmt.Printf("Card Bank Name: %s\n", cardResponse.CardBankName)

		// Saklanan kartları listeleme
		fmt.Println("\n=== Kart Listesi ===")

		listRequest := &iyzipay.RetrieveCardListRequest{
			Locale:         iyzipay.LocaleTR,
			ConversationID: "123456789",
			CardUserKey:    cardResponse.CardUserKey,
		}

		listResponse, err := client.Card.List(ctx, listRequest)
		if err != nil {
			log.Printf("Card List Error: %v", err)
		} else {
			fmt.Printf("Card List Status: %s\n", listResponse.Status)
			if listResponse.Status == "success" {
				fmt.Printf("Total Cards: %d\n", len(listResponse.CardDetails))

				for i, card := range listResponse.CardDetails {
					fmt.Printf("\nCard %d:\n", i+1)
					fmt.Printf("  Card Token: %s\n", card.CardToken)
					fmt.Printf("  Card User Key: %s\n", card.CardUserKey)
					fmt.Printf("  Card Alias: %s\n", card.CardAlias)
					fmt.Printf("  BIN Number: %s\n", card.BinNumber)
					fmt.Printf("  Last Four Digits: %s\n", card.LastFourDigits)
					fmt.Printf("  Card Type: %s\n", card.CardType)
					fmt.Printf("  Card Association: %s\n", card.CardAssociation)
					fmt.Printf("  Card Family: %s\n", card.CardFamily)
					fmt.Printf("  Card Bank Code: %s\n", card.CardBankCode)
					fmt.Printf("  Card Bank Name: %s\n", card.CardBankName)
				}
			} else {
				fmt.Printf("Card List Failed - Error: %s\n", listResponse.ErrorMessage)
			}
		}

		// Saklanan kart ile ödeme örneği
		fmt.Println("\n=== Saklanan Kart ile Ödeme ===")

		savedCardPaymentRequest := &iyzipay.PaymentRequest{
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
				CardUserKey: cardResponse.CardUserKey,
				CardToken:   cardResponse.CardToken,
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

		savedCardPaymentResponse, err := client.Payment.Create(ctx, savedCardPaymentRequest)
		if err != nil {
			log.Printf("Saved Card Payment Error: %v", err)
		} else {
			fmt.Printf("Saved Card Payment Status: %s\n", savedCardPaymentResponse.Status)
			if savedCardPaymentResponse.Status == "success" {
				fmt.Printf("Payment ID: %s\n", savedCardPaymentResponse.PaymentID)
				fmt.Printf("Card Token: %s\n", savedCardPaymentResponse.CardToken)
				fmt.Printf("Card User Key: %s\n", savedCardPaymentResponse.CardUserKey)
				fmt.Println("✓ Payment completed successfully with saved card")
			} else {
				fmt.Printf("Saved Card Payment Failed - Error: %s\n", savedCardPaymentResponse.ErrorMessage)
			}
		}

		// Kartı silme
		fmt.Println("\n=== Kart Silme ===")

		deleteRequest := &iyzipay.DeleteCardRequest{
			Locale:         iyzipay.LocaleTR,
			ConversationID: "123456789",
			CardToken:      cardResponse.CardToken,
			CardUserKey:    cardResponse.CardUserKey,
		}

		deleteResponse, err := client.Card.Delete(ctx, deleteRequest)
		if err != nil {
			log.Printf("Card Delete Error: %v", err)
		} else {
			fmt.Printf("Card Delete Status: %s\n", deleteResponse.Status)
			if deleteResponse.Status == "success" {
				fmt.Println("✓ Card deleted successfully")
			} else {
				fmt.Printf("Card Delete Failed - Error: %s\n", deleteResponse.ErrorMessage)
			}
		}

	} else {
		fmt.Printf("Card Create Failed - Error Code: %s\n", cardResponse.ErrorCode)
		fmt.Printf("Error Message: %s\n", cardResponse.ErrorMessage)
		fmt.Printf("Error Group: %s\n", cardResponse.ErrorGroup)
	}

	// BIN Number kontrolü
	fmt.Println("\n=== BIN Number Kontrolü ===")

	binRequest := &iyzipay.RetrieveBinNumberRequest{
		Locale:         iyzipay.LocaleTR,
		ConversationID: "123456789",
		BinNumber:      "552879",
	}

	binResponse, err := client.BinNumber.Retrieve(ctx, binRequest)
	if err != nil {
		log.Printf("BIN Number Error: %v", err)
	} else {
		fmt.Printf("BIN Number Status: %s\n", binResponse.Status)
		if binResponse.Status == "success" {
			fmt.Printf("BIN Number: %s\n", binResponse.BinNumber)
			fmt.Printf("Card Type: %s\n", binResponse.CardType)
			fmt.Printf("Card Association: %s\n", binResponse.CardAssociation)
			fmt.Printf("Card Family: %s\n", binResponse.CardFamily)
			fmt.Printf("Bank Name: %s\n", binResponse.BankName)
			fmt.Printf("Bank Code: %s\n", binResponse.BankCode)
			fmt.Printf("Commercial: %v\n", binResponse.Commercial)
		} else {
			fmt.Printf("BIN Number Failed - Error: %s\n", binResponse.ErrorMessage)
		}
	}

	// Taksit bilgilerini alma
	fmt.Println("\n=== Taksit Bilgileri ===")

	installmentRequest := &iyzipay.RetrieveInstallmentInfoRequest{
		Locale:         iyzipay.LocaleTR,
		ConversationID: "123456789",
		BinNumber:      "552879",
		Price:          "100.0",
	}

	installmentResponse, err := client.InstallmentInfo.Retrieve(ctx, installmentRequest)
	if err != nil {
		log.Printf("Installment Info Error: %v", err)
	} else {
		fmt.Printf("Installment Info Status: %s\n", installmentResponse.Status)
		if installmentResponse.Status == "success" {
			fmt.Printf("Total Banks: %d\n", len(installmentResponse.InstallmentDetails))

			for i, detail := range installmentResponse.InstallmentDetails {
				fmt.Printf("\nBank %d:\n", i+1)
				fmt.Printf("  BIN Number: %s\n", detail.BinNumber)
				fmt.Printf("  Bank Name: %s\n", detail.BankName)
				fmt.Printf("  Bank Code: %s\n", detail.BankCode)
				fmt.Printf("  Card Type: %s\n", detail.CardType)
				fmt.Printf("  Card Association: %s\n", detail.CardAssociation)
				fmt.Printf("  Card Family: %s\n", detail.CardFamilyName)
				fmt.Printf("  Force 3DS: %v\n", detail.Force3DS)
				fmt.Printf("  Force CVC: %v\n", detail.ForceCvc)
				fmt.Printf("  Commercial: %v\n", detail.Commercial)

				if len(detail.InstallmentPrices) > 0 {
					fmt.Println("  Installment Options:")
					for _, price := range detail.InstallmentPrices {
						fmt.Printf("    %d Taksit: %s TL (Toplam: %s TL)\n",
							price.InstallmentNumber,
							price.InstallmentPrice,
							price.TotalPrice)
					}
				}
			}
		} else {
			fmt.Printf("Installment Info Failed - Error: %s\n", installmentResponse.ErrorMessage)
		}
	}
}
