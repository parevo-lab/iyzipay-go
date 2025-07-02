package iyzipay

import (
	"reflect"
	"strings"
	"testing"
)

func TestFormatPrice(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "string with decimal",
			input:    "10.5",
			expected: "10.5",
		},
		{
			name:     "string without decimal",
			input:    "10",
			expected: "10.0",
		},
		{
			name:     "float64 with decimal",
			input:    10.5,
			expected: "10.5",
		},
		{
			name:     "float64 without decimal",
			input:    10.0,
			expected: "10.0",
		},
		{
			name:     "int",
			input:    10,
			expected: "10.0",
		},
		{
			name:     "int64",
			input:    int64(10),
			expected: "10.0",
		},
		{
			name:     "float32",
			input:    float32(10.5),
			expected: "10.5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatPrice(tt.input)
			if result != tt.expected {
				t.Errorf("formatPrice(%v) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGenerateRandomString(t *testing.T) {
	size := 8
	result := generateRandomString(size)
	
	if len(result) < size {
		t.Errorf("Expected random string length >= %d, got %d", size, len(result))
	}

	// Test that multiple calls return different strings
	result2 := generateRandomString(size)
	if result == result2 {
		t.Error("Two consecutive calls should return different random strings")
	}
}

func TestGenerateHashV1(t *testing.T) {
	apiKey := "test-api-key"
	randomString := "12345678"
	secretKey := "test-secret"
	pkiString := "[locale=tr,conversationId=123]"

	hash := generateHashV1(apiKey, randomString, secretKey, pkiString)
	
	if hash == "" {
		t.Error("Hash should not be empty")
	}

	// Test that same inputs produce same hash
	hash2 := generateHashV1(apiKey, randomString, secretKey, pkiString)
	if hash != hash2 {
		t.Error("Same inputs should produce same hash")
	}

	// Test that different inputs produce different hash
	hash3 := generateHashV1(apiKey, "87654321", secretKey, pkiString)
	if hash == hash3 {
		t.Error("Different inputs should produce different hash")
	}
}

func TestGenerateHashV2(t *testing.T) {
	apiKey := "test-api-key"
	randomString := "12345678"
	secretKey := "test-secret"
	uri := "/payment/auth"
	body := map[string]interface{}{
		"locale":         "tr",
		"conversationId": "123",
	}

	hash := generateHashV2(apiKey, randomString, secretKey, uri, body)
	
	if hash == "" {
		t.Error("Hash should not be empty")
	}

	// Test that same inputs produce same hash
	hash2 := generateHashV2(apiKey, randomString, secretKey, uri, body)
	if hash != hash2 {
		t.Error("Same inputs should produce same hash")
	}

	// Test that different inputs produce different hash
	hash3 := generateHashV2(apiKey, "87654321", secretKey, uri, body)
	if hash == hash3 {
		t.Error("Different inputs should produce different hash")
	}
}

func TestGenerateAuthorizationHeaderV1(t *testing.T) {
	apiKey := "test-api-key"
	randomString := "12345678"
	secretKey := "test-secret"
	pkiString := "[locale=tr,conversationId=123]"

	header := generateAuthorizationHeaderV1(apiKey, randomString, secretKey, pkiString)
	
	if !strings.HasPrefix(header, HeaderIyziWSV1+" ") {
		t.Errorf("Header should start with '%s ', got %s", HeaderIyziWSV1, header)
	}

	if !strings.Contains(header, apiKey) {
		t.Error("Header should contain API key")
	}

	if !strings.Contains(header, Separator) {
		t.Error("Header should contain separator")
	}
}

func TestGenerateAuthorizationHeaderV2(t *testing.T) {
	apiKey := "test-api-key"
	randomString := "12345678"
	secretKey := "test-secret"
	uri := "/payment/auth"
	body := map[string]interface{}{
		"locale":         "tr",
		"conversationId": "123",
	}

	header := generateAuthorizationHeaderV2(apiKey, randomString, secretKey, uri, body)
	
	if !strings.HasPrefix(header, HeaderIyziWSV2+" ") {
		t.Errorf("Header should start with '%s ', got %s", HeaderIyziWSV2, header)
	}
}

func TestCalculateHMACSignature(t *testing.T) {
	params := []string{"param1", "param2", "param3"}
	secretKey := "test-secret"

	signature := calculateHMACSignature(params, secretKey)
	
	if signature == "" {
		t.Error("Signature should not be empty")
	}

	// Test that same inputs produce same signature
	signature2 := calculateHMACSignature(params, secretKey)
	if signature != signature2 {
		t.Error("Same inputs should produce same signature")
	}

	// Test that different inputs produce different signature
	differentParams := []string{"param1", "param2", "param4"}
	signature3 := calculateHMACSignature(differentParams, secretKey)
	if signature == signature3 {
		t.Error("Different inputs should produce different signature")
	}
}

func TestPKIString(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: "",
		},
		{
			name: "simple struct",
			input: struct {
				Locale         string `json:"locale"`
				ConversationID string `json:"conversationId"`
			}{
				Locale:         "tr",
				ConversationID: "123",
			},
			expected: "[locale=tr,conversationId=123]",
		},
		{
			name: "struct with price",
			input: struct {
				Price string `json:"price"`
			}{
				Price: "10",
			},
			expected: "[price=10.0]",
		},
		{
			name:     "slice",
			input:    []string{"item1", "item2"},
			expected: "[item1,item2]",
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PKIString(tt.input)
			if result != tt.expected {
				t.Errorf("PKIString(%v) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMergeObjects(t *testing.T) {
	obj1 := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	obj2 := map[string]interface{}{
		"key2": "new_value2", // Should override
		"key3": "value3",
	}

	result := mergeObjects(obj1, obj2)

	if result["key1"] != "value1" {
		t.Errorf("Expected key1=value1, got %v", result["key1"])
	}

	if result["key2"] != "new_value2" {
		t.Errorf("Expected key2=new_value2, got %v", result["key2"])
	}

	if result["key3"] != "value3" {
		t.Errorf("Expected key3=value3, got %v", result["key3"])
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(result))
	}
}

func TestValidateConfigFunction(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				BaseURL:   "https://test.com",
				APIKey:    "test-key",
				SecretKey: "test-secret",
			},
			wantErr: false,
		},
		{
			name: "empty base URL",
			config: &Config{
				BaseURL:   "",
				APIKey:    "test-key",
				SecretKey: "test-secret",
			},
			wantErr: true,
		},
		{
			name: "empty API key",
			config: &Config{
				BaseURL:   "https://test.com",
				APIKey:    "",
				SecretKey: "test-secret",
			},
			wantErr: true,
		},
		{
			name: "empty secret key",
			config: &Config{
				BaseURL:   "https://test.com",
				APIKey:    "test-key",
				SecretKey: "",
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

func TestFlexibleUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		target   interface{}
		expected interface{}
		wantErr  bool
	}{
		{
			name:     "string to string (normal case)",
			jsonData: `{"price": "10.50"}`,
			target:   &struct{ Price string `json:"price"` }{},
			expected: &struct{ Price string `json:"price"` }{Price: "10.50"},
			wantErr:  false,
		},
		{
			name:     "number to string (flexible case)",
			jsonData: `{"price": 10.50}`,
			target:   &struct{ Price string `json:"price"` }{},
			expected: &struct{ Price string `json:"price"` }{Price: "10.5"},
			wantErr:  false,
		},
		{
			name:     "integer to string",
			jsonData: `{"price": 10}`,
			target:   &struct{ Price string `json:"price"` }{},
			expected: &struct{ Price string `json:"price"` }{Price: "10"},
			wantErr:  false,
		},
		{
			name:     "string to int",
			jsonData: `{"count": "42"}`,
			target:   &struct{ Count int `json:"count"` }{},
			expected: &struct{ Count int `json:"count"` }{Count: 42},
			wantErr:  false,
		},
		{
			name:     "float to int",
			jsonData: `{"count": 42.0}`,
			target:   &struct{ Count int `json:"count"` }{},
			expected: &struct{ Count int `json:"count"` }{Count: 42},
			wantErr:  false,
		},
		{
			name:     "bool to string",
			jsonData: `{"active": true}`,
			target:   &struct{ Active string `json:"active"` }{},
			expected: &struct{ Active string `json:"active"` }{Active: "true"},
			wantErr:  false,
		},
		{
			name:     "string to bool",
			jsonData: `{"active": "true"}`,
			target:   &struct{ Active bool `json:"active"` }{},
			expected: &struct{ Active bool `json:"active"` }{Active: true},
			wantErr:  false,
		},
		{
			name:     "complex payment response",
			jsonData: `{"status": "success", "price": 10.50, "installment": 1, "fraudStatus": "1", "paymentId": "12345"}`,
			target: &struct {
				Status      string `json:"status"`
				Price       string `json:"price"`
				Installment int    `json:"installment"`
				FraudStatus int    `json:"fraudStatus"`
				PaymentID   string `json:"paymentId"`
			}{},
			expected: &struct {
				Status      string `json:"status"`
				Price       string `json:"price"`
				Installment int    `json:"installment"`
				FraudStatus int    `json:"fraudStatus"`
				PaymentID   string `json:"paymentId"`
			}{
				Status:      "success",
				Price:       "10.5",
				Installment: 1,
				FraudStatus: 1,
				PaymentID:   "12345",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FlexibleUnmarshal([]byte(tt.jsonData), tt.target)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("FlexibleUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Use reflection to compare the results
				if !compareStructs(tt.target, tt.expected) {
					t.Errorf("FlexibleUnmarshal() result = %+v, expected %+v", tt.target, tt.expected)
				}
			}
		})
	}
}

func compareStructs(a, b interface{}) bool {
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	
	if aValue.Kind() == reflect.Ptr {
		aValue = aValue.Elem()
	}
	if bValue.Kind() == reflect.Ptr {
		bValue = bValue.Elem()
	}
	
	if aValue.Type() != bValue.Type() {
		return false
	}
	
	for i := 0; i < aValue.NumField(); i++ {
		aField := aValue.Field(i)
		bField := bValue.Field(i)
		
		if !reflect.DeepEqual(aField.Interface(), bField.Interface()) {
			return false
		}
	}
	
	return true
}