package iyzipay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Config represents the configuration for İyzipay client
type Config struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	HTTPClient *http.Client
}

// Client represents the İyzipay API client
type Client struct {
	config *Config

	// Services
	APITest                    *APITestService
	Payment                    *PaymentService
	BasicPayment               *BasicPaymentService
	ThreedsInitialize          *ThreedsInitializeService
	ThreedsPayment             *ThreedsPaymentService
	CheckoutForm               *CheckoutFormService
	Card                       *CardService
	Refund                     *RefundService
	Cancel                     *CancelService
	SubMerchant                *SubMerchantService
	BKM                        *BKMService
	APM                        *APMService
	Subscription               *SubscriptionService
	InstallmentInfo            *InstallmentInfoService
	BinNumber                  *BinNumberService
	PaymentItem                *PaymentItemService
	CrossBooking               *CrossBookingService
	RefundToBalance            *RefundToBalanceService
	SettlementToBalance        *SettlementToBalanceService
	UniversalCardStorage       *UniversalCardStorageService
}

// NewClient creates a new İyzipay client with the given configuration
func NewClient(config *Config) *Client {
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	if err := validateConfig(config); err != nil {
		panic(fmt.Sprintf("invalid config: %v", err))
	}

	client := &Client{
		config: config,
	}

	// Initialize services
	client.APITest = &APITestService{client: client}
	client.Payment = &PaymentService{client: client}
	client.BasicPayment = &BasicPaymentService{client: client}
	client.ThreedsInitialize = &ThreedsInitializeService{client: client}
	client.ThreedsPayment = &ThreedsPaymentService{client: client}
	client.CheckoutForm = &CheckoutFormService{client: client}
	client.Card = &CardService{client: client}
	client.Refund = &RefundService{client: client}
	client.Cancel = &CancelService{client: client}
	client.SubMerchant = &SubMerchantService{client: client}
	client.BKM = &BKMService{client: client}
	client.APM = &APMService{client: client}
	client.Subscription = &SubscriptionService{client: client}
	client.InstallmentInfo = &InstallmentInfoService{client: client}
	client.BinNumber = &BinNumberService{client: client}
	client.PaymentItem = &PaymentItemService{client: client}
	client.CrossBooking = &CrossBookingService{client: client}
	client.RefundToBalance = &RefundToBalanceService{client: client}
	client.SettlementToBalance = &SettlementToBalanceService{client: client}
	client.UniversalCardStorage = &UniversalCardStorageService{client: client}

	return client
}

// NewClientFromEnv creates a new İyzipay client using environment variables
func NewClientFromEnv() *Client {
	config := &Config{
		APIKey:    os.Getenv("IYZIPAY_API_KEY"),
		SecretKey: os.Getenv("IYZIPAY_SECRET_KEY"),
		BaseURL:   os.Getenv("IYZIPAY_BASE_URL"),
	}

	// Set default base URL if not provided
	if config.BaseURL == "" {
		config.BaseURL = "https://sandbox-api.iyzipay.com"
	}

	return NewClient(config)
}

// makeRequest performs HTTP request with İyzipay authentication
func (c *Client) makeRequest(ctx context.Context, method, endpoint string, body interface{}) (*http.Response, error) {
	url := c.config.BaseURL + endpoint
	
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	// Generate authentication headers
	randomString := generateRandomString(RandomStringSize)
	req.Header.Set(HeaderRandomString, randomString)
	req.Header.Set(HeaderClientVersion, ClientVersion)

	// Generate PKI string for v1 auth
	pkiString := ""
	if body != nil {
		pkiString = PKIString(body)
	}

	// Set both v1 and v2 authorization headers
	authV1 := generateAuthorizationHeaderV1(c.config.APIKey, randomString, c.config.SecretKey, pkiString)
	authV2 := generateAuthorizationHeaderV2(c.config.APIKey, randomString, c.config.SecretKey, endpoint, body)
	
	req.Header.Set(HeaderAuthorization, authV2)
	req.Header.Set(HeaderAuthorizationFallback, authV1)

	return c.config.HTTPClient.Do(req)
}

// doRequest performs the request and handles the response
func (c *Client) doRequest(ctx context.Context, method, endpoint string, body interface{}, result interface{}) error {
	resp, err := c.makeRequest(ctx, method, endpoint, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(respBody))
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// SetHTTPClient sets custom HTTP client
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.config.HTTPClient = httpClient
}

// GetConfig returns the client configuration
func (c *Client) GetConfig() *Config {
	return c.config
}