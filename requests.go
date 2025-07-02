package iyzipay

// Request types for various İyzipay API endpoints

// RetrievePaymentRequest represents retrieve payment request
type RetrievePaymentRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	PaymentID      string `json:"paymentId"`
	IP             string `json:"ip"`
}

// ThreedsInitializeResponse represents 3DS initialize response
type ThreedsInitializeResponse struct {
	BaseResponse
	PaymentID               string `json:"paymentId"`
	ThreedsFormData         string `json:"threeDSHtmlContent"`
	RedirectURL             string `json:"redirectUrl"`
	PaymentStatus           string `json:"paymentStatus"`
	MdStatus                int    `json:"mdStatus"`
	Signature               string `json:"signature"`
}

// ThreedsPaymentRequest represents 3DS payment completion request
type ThreedsPaymentRequest struct {
	Locale           string `json:"locale"`
	ConversationID   string `json:"conversationId"`
	PaymentID        string `json:"paymentId"`
	ConversationData string `json:"conversationData"`
}

// CheckoutFormInitializeRequest represents checkout form initialize request
type CheckoutFormInitializeRequest struct {
	Locale             string        `json:"locale"`
	ConversationID     string        `json:"conversationId"`
	Price              string        `json:"price"`
	PaidPrice          string        `json:"paidPrice"`
	Currency           string        `json:"currency"`
	BasketID           string        `json:"basketId"`
	PaymentGroup       string        `json:"paymentGroup"`
	CallbackURL        string        `json:"callbackUrl"`
	EnabledInstallments []int        `json:"enabledInstallments"`
	Buyer              *Buyer        `json:"buyer"`
	ShippingAddress    *Address      `json:"shippingAddress"`
	BillingAddress     *Address      `json:"billingAddress"`
	BasketItems        []BasketItem  `json:"basketItems"`
	PaymentSource      string        `json:"paymentSource"`
	ForceThreeDS       *bool         `json:"forceThreeDS,omitempty"`
	CardUserKey        string        `json:"cardUserKey"`
	PosOrderID         string        `json:"posOrderId"`
}

// CheckoutFormInitializeResponse represents checkout form initialize response
type CheckoutFormInitializeResponse struct {
	BaseResponse
	Token           string `json:"token"`
	CheckoutFormURL string `json:"checkoutFormContent"`
	TokenExpireTime int64  `json:"tokenExpireTime"`
	PaymentPageURL  string `json:"paymentPageUrl"`
	Signature       string `json:"signature"`
}

// RetrieveCheckoutFormRequest represents retrieve checkout form request
type RetrieveCheckoutFormRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	Token          string `json:"token"`
}

// CheckoutFormResponse represents checkout form response
type CheckoutFormResponse struct {
	BaseResponse
	Token               string            `json:"token"`
	CallbackURL         string            `json:"callbackUrl"`
	PaymentStatus       string            `json:"paymentStatus"`
	PaymentID           string            `json:"paymentId"`
	Price               string            `json:"price"`
	PaidPrice           string            `json:"paidPrice"`
	Installment         int               `json:"installment"`
	Currency            string            `json:"currency"`
	BasketID            string            `json:"basketId"`
	ItemTransactions    []ItemTransaction `json:"itemTransactions"`
	MdStatus            int               `json:"mdStatus"`
	Signature           string            `json:"signature"`
}

// CreateCardRequest represents create card request
type CreateCardRequest struct {
	Locale         string           `json:"locale"`
	ConversationID string           `json:"conversationId"`
	Email          string           `json:"email"`
	ExternalID     string           `json:"externalId"`
	Card           *CardInformation `json:"card"`
}

// CardResponse represents card response
type CardResponse struct {
	BaseResponse
	ExternalID  string `json:"externalId"`
	Email       string `json:"email"`
	CardToken   string `json:"cardToken"`
	CardUserKey string `json:"cardUserKey"`
	BinNumber   string `json:"binNumber"`
	CardType    string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily  string `json:"cardFamily"`
	CardBankCode string `json:"cardBankCode"`
	CardBankName string `json:"cardBankName"`
}

// DeleteCardRequest represents delete card request
type DeleteCardRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	CardToken      string `json:"cardToken"`
	CardUserKey    string `json:"cardUserKey"`
}

// RetrieveCardListRequest represents retrieve card list request
type RetrieveCardListRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	CardUserKey    string `json:"cardUserKey"`
}

// CardListResponse represents card list response
type CardListResponse struct {
	BaseResponse
	CardDetails []CardDetail `json:"cardDetails"`
}

// CardDetail represents card detail in list
type CardDetail struct {
	CardToken       string `json:"cardToken"`
	CardUserKey     string `json:"cardUserKey"`
	CardAlias       string `json:"cardAlias"`
	BinNumber       string `json:"binNumber"`
	LastFourDigits  string `json:"lastFourDigits"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily      string `json:"cardFamily"`
	CardBankCode    string `json:"cardBankCode"`
	CardBankName    string `json:"cardBankName"`
}

// RefundResponse represents refund response
type RefundResponse struct {
	BaseResponse
	PaymentID            string `json:"paymentId"`
	PaymentTransactionID string `json:"paymentTransactionId"`
	Price                string `json:"price"`
	Currency             string `json:"currency"`
	ConnectorName        string `json:"connectorName"`
	AuthCode             string `json:"authCode"`
	HostReference        string `json:"hostReference"`
}

// CancelResponse represents cancel response
type CancelResponse struct {
	BaseResponse
	PaymentID     string `json:"paymentId"`
	Price         string `json:"price"`
	Currency      string `json:"currency"`
	ConnectorName string `json:"connectorName"`
	AuthCode      string `json:"authCode"`
	HostReference string `json:"hostReference"`
}

// CreateSubMerchantRequest represents create sub merchant request
type CreateSubMerchantRequest struct {
	Locale                string `json:"locale"`
	ConversationID        string `json:"conversationId"`
	SubMerchantExternalID string `json:"subMerchantExternalId"`
	SubMerchantType       string `json:"subMerchantType"`
	Address               string `json:"address"`
	ContactName           string `json:"contactName"`
	ContactSurname        string `json:"contactSurname"`
	Email                 string `json:"email"`
	GsmNumber             string `json:"gsmNumber"`
	Name                  string `json:"name"`
	IBAN                  string `json:"iban"`
	IdentityNumber        string `json:"identityNumber"`
	Currency              string `json:"currency"`
	TaxOffice             string `json:"taxOffice"`
	TaxNumber             string `json:"taxNumber"`
	LegalCompanyTitle     string `json:"legalCompanyTitle"`
	SwiftCode             string `json:"swiftCode"`
}

// UpdateSubMerchantRequest represents update sub merchant request
type UpdateSubMerchantRequest struct {
	Locale                string `json:"locale"`
	ConversationID        string `json:"conversationId"`
	SubMerchantKey        string `json:"subMerchantKey"`
	IBAN                  string `json:"iban"`
	Address               string `json:"address"`
	ContactName           string `json:"contactName"`
	ContactSurname        string `json:"contactSurname"`
	Email                 string `json:"email"`
	GsmNumber             string `json:"gsmNumber"`
	Name                  string `json:"name"`
	IdentityNumber        string `json:"identityNumber"`
	Currency              string `json:"currency"`
	TaxOffice             string `json:"taxOffice"`
	TaxNumber             string `json:"taxNumber"`
	LegalCompanyTitle     string `json:"legalCompanyTitle"`
	SwiftCode             string `json:"swiftCode"`
}

// RetrieveSubMerchantRequest represents retrieve sub merchant request
type RetrieveSubMerchantRequest struct {
	Locale                string `json:"locale"`
	ConversationID        string `json:"conversationId"`
	SubMerchantExternalID string `json:"subMerchantExternalId"`
}

// SubMerchantResponse represents sub merchant response
type SubMerchantResponse struct {
	BaseResponse
	SubMerchantKey        string `json:"subMerchantKey"`
	SubMerchantExternalID string `json:"subMerchantExternalId"`
	SubMerchantType       string `json:"subMerchantType"`
	Address               string `json:"address"`
	ContactName           string `json:"contactName"`
	ContactSurname        string `json:"contactSurname"`
	Email                 string `json:"email"`
	GsmNumber             string `json:"gsmNumber"`
	Name                  string `json:"name"`
	IBAN                  string `json:"iban"`
	IdentityNumber        string `json:"identityNumber"`
	Currency              string `json:"currency"`
}

// BKMInitializeRequest represents BKM initialize request
type BKMInitializeRequest struct {
	Locale             string        `json:"locale"`
	ConversationID     string        `json:"conversationId"`
	Price              string        `json:"price"`
	BasketID           string        `json:"basketId"`
	PaymentGroup       string        `json:"paymentGroup"`
	Buyer              *Buyer        `json:"buyer"`
	ShippingAddress    *Address      `json:"shippingAddress"`
	BillingAddress     *Address      `json:"billingAddress"`
	BasketItems        []BasketItem  `json:"basketItems"`
	CallbackURL        string        `json:"callbackUrl"`
	PaymentSource      string        `json:"paymentSource"`
	EnabledInstallments []BkmInstallment `json:"enabledInstallments"`
}

// BasicBKMInitializeRequest represents basic BKM initialize request
type BasicBKMInitializeRequest struct {
	Locale             string           `json:"locale"`
	ConversationID     string           `json:"conversationId"`
	Price              string           `json:"price"`
	CallbackURL        string           `json:"callbackUrl"`
	BuyerEmail         string           `json:"buyerEmail"`
	BuyerID            string           `json:"buyerId"`
	BuyerIP            string           `json:"buyerIp"`
	PosOrderID         string           `json:"posOrderId"`
	EnabledInstallments []BkmInstallment `json:"enabledInstallments"`
}

// BKMInitializeResponse represents BKM initialize response
type BKMInitializeResponse struct {
	BaseResponse
	HtmlContent string `json:"htmlContent"`
	Token       string `json:"token"`
	Signature   string `json:"signature"`
}

// RetrieveBKMRequest represents retrieve BKM request
type RetrieveBKMRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	Token          string `json:"token"`
}

// BKMResponse represents BKM response
type BKMResponse struct {
	BaseResponse
	Token               string            `json:"token"`
	PaymentStatus       string            `json:"paymentStatus"`
	PaymentID           string            `json:"paymentId"`
	Price               string            `json:"price"`
	PaidPrice           string            `json:"paidPrice"`
	Currency            string            `json:"currency"`
	BasketID            string            `json:"basketId"`
	Installment         int               `json:"installment"`
	ItemTransactions    []ItemTransaction `json:"itemTransactions"`
	Signature           string            `json:"signature"`
}

// RetrieveAPMRequest represents retrieve APM request
type RetrieveAPMRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	PaymentID      string `json:"paymentId"`
}

// APMInitializeResponse represents APM initialize response
type APMInitializeResponse struct {
	BaseResponse
	PaymentID   string `json:"paymentId"`
	RedirectURL string `json:"redirectUrl"`
	HtmlContent string `json:"htmlContent"`
	Signature   string `json:"signature"`
}

// APMResponse represents APM response
type APMResponse struct {
	BaseResponse
	PaymentID           string            `json:"paymentId"`
	PaymentStatus       string            `json:"paymentStatus"`
	Price               string            `json:"price"`
	PaidPrice           string            `json:"paidPrice"`
	Currency            string            `json:"currency"`
	MerchantOrderID     string            `json:"merchantOrderId"`
	BasketID            string            `json:"basketId"`
	ItemTransactions    []ItemTransaction `json:"itemTransactions"`
	Signature           string            `json:"signature"`
}

// CreateSubscriptionInitRequest represents subscription init request
type CreateSubscriptionInitRequest struct {
	Locale                    string               `json:"locale"`
	ConversationID            string               `json:"conversationId"`
	PricingPlanReferenceCode  string               `json:"pricingPlanReferenceCode"`
	SubscriptionInitialStatus string               `json:"subscriptionInitialStatus"`
	Customer                  *SubscriptionCustomer `json:"customer"`
	PaymentCard               *SubscriptionCard     `json:"paymentCard"`
}

// SubscriptionInitializeResponse represents subscription initialize response
type SubscriptionInitializeResponse struct {
	BaseResponse
	SubscriptionReferenceCode string `json:"subscriptionReferenceCode"`
	ParentReferenceCode       string `json:"parentReferenceCode"`
	PricingPlanReferenceCode  string `json:"pricingPlanReferenceCode"`
	SubscriptionStatus        string `json:"subscriptionStatus"`
	TrialDays                 int    `json:"trialDays"`
	TrialStartDate            string `json:"trialStartDate"`
	TrialEndDate              string `json:"trialEndDate"`
	StartDate                 string `json:"startDate"`
	EndDate                   string `json:"endDate"`
}

// RetrieveInstallmentInfoRequest represents retrieve installment info request
type RetrieveInstallmentInfoRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	BinNumber      string `json:"binNumber"`
	Price          string `json:"price"`
}

// InstallmentInfoResponse represents installment info response
type InstallmentInfoResponse struct {
	BaseResponse
	InstallmentDetails []InstallmentDetail `json:"installmentDetails"`
}

// InstallmentDetail represents installment detail
type InstallmentDetail struct {
	BinNumber         string              `json:"binNumber"`
	Price             string              `json:"price"`
	CardType          string              `json:"cardType"`
	CardAssociation   string              `json:"cardAssociation"`
	CardFamilyName    string              `json:"cardFamilyName"`
	Force3DS          bool                `json:"force3ds"`
	BankCode          string              `json:"bankCode"`
	BankName          string              `json:"bankName"`
	ForceCvc          bool                `json:"forceCvc"`
	Commercial        bool                `json:"commercial"`
	InstallmentPrices []InstallmentPrice  `json:"installmentPrices"`
}

// InstallmentPrice represents installment price
type InstallmentPrice struct {
	InstallmentNumber       int    `json:"installmentNumber"`
	Price                   string `json:"price"`
	TotalPrice              string `json:"totalPrice"`
	InstallmentPrice        string `json:"installmentPrice"`
}

// RetrieveBinNumberRequest represents retrieve BIN number request
type RetrieveBinNumberRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	BinNumber      string `json:"binNumber"`
}

// BinNumberResponse represents BIN number response
type BinNumberResponse struct {
	BaseResponse
	BinNumber       string `json:"binNumber"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily      string `json:"cardFamily"`
	BankName        string `json:"bankName"`
	BankCode        string `json:"bankCode"`
	Commercial      bool   `json:"commercial"`
}

// UpdatePaymentItemRequest represents update payment item request
type UpdatePaymentItemRequest struct {
	Locale               string        `json:"locale"`
	ConversationID       string        `json:"conversationId"`
	SubMerchantKey       string        `json:"subMerchantKey"`
	PaymentTransactionID string        `json:"paymentTransactionId"`
	SubMerchantPrice     string        `json:"subMerchantPrice"`
	PaymentItems         []PaymentItem `json:"paymentItems"`
}

// PaymentItemResponse represents payment item response
type PaymentItemResponse struct {
	BaseResponse
	PaymentItems []PaymentItem `json:"paymentItems"`
}

// CrossBookingRequest represents cross booking request
type CrossBookingRequest struct {
	Locale               string `json:"locale"`
	ConversationID       string `json:"conversationId"`
	SubMerchantKey       string `json:"subMerchantKey"`
	Price                string `json:"price"`
	Reason               string `json:"reason"`
	Currency             string `json:"currency"`
}

// CrossBookingResponse represents cross booking response
type CrossBookingResponse struct {
	BaseResponse
	SubMerchantKey string `json:"subMerchantKey"`
	Price          string `json:"price"`
	Currency       string `json:"currency"`
}

// RefundToBalanceRequest represents refund to balance request
type RefundToBalanceRequest struct {
	Locale               string `json:"locale"`
	ConversationID       string `json:"conversationId"`
	PaymentTransactionID string `json:"paymentTransactionId"`
	Price                string `json:"price"`
	Currency             string `json:"currency"`
	CallbackURL          string `json:"callbackUrl"`
}

// RefundToBalanceResponse represents refund to balance response
type RefundToBalanceResponse struct {
	BaseResponse
	PaymentID            string `json:"paymentId"`
	PaymentTransactionID string `json:"paymentTransactionId"`
	Price                string `json:"price"`
	Currency             string `json:"currency"`
}

// SettlementToBalanceRequest represents settlement to balance request
type SettlementToBalanceRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	SubMerchantKey string `json:"subMerchantKey"`
	Price          string `json:"price"`
	Currency       string `json:"currency"`
	CallbackURL    string `json:"callbackUrl"`
}

// SettlementToBalanceResponse represents settlement to balance response
type SettlementToBalanceResponse struct {
	BaseResponse
	SubMerchantKey string `json:"subMerchantKey"`
	Price          string `json:"price"`
	Currency       string `json:"currency"`
}

// UniversalCardStorageInitializeRequest represents universal card storage initialize request
type UniversalCardStorageInitializeRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	Email          string `json:"email"`
	GsmNumber      string `json:"gsmNumber"`
	CardAlias      string `json:"cardAlias"`
	CallbackURL    string `json:"callbackUrl"`
}

// UniversalCardStorageInitializeResponse represents universal card storage initialize response
type UniversalCardStorageInitializeResponse struct {
	BaseResponse
	UcsToken    string `json:"ucsToken"`
	UcsURL      string `json:"ucsUrl"`
	UcsContent  string `json:"ucsContent"`
}