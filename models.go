package iyzipay

// Address represents address information
type Address struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

// SubscriptionAddress extends Address with district field
type SubscriptionAddress struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
	District    string `json:"district"`
}

// Buyer represents buyer information
type Buyer struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	IdentityNumber      string `json:"identityNumber"`
	Email               string `json:"email"`
	GsmNumber           string `json:"gsmNumber"`
	RegistrationDate    string `json:"registrationDate"`
	LastLoginDate       string `json:"lastLoginDate"`
	RegistrationAddress string `json:"registrationAddress"`
	City                string `json:"city"`
	Country             string `json:"country"`
	ZipCode             string `json:"zipCode"`
	IP                  string `json:"ip"`
}

// PaymentCard represents payment card information
type PaymentCard struct {
	CardHolderName       string `json:"cardHolderName"`
	CardNumber           string `json:"cardNumber"`
	ExpireYear           string `json:"expireYear"`
	ExpireMonth          string `json:"expireMonth"`
	CVC                  string `json:"cvc"`
	RegisterCard         *bool  `json:"registerCard,omitempty"`
	CardAlias            string `json:"cardAlias"`
	CardToken            string `json:"cardToken"`
	CardUserKey          string `json:"cardUserKey"`
	ConsumerToken        string `json:"consumerToken"`
	RegisterConsumerCard *bool  `json:"registerConsumerCard,omitempty"`
	UcsToken             string `json:"ucsToken"`
}

// SubscriptionCard represents subscription card information
type SubscriptionCard struct {
	CardHolderName       string `json:"cardHolderName"`
	CardNumber           string `json:"cardNumber"`
	ExpireYear           string `json:"expireYear"`
	ExpireMonth          string `json:"expireMonth"`
	CVC                  string `json:"cvc"`
	CardUserKey          string `json:"cardUserKey"`
	CardToken            string `json:"cardToken"`
	UcsToken             string `json:"ucsToken"`
	ConsumerToken        string `json:"consumerToken"`
	RegisterConsumerCard *bool  `json:"registerConsumerCard,omitempty"`
}

// CardInformation represents card information for storage
type CardInformation struct {
	CardAlias      string `json:"cardAlias"`
	CardNumber     string `json:"cardNumber"`
	ExpireYear     string `json:"expireYear"`
	ExpireMonth    string `json:"expireMonth"`
	CardHolderName string `json:"cardHolderName"`
}

// BasketItem represents basket item information
type BasketItem struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Category1        string `json:"category1"`
	Category2        string `json:"category2"`
	ItemType         string `json:"itemType"`
	Price            string `json:"price"`
	SubMerchantKey   string `json:"subMerchantKey"`
	SubMerchantPrice string `json:"subMerchantPrice"`
	WithholdingTax   string `json:"withholdingTax"`
}

// PaymentItem represents payment item information
type PaymentItem struct {
	SubMerchantKey       string `json:"subMerchantKey"`
	PaymentTransactionID string `json:"paymentTransactionId"`
	SubMerchantPrice     string `json:"subMerchantPrice"`
	WithholdingTax       string `json:"withholdingTax"`
}

// BkmInstallmentPrice represents BKM installment price
type BkmInstallmentPrice struct {
	InstallmentNumber int    `json:"installmentNumber"`
	TotalPrice        string `json:"totalPrice"`
}

// BkmInstallment represents BKM installment information
type BkmInstallment struct {
	BankID            string                `json:"bankId"`
	InstallmentPrices []BkmInstallmentPrice `json:"installmentPrices"`
}

// Pagination represents pagination information
type Pagination struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	Page           int    `json:"page"`
	Count          int    `json:"count"`
}

// SubscriptionCustomer represents subscription customer information
type SubscriptionCustomer struct {
	Name            string               `json:"name"`
	Surname         string               `json:"surname"`
	IdentityNumber  string               `json:"identityNumber"`
	Email           string               `json:"email"`
	GsmNumber       string               `json:"gsmNumber"`
	BillingAddress  *SubscriptionAddress `json:"billingAddress"`
	ShippingAddress *SubscriptionAddress `json:"shippingAddress"`
}

// SubscriptionProduct represents subscription product information
type SubscriptionProduct struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
}

// SubscriptionPricingPlan represents subscription pricing plan information
type SubscriptionPricingPlan struct {
	Locale               string `json:"locale"`
	ConversationID       string `json:"conversationId"`
	Name                 string `json:"name"`
	Price                string `json:"price"`
	CurrencyCode         string `json:"currencyCode"`
	PaymentInterval      string `json:"paymentInterval"`
	PaymentIntervalCount int    `json:"paymentIntervalCount"`
	TrialPeriodDays      int    `json:"trialPeriodDays"`
	PlanPaymentType      string `json:"planPaymentType"`
}

// PaymentRequest represents payment request
type PaymentRequest struct {
	Locale          string        `json:"locale"`
	ConversationID  string        `json:"conversationId"`
	Price           string        `json:"price"`
	PaidPrice       string        `json:"paidPrice"`
	Currency        string        `json:"currency"`
	Installment     int           `json:"installment"`
	BasketID        string        `json:"basketId"`
	PaymentChannel  string        `json:"paymentChannel"`
	PaymentGroup    string        `json:"paymentGroup"`
	PaymentCard     *PaymentCard  `json:"paymentCard"`
	Buyer           *Buyer        `json:"buyer"`
	ShippingAddress *Address      `json:"shippingAddress"`
	BillingAddress  *Address      `json:"billingAddress"`
	BasketItems     []BasketItem  `json:"basketItems"`
	PaymentSource   string        `json:"paymentSource"`
	GsmNumber       string        `json:"gsmNumber"`
	PosOrderID      string        `json:"posOrderId"`
	ConnectorName   string        `json:"connectorName"`
	CallbackURL     string        `json:"callbackUrl"`
}

// BasicPaymentRequest represents basic payment request
type BasicPaymentRequest struct {
	Locale         string       `json:"locale"`
	ConversationID string       `json:"conversationId"`
	Price          string       `json:"price"`
	PaidPrice      string       `json:"paidPrice"`
	Installment    int          `json:"installment"`
	BuyerEmail     string       `json:"buyerEmail"`
	BuyerID        string       `json:"buyerId"`
	BuyerIP        string       `json:"buyerIp"`
	PosOrderID     string       `json:"posOrderId"`
	PaymentCard    *PaymentCard `json:"paymentCard"`
	Currency       string       `json:"currency"`
	ConnectorName  string       `json:"connectorName"`
	CallbackURL    string       `json:"callbackUrl"`
}

// APMRequest represents APM payment request
type APMRequest struct {
	Locale                  string        `json:"locale"`
	ConversationID          string        `json:"conversationId"`
	Price                   string        `json:"price"`
	PaidPrice               string        `json:"paidPrice"`
	PaymentChannel          string        `json:"paymentChannel"`
	PaymentGroup            string        `json:"paymentGroup"`
	PaymentSource           string        `json:"paymentSource"`
	Currency                string        `json:"currency"`
	MerchantOrderID         string        `json:"merchantOrderId"`
	CountryCode             string        `json:"countryCode"`
	AccountHolderName       string        `json:"accountHolderName"`
	MerchantCallbackURL     string        `json:"merchantCallbackUrl"`
	MerchantErrorURL        string        `json:"merchantErrorUrl"`
	MerchantNotificationURL string        `json:"merchantNotificationUrl"`
	APMType                 string        `json:"apmType"`
	BasketID                string        `json:"basketId"`
	Buyer                   *Buyer        `json:"buyer"`
	ShippingAddress         *Address      `json:"shippingAddress"`
	BillingAddress          *Address      `json:"billingAddress"`
	BasketItems             []BasketItem  `json:"basketItems"`
}

// RefundRequest represents refund request
type RefundRequest struct {
	Locale               string `json:"locale"`
	ConversationID       string `json:"conversationId"`
	PaymentTransactionID string `json:"paymentTransactionId"`
	Price                string `json:"price"`
	Currency             string `json:"currency"`
	IP                   string `json:"ip"`
	Reason               string `json:"reason"`
	Description          string `json:"description"`
}

// CancelRequest represents cancel request
type CancelRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	PaymentID      string `json:"paymentId"`
	IP             string `json:"ip"`
	Reason         string `json:"reason"`
	Description    string `json:"description"`
}

// PaymentResponse represents payment response
type PaymentResponse struct {
	Status              string `json:"status"`
	Locale              string `json:"locale"`
	SystemTime          int64  `json:"systemTime"`
	ConversationID      string `json:"conversationId"`
	Price               string `json:"price"`
	PaidPrice           string `json:"paidPrice"`
	Installment         int    `json:"installment"`
	PaymentID           string `json:"paymentId"`
	FraudStatus         int    `json:"fraudStatus"`
	MerchantCommission  string `json:"merchantCommissionRate"`
	IyziCommission      string `json:"iyziCommissionRateAmount"`
	IyziCommissionFee   string `json:"iyziCommissionFee"`
	CardType            string `json:"cardType"`
	CardAssociation     string `json:"cardAssociation"`
	CardFamily          string `json:"cardFamily"`
	CardToken           string `json:"cardToken"`
	CardUserKey         string `json:"cardUserKey"`
	BinNumber           string `json:"binNumber"`
	BasketID            string `json:"basketId"`
	Currency            string `json:"currency"`
	ItemTransactions    []ItemTransaction `json:"itemTransactions"`
	ConnectorName       string `json:"connectorName"`
	AuthCode            string `json:"authCode"`
	Phase               string `json:"phase"`
	LastFourDigits      string `json:"lastFourDigits"`
	PosOrderID          string `json:"posOrderId"`
	PaymentSource       string `json:"paymentSource"`
	ErrorCode           string `json:"errorCode"`
	ErrorMessage        string `json:"errorMessage"`
	ErrorGroup          string `json:"errorGroup"`
	Signature           string `json:"signature"`
}

// ItemTransaction represents item transaction
type ItemTransaction struct {
	ItemID                    string `json:"itemId"`
	PaymentTransactionID      string `json:"paymentTransactionId"`
	TransactionStatus         int    `json:"transactionStatus"`
	Price                     string `json:"price"`
	PaidPrice                 string `json:"paidPrice"`
	MerchantCommissionRate    string `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount string `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount  string `json:"iyziCommissionRateAmount"`
	IyziCommissionFee         string `json:"iyziCommissionFee"`
	BlockageRate              string `json:"blockageRate"`
	BlockageRateAmountMerchant string `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant string `json:"blockageRateAmountSubMerchant"`
	BlockageResolvedDate      string `json:"blockageResolvedDate"`
	SubMerchantKey            string `json:"subMerchantKey"`
	SubMerchantPrice          string `json:"subMerchantPrice"`
	SubMerchantPayoutRate     string `json:"subMerchantPayoutRate"`
	SubMerchantPayoutAmount   string `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount      string `json:"merchantPayoutAmount"`
	ConvertedPayout           ConvertedPayout `json:"convertedPayout"`
}

// ConvertedPayout represents converted payout information
type ConvertedPayout struct {
	PaidPrice                    string `json:"paidPrice"`
	IyziCommissionRateAmount     string `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            string `json:"iyziCommissionFee"`
	BlockageRateAmountMerchant   string `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant string `json:"blockageRateAmountSubMerchant"`
	SubMerchantPayoutAmount       string `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          string `json:"merchantPayoutAmount"`
	IyziConversionRate            string `json:"iyziConversionRate"`
	IyziConversionRateAmount      string `json:"iyziConversionRateAmount"`
	Currency                      string `json:"currency"`
}

// BaseResponse represents base response structure
type BaseResponse struct {
	Status         string `json:"status"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationID string `json:"conversationId"`
	ErrorCode      string `json:"errorCode"`
	ErrorMessage   string `json:"errorMessage"`
	ErrorGroup     string `json:"errorGroup"`
}

// APITestResponse represents API test response
type APITestResponse struct {
	BaseResponse
	SystemTime int64 `json:"systemTime"`
}