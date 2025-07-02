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

// FlexibleUnmarshal performs flexible JSON unmarshal that handles type mismatches
func FlexibleUnmarshal(data []byte, v interface{}) error {
	// First try normal unmarshal
	if err := json.Unmarshal(data, v); err == nil {
		return nil
	}

	// If normal unmarshal fails, try flexible unmarshal
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		return err
	}

	return flexibleMapToStruct(rawData, v)
}

// flexibleMapToStruct converts map to struct with flexible type handling
func flexibleMapToStruct(data map[string]interface{}, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("destination must be a non-nil pointer")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to struct")
	}

	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i)

		if !field.CanSet() {
			continue
		}

		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		tagParts := strings.Split(jsonTag, ",")
		fieldName := tagParts[0]

		value, exists := data[fieldName]
		if !exists {
			continue
		}

		if err := setFlexibleValue(field, value); err != nil {
			return fmt.Errorf("error setting field %s: %v", fieldName, err)
		}
	}

	return nil
}

// setFlexibleValue sets value to field with flexible type conversion
func setFlexibleValue(field reflect.Value, value interface{}) error {
	if value == nil {
		return nil
	}

	targetType := field.Type()
	sourceValue := reflect.ValueOf(value)

	// Handle pointer types
	if targetType.Kind() == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(targetType.Elem()))
		}
		return setFlexibleValue(field.Elem(), value)
	}

	// Handle same types
	if sourceValue.Type().AssignableTo(targetType) {
		field.Set(sourceValue)
		return nil
	}

	// Handle conversions
	switch targetType.Kind() {
	case reflect.String:
		return setStringValue(field, value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setIntValue(field, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUintValue(field, value)
	case reflect.Float32, reflect.Float64:
		return setFloatValue(field, value)
	case reflect.Bool:
		return setBoolValue(field, value)
	case reflect.Slice:
		return setSliceValue(field, value)
	case reflect.Struct:
		return setStructValue(field, value)
	default:
		return fmt.Errorf("unsupported type conversion from %T to %s", value, targetType)
	}
}

// setStringValue converts any value to string
func setStringValue(field reflect.Value, value interface{}) error {
	switch v := value.(type) {
	case string:
		field.SetString(v)
	case int, int8, int16, int32, int64:
		field.SetString(fmt.Sprintf("%d", v))
	case uint, uint8, uint16, uint32, uint64:
		field.SetString(fmt.Sprintf("%d", v))
	case float32, float64:
		field.SetString(fmt.Sprintf("%g", v))
	case bool:
		field.SetString(fmt.Sprintf("%t", v))
	default:
		field.SetString(fmt.Sprintf("%v", v))
	}
	return nil
}

// setIntValue converts any value to int
func setIntValue(field reflect.Value, value interface{}) error {
	switch v := value.(type) {
	case int:
		field.SetInt(int64(v))
	case int8:
		field.SetInt(int64(v))
	case int16:
		field.SetInt(int64(v))
	case int32:
		field.SetInt(int64(v))
	case int64:
		field.SetInt(v)
	case uint, uint8, uint16, uint32, uint64:
		field.SetInt(int64(reflect.ValueOf(v).Uint()))
	case float32:
		field.SetInt(int64(v))
	case float64:
		field.SetInt(int64(v))
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			field.SetInt(i)
		} else if f, err := strconv.ParseFloat(v, 64); err == nil {
			field.SetInt(int64(f))
		} else {
			return fmt.Errorf("cannot convert string %q to int", v)
		}
	default:
		return fmt.Errorf("cannot convert %T to int", v)
	}
	return nil
}

// setUintValue converts any value to uint
func setUintValue(field reflect.Value, value interface{}) error {
	switch v := value.(type) {
	case uint:
		field.SetUint(uint64(v))
	case uint8:
		field.SetUint(uint64(v))
	case uint16:
		field.SetUint(uint64(v))
	case uint32:
		field.SetUint(uint64(v))
	case uint64:
		field.SetUint(v)
	case int, int8, int16, int32, int64:
		val := reflect.ValueOf(v).Int()
		if val < 0 {
			return fmt.Errorf("cannot convert negative int %d to uint", val)
		}
		field.SetUint(uint64(val))
	case float32:
		if v < 0 {
			return fmt.Errorf("cannot convert negative float %f to uint", v)
		}
		field.SetUint(uint64(v))
	case float64:
		if v < 0 {
			return fmt.Errorf("cannot convert negative float %f to uint", v)
		}
		field.SetUint(uint64(v))
	case string:
		if i, err := strconv.ParseUint(v, 10, 64); err == nil {
			field.SetUint(i)
		} else if f, err := strconv.ParseFloat(v, 64); err == nil {
			if f < 0 {
				return fmt.Errorf("cannot convert negative float %f to uint", f)
			}
			field.SetUint(uint64(f))
		} else {
			return fmt.Errorf("cannot convert string %q to uint", v)
		}
	default:
		return fmt.Errorf("cannot convert %T to uint", v)
	}
	return nil
}

// setFloatValue converts any value to float
func setFloatValue(field reflect.Value, value interface{}) error {
	switch v := value.(type) {
	case float32:
		field.SetFloat(float64(v))
	case float64:
		field.SetFloat(v)
	case int, int8, int16, int32, int64:
		field.SetFloat(float64(reflect.ValueOf(v).Int()))
	case uint, uint8, uint16, uint32, uint64:
		field.SetFloat(float64(reflect.ValueOf(v).Uint()))
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			field.SetFloat(f)
		} else {
			return fmt.Errorf("cannot convert string %q to float", v)
		}
	default:
		return fmt.Errorf("cannot convert %T to float", v)
	}
	return nil
}

// setBoolValue converts any value to bool
func setBoolValue(field reflect.Value, value interface{}) error {
	switch v := value.(type) {
	case bool:
		field.SetBool(v)
	case string:
		if b, err := strconv.ParseBool(v); err == nil {
			field.SetBool(b)
		} else {
			return fmt.Errorf("cannot convert string %q to bool", v)
		}
	case int, int8, int16, int32, int64:
		field.SetBool(reflect.ValueOf(v).Int() != 0)
	case uint, uint8, uint16, uint32, uint64:
		field.SetBool(reflect.ValueOf(v).Uint() != 0)
	case float32, float64:
		field.SetBool(reflect.ValueOf(v).Float() != 0)
	default:
		return fmt.Errorf("cannot convert %T to bool", v)
	}
	return nil
}

// setSliceValue converts slice values
func setSliceValue(field reflect.Value, value interface{}) error {
	sourceSlice := reflect.ValueOf(value)
	if sourceSlice.Kind() != reflect.Slice {
		return fmt.Errorf("source is not a slice")
	}

	elementType := field.Type().Elem()
	newSlice := reflect.MakeSlice(field.Type(), sourceSlice.Len(), sourceSlice.Len())

	for i := 0; i < sourceSlice.Len(); i++ {
		sourceElem := sourceSlice.Index(i)
		targetElem := newSlice.Index(i)

		if elementType.Kind() == reflect.Ptr {
			newElem := reflect.New(elementType.Elem())
			if err := setFlexibleValue(newElem.Elem(), sourceElem.Interface()); err != nil {
				return err
			}
			targetElem.Set(newElem)
		} else {
			if err := setFlexibleValue(targetElem, sourceElem.Interface()); err != nil {
				return err
			}
		}
	}

	field.Set(newSlice)
	return nil
}

// setStructValue converts struct values
func setStructValue(field reflect.Value, value interface{}) error {
	if mapValue, ok := value.(map[string]interface{}); ok {
		return flexibleMapToStruct(mapValue, field.Addr().Interface())
	}
	return fmt.Errorf("cannot convert %T to struct", value)
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