package iyzipay

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// formatPrice formats price to string format with decimal point
func formatPrice(price interface{}) string {
	switch v := price.(type) {
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return formatFloatPrice(f)
		}
		return v
	case float64:
		return formatFloatPrice(v)
	case float32:
		return formatFloatPrice(float64(v))
	case int:
		return fmt.Sprintf("%d.0", v)
	case int64:
		return fmt.Sprintf("%d.0", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func formatFloatPrice(f float64) string {
	if math.Trunc(f) == f {
		return fmt.Sprintf("%.0f.0", f)
	}
	return fmt.Sprintf("%g", f)
}

// generateRandomString generates a random string for authentication
func generateRandomString(size int) string {
	now := time.Now().UnixNano()
	b := make([]byte, 8)
	rand.Read(b)
	randomPart := hex.EncodeToString(b)
	return fmt.Sprintf("%d%s", now, randomPart[:size])
}

// generateAuthorizationHeaderV1 generates authorization header version 1 (fallback)
func generateAuthorizationHeaderV1(apiKey, randomString, secretKey, pkiString string) string {
	hash := generateHashV1(apiKey, randomString, secretKey, pkiString)
	return fmt.Sprintf("%s %s%s%s", HeaderIyziWSV1, apiKey, Separator, hash)
}

// generateAuthorizationHeaderV2 generates authorization header version 2 (primary)
func generateAuthorizationHeaderV2(apiKey, randomString, secretKey, uri string, body interface{}) string {
	signature := generateHashV2(apiKey, randomString, secretKey, uri, body)
	authParams := []string{
		fmt.Sprintf("apiKey%s%s", Separator, apiKey),
		fmt.Sprintf("randomKey%s%s", Separator, randomString),
		fmt.Sprintf("signature%s%s", Separator, signature),
	}
	authString := strings.Join(authParams, "&")
	return fmt.Sprintf("%s %s", HeaderIyziWSV2, base64.StdEncoding.EncodeToString([]byte(authString)))
}

// generateHashV1 generates SHA1 hash for authorization v1
func generateHashV1(apiKey, randomString, secretKey, pkiString string) string {
	data := apiKey + randomString + secretKey + pkiString
	h := sha1.New()
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// generateHashV2 generates HMAC-SHA256 hash for authorization v2
func generateHashV2(apiKey, randomString, secretKey, uri string, body interface{}) string {
	bodyJSON, _ := json.Marshal(body)
	data := randomString + uri + string(bodyJSON)
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// calculateHMACSignature calculates HMAC-SHA256 signature for verification
func calculateHMACSignature(params []string, secretKey string) string {
	dataToCheck := strings.Join(params, ":")
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(dataToCheck))
	return hex.EncodeToString(h.Sum(nil))
}

// PKIString generates PKI string from request object
func PKIString(obj interface{}) string {
	return generatePKIString(obj)
}

func generatePKIString(obj interface{}) string {
	if obj == nil {
		return ""
	}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		return structToPKI(v)
	case reflect.Slice, reflect.Array:
		return sliceToPKI(v)
	case reflect.Map:
		return mapToPKI(v)
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}

func structToPKI(v reflect.Value) string {
	t := v.Type()
	var pairs []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}

		// Get JSON tag name
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Parse JSON tag (remove omitempty and other options)
		tagParts := strings.Split(jsonTag, ",")
		fieldName := tagParts[0]

		// Skip if field is nil pointer or empty
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		value := generatePKIString(field.Interface())
		if value != "" {
			// Format price fields
			if strings.Contains(strings.ToLower(fieldName), "price") || 
			   strings.Contains(strings.ToLower(fieldName), "amount") {
				if field.Kind() == reflect.String {
					value = formatPrice(value)
				}
			}
			pairs = append(pairs, fmt.Sprintf("%s=%s", fieldName, value))
		}
	}

	if len(pairs) == 0 {
		return ""
	}

	return fmt.Sprintf("[%s]", strings.Join(pairs, ","))
}

func sliceToPKI(v reflect.Value) string {
	if v.Len() == 0 {
		return ""
	}

	var items []string
	for i := 0; i < v.Len(); i++ {
		item := generatePKIString(v.Index(i).Interface())
		if item != "" {
			items = append(items, item)
		}
	}

	if len(items) == 0 {
		return ""
	}

	return fmt.Sprintf("[%s]", strings.Join(items, ","))
}

func mapToPKI(v reflect.Value) string {
	if v.Len() == 0 {
		return ""
	}

	var pairs []string
	keys := v.MapKeys()
	
	// Sort keys for consistent output
	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprintf("%v", keys[i].Interface()) < fmt.Sprintf("%v", keys[j].Interface())
	})

	for _, key := range keys {
		value := generatePKIString(v.MapIndex(key).Interface())
		if value != "" {
			pairs = append(pairs, fmt.Sprintf("%v=%s", key.Interface(), value))
		}
	}

	if len(pairs) == 0 {
		return ""
	}

	return fmt.Sprintf("[%s]", strings.Join(pairs, ","))
}

// mergeObjects merges two maps/objects
func mergeObjects(obj1, obj2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	
	for k, v := range obj1 {
		merged[k] = v
	}
	
	for k, v := range obj2 {
		merged[k] = v
	}
	
	return merged
}

// validateConfig validates the client configuration
func validateConfig(config *Config) error {
	if config.BaseURL == "" {
		return fmt.Errorf("baseURL cannot be empty")
	}
	if config.APIKey == "" {
		return fmt.Errorf("apiKey cannot be empty")
	}
	if config.SecretKey == "" {
		return fmt.Errorf("secretKey cannot be empty")
	}
	return nil
}