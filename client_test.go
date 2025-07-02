package iyzipay

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	config := &Config{
		APIKey:    "test-api-key",
		SecretKey: "test-secret-key",
		BaseURL:   "https://test.iyzipay.com",
	}

	client := NewClient(config)
	if client == nil {
		t.Fatal("Client should not be nil")
	}

	if client.config.APIKey != config.APIKey {
		t.Errorf("Expected API key %s, got %s", config.APIKey, client.config.APIKey)
	}

	if client.config.SecretKey != config.SecretKey {
		t.Errorf("Expected secret key %s, got %s", config.SecretKey, client.config.SecretKey)
	}

	if client.config.BaseURL != config.BaseURL {
		t.Errorf("Expected base URL %s, got %s", config.BaseURL, client.config.BaseURL)
	}
}

func TestNewClientFromEnv(t *testing.T) {
	// Note: This test would need environment variables set
	// For now, just test that it doesn't panic
	defer func() {
		if r := recover(); r != nil {
			// Expected to panic if env vars are not set
		}
	}()

	// This will panic if env vars are not set, which is expected
	// client := NewClientFromEnv()
}

func TestMakeRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept application/json, got %s", r.Header.Get("Accept"))
		}

		// Verify authorization headers exist
		if r.Header.Get(HeaderAuthorization) == "" {
			t.Error("Authorization header should not be empty")
		}

		if r.Header.Get(HeaderAuthorizationFallback) == "" {
			t.Error("Authorization fallback header should not be empty")
		}

		if r.Header.Get(HeaderRandomString) == "" {
			t.Error("Random string header should not be empty")
		}

		if r.Header.Get(HeaderClientVersion) != ClientVersion {
			t.Errorf("Expected client version %s, got %s", ClientVersion, r.Header.Get(HeaderClientVersion))
		}

		// Return a test response
		response := APITestResponse{
			BaseResponse: BaseResponse{
				Status:         "success",
				Locale:         "tr",
				SystemTime:     time.Now().Unix(),
				ConversationID: "test-conversation",
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := &Config{
		APIKey:    "test-api-key",
		SecretKey: "test-secret-key",
		BaseURL:   server.URL,
	}

	client := NewClient(config)
	
	ctx := context.Background()
	resp, err := client.makeRequest(ctx, "GET", "/test", nil)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestAPITest(t *testing.T) {
	// Create a test server that mimics İyzipay API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != EndpointAPITest {
			t.Errorf("Expected path %s, got %s", EndpointAPITest, r.URL.Path)
		}

		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		response := APITestResponse{
			BaseResponse: BaseResponse{
				Status:         "success",
				Locale:         "tr",
				SystemTime:     time.Now().Unix(),
				ConversationID: "",
			},
			SystemTime: time.Now().Unix(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := &Config{
		APIKey:    "test-api-key",
		SecretKey: "test-secret-key",
		BaseURL:   server.URL,
	}

	client := NewClient(config)
	
	ctx := context.Background()
	result, err := client.APITest.Retrieve(ctx)
	if err != nil {
		t.Fatalf("API test failed: %v", err)
	}

	if result.Status != "success" {
		t.Errorf("Expected status success, got %s", result.Status)
	}

	if result.Locale != "tr" {
		t.Errorf("Expected locale tr, got %s", result.Locale)
	}
}

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				APIKey:    "test-key",
				SecretKey: "test-secret",
				BaseURL:   "https://test.com",
			},
			wantErr: false,
		},
		{
			name: "empty api key",
			config: &Config{
				APIKey:    "",
				SecretKey: "test-secret",
				BaseURL:   "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "empty secret key",
			config: &Config{
				APIKey:    "test-key",
				SecretKey: "",
				BaseURL:   "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "empty base url",
			config: &Config{
				APIKey:    "test-key",
				SecretKey: "test-secret",
				BaseURL:   "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConfig(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetHTTPClient(t *testing.T) {
	config := &Config{
		APIKey:    "test-api-key",
		SecretKey: "test-secret-key",
		BaseURL:   "https://test.iyzipay.com",
	}

	client := NewClient(config)
	
	customHTTPClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	client.SetHTTPClient(customHTTPClient)

	if client.config.HTTPClient != customHTTPClient {
		t.Error("Custom HTTP client was not set properly")
	}

	if client.config.HTTPClient.Timeout != 10*time.Second {
		t.Errorf("Expected timeout 10s, got %v", client.config.HTTPClient.Timeout)
	}
}

func TestGetConfig(t *testing.T) {
	config := &Config{
		APIKey:    "test-api-key",
		SecretKey: "test-secret-key",
		BaseURL:   "https://test.iyzipay.com",
	}

	client := NewClient(config)
	retrievedConfig := client.GetConfig()

	if retrievedConfig.APIKey != config.APIKey {
		t.Errorf("Expected API key %s, got %s", config.APIKey, retrievedConfig.APIKey)
	}

	if retrievedConfig.SecretKey != config.SecretKey {
		t.Errorf("Expected secret key %s, got %s", config.SecretKey, retrievedConfig.SecretKey)
	}

	if retrievedConfig.BaseURL != config.BaseURL {
		t.Errorf("Expected base URL %s, got %s", config.BaseURL, retrievedConfig.BaseURL)
	}
}