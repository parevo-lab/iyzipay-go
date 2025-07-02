package iyzipay

// CalculateHMACSignature calculates HMAC-SHA256 signature for verification
// This function is exported for use in examples and external verification
func CalculateHMACSignature(params []string, secretKey string) string {
	return calculateHMACSignature(params, secretKey)
}