package iyzipay

import (
	"context"
	"net/http"
)

// APITestService handles API test operations
type APITestService struct {
	client *Client
}

// Retrieve performs API test
func (s *APITestService) Retrieve(ctx context.Context) (*APITestResponse, error) {
	var response APITestResponse
	err := s.client.doRequest(ctx, http.MethodGet, EndpointAPITest, nil, &response)
	return &response, err
}

// PaymentService handles payment operations
type PaymentService struct {
	client *Client
}

// Create creates a new payment
func (s *PaymentService) Create(ctx context.Context, request *PaymentRequest) (*PaymentResponse, error) {
	var response PaymentResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPaymentAuth, request, &response)
	return &response, err
}

// Retrieve retrieves payment details
func (s *PaymentService) Retrieve(ctx context.Context, request *RetrievePaymentRequest) (*PaymentResponse, error) {
	var response PaymentResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPaymentDetail, request, &response)
	return &response, err
}

// BasicPaymentService handles basic payment operations
type BasicPaymentService struct {
	client *Client
}

// Create creates a new basic payment
func (s *BasicPaymentService) Create(ctx context.Context, request *BasicPaymentRequest) (*PaymentResponse, error) {
	var response PaymentResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPaymentAuthBasic, request, &response)
	return &response, err
}

// ThreedsInitializeService handles 3DS initialization
type ThreedsInitializeService struct {
	client *Client
}

// Create initializes 3DS payment
func (s *ThreedsInitializeService) Create(ctx context.Context, request *PaymentRequest) (*ThreedsInitializeResponse, error) {
	var response ThreedsInitializeResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPayment3DSecureInitialize, request, &response)
	return &response, err
}

// CreateBasic initializes basic 3DS payment
func (s *ThreedsInitializeService) CreateBasic(ctx context.Context, request *BasicPaymentRequest) (*ThreedsInitializeResponse, error) {
	var response ThreedsInitializeResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPayment3DSecureInitializeBasic, request, &response)
	return &response, err
}

// ThreedsPaymentService handles 3DS payment completion
type ThreedsPaymentService struct {
	client *Client
}

// Create completes 3DS payment
func (s *ThreedsPaymentService) Create(ctx context.Context, request *ThreedsPaymentRequest) (*PaymentResponse, error) {
	var response PaymentResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPayment3DSecureAuth, request, &response)
	return &response, err
}

// CreateBasic completes basic 3DS payment
func (s *ThreedsPaymentService) CreateBasic(ctx context.Context, request *ThreedsPaymentRequest) (*PaymentResponse, error) {
	var response PaymentResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPayment3DSecureAuthBasic, request, &response)
	return &response, err
}

// CheckoutFormService handles checkout form operations
type CheckoutFormService struct {
	client *Client
}

// Initialize initializes checkout form
func (s *CheckoutFormService) Initialize(ctx context.Context, request *CheckoutFormInitializeRequest) (*CheckoutFormInitializeResponse, error) {
	var response CheckoutFormInitializeResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointCheckoutFormInitializeAuth, request, &response)
	return &response, err
}

// Retrieve retrieves checkout form result
func (s *CheckoutFormService) Retrieve(ctx context.Context, request *RetrieveCheckoutFormRequest) (*CheckoutFormResponse, error) {
	var response CheckoutFormResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointCheckoutFormAuthDetail, request, &response)
	return &response, err
}

// CardService handles card storage operations
type CardService struct {
	client *Client
}

// Create creates/stores a card
func (s *CardService) Create(ctx context.Context, request *CreateCardRequest) (*CardResponse, error) {
	var response CardResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointCardStorageCard, request, &response)
	return &response, err
}

// Delete deletes a stored card
func (s *CardService) Delete(ctx context.Context, request *DeleteCardRequest) (*BaseResponse, error) {
	var response BaseResponse
	err := s.client.doRequest(ctx, http.MethodDelete, EndpointCardStorageCard, request, &response)
	return &response, err
}

// List retrieves list of stored cards
func (s *CardService) List(ctx context.Context, request *RetrieveCardListRequest) (*CardListResponse, error) {
	var response CardListResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointCardStorageCards, request, &response)
	return &response, err
}

// RefundService handles refund operations
type RefundService struct {
	client *Client
}

// Create creates a refund
func (s *RefundService) Create(ctx context.Context, request *RefundRequest) (*RefundResponse, error) {
	var response RefundResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPaymentRefund, request, &response)
	return &response, err
}

// CancelService handles payment cancellation
type CancelService struct {
	client *Client
}

// Create cancels a payment
func (s *CancelService) Create(ctx context.Context, request *CancelRequest) (*CancelResponse, error) {
	var response CancelResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPaymentCancel, request, &response)
	return &response, err
}

// SubMerchantService handles sub merchant operations
type SubMerchantService struct {
	client *Client
}

// Create creates a sub merchant
func (s *SubMerchantService) Create(ctx context.Context, request *CreateSubMerchantRequest) (*SubMerchantResponse, error) {
	var response SubMerchantResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointSubMerchant, request, &response)
	return &response, err
}

// Update updates a sub merchant
func (s *SubMerchantService) Update(ctx context.Context, request *UpdateSubMerchantRequest) (*SubMerchantResponse, error) {
	var response SubMerchantResponse
	err := s.client.doRequest(ctx, http.MethodPut, EndpointSubMerchant, request, &response)
	return &response, err
}

// Retrieve retrieves sub merchant details
func (s *SubMerchantService) Retrieve(ctx context.Context, request *RetrieveSubMerchantRequest) (*SubMerchantResponse, error) {
	var response SubMerchantResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointSubMerchantDetail, request, &response)
	return &response, err
}

// BKMService handles BKM Express operations
type BKMService struct {
	client *Client
}

// Initialize initializes BKM payment
func (s *BKMService) Initialize(ctx context.Context, request *BKMInitializeRequest) (*BKMInitializeResponse, error) {
	var response BKMInitializeResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointBKMInitialize, request, &response)
	return &response, err
}

// InitializeBasic initializes basic BKM payment
func (s *BKMService) InitializeBasic(ctx context.Context, request *BasicBKMInitializeRequest) (*BKMInitializeResponse, error) {
	var response BKMInitializeResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointBKMInitializeBasic, request, &response)
	return &response, err
}

// Retrieve retrieves BKM payment result
func (s *BKMService) Retrieve(ctx context.Context, request *RetrieveBKMRequest) (*BKMResponse, error) {
	var response BKMResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointBKMAuthDetail, request, &response)
	return &response, err
}

// APMService handles Alternative Payment Methods
type APMService struct {
	client *Client
}

// Initialize initializes APM payment
func (s *APMService) Initialize(ctx context.Context, request *APMRequest) (*APMInitializeResponse, error) {
	var response APMInitializeResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointAPMInitialize, request, &response)
	return &response, err
}

// Retrieve retrieves APM payment result
func (s *APMService) Retrieve(ctx context.Context, request *RetrieveAPMRequest) (*APMResponse, error) {
	var response APMResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointAPMRetrieve, request, &response)
	return &response, err
}

// SubscriptionService handles subscription operations
type SubscriptionService struct {
	client *Client
}

// Initialize initializes a subscription
func (s *SubscriptionService) Initialize(ctx context.Context, request *CreateSubscriptionInitRequest) (*SubscriptionInitializeResponse, error) {
	var response SubscriptionInitializeResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointSubscriptionInitialize, request, &response)
	return &response, err
}

// InstallmentInfoService handles installment information
type InstallmentInfoService struct {
	client *Client
}

// Retrieve retrieves installment information
func (s *InstallmentInfoService) Retrieve(ctx context.Context, request *RetrieveInstallmentInfoRequest) (*InstallmentInfoResponse, error) {
	var response InstallmentInfoResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPaymentInstallment, request, &response)
	return &response, err
}

// BinNumberService handles BIN number operations
type BinNumberService struct {
	client *Client
}

// Retrieve retrieves BIN number information
func (s *BinNumberService) Retrieve(ctx context.Context, request *RetrieveBinNumberRequest) (*BinNumberResponse, error) {
	var response BinNumberResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointPaymentBinCheck, request, &response)
	return &response, err
}

// PaymentItemService handles payment item operations
type PaymentItemService struct {
	client *Client
}

// Update updates payment item
func (s *PaymentItemService) Update(ctx context.Context, request *UpdatePaymentItemRequest) (*PaymentItemResponse, error) {
	var response PaymentItemResponse
	err := s.client.doRequest(ctx, http.MethodPut, EndpointPaymentItem, request, &response)
	return &response, err
}

// CrossBookingService handles cross booking operations
type CrossBookingService struct {
	client *Client
}

// Send sends cross booking
func (s *CrossBookingService) Send(ctx context.Context, request *CrossBookingRequest) (*CrossBookingResponse, error) {
	var response CrossBookingResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointCrossBookingSend, request, &response)
	return &response, err
}

// Receive receives cross booking
func (s *CrossBookingService) Receive(ctx context.Context, request *CrossBookingRequest) (*CrossBookingResponse, error) {
	var response CrossBookingResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointCrossBookingReceive, request, &response)
	return &response, err
}

// RefundToBalanceService handles refund to balance operations
type RefundToBalanceService struct {
	client *Client
}

// Create creates refund to balance
func (s *RefundToBalanceService) Create(ctx context.Context, request *RefundToBalanceRequest) (*RefundToBalanceResponse, error) {
	var response RefundToBalanceResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointRefundToBalance, request, &response)
	return &response, err
}

// SettlementToBalanceService handles settlement to balance operations
type SettlementToBalanceService struct {
	client *Client
}

// Create creates settlement to balance
func (s *SettlementToBalanceService) Create(ctx context.Context, request *SettlementToBalanceRequest) (*SettlementToBalanceResponse, error) {
	var response SettlementToBalanceResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointSettlementToBalance, request, &response)
	return &response, err
}

// UniversalCardStorageService handles universal card storage operations
type UniversalCardStorageService struct {
	client *Client
}

// Initialize initializes universal card storage
func (s *UniversalCardStorageService) Initialize(ctx context.Context, request *UniversalCardStorageInitializeRequest) (*UniversalCardStorageInitializeResponse, error) {
	var response UniversalCardStorageInitializeResponse
	err := s.client.doRequest(ctx, http.MethodPost, EndpointUniversalCardStorageInitialize, request, &response)
	return &response, err
}