package iyzipay

// CalculateHMACSignature calculates HMAC-SHA256 signature for verification
// This function is exported for use in examples and external verification
func CalculateHMACSignature(params []string, secretKey string) string {
	return calculateHMACSignature(params, secretKey)
}

// FlexibleJSONUnmarshal performs flexible JSON unmarshal that handles type mismatches
// This function is exported for use when manually handling JSON responses
func FlexibleJSONUnmarshal(data []byte, v interface{}) error {
	return FlexibleUnmarshal(data, v)
}