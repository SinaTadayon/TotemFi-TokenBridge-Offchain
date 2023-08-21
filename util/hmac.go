package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Signer signs provided payloads.
type Signer interface {
	// Sign signs provided payload and returns encoded string sum.
	Sign(payload []byte) string

	Verify(payload []byte, hash string) bool
}

// hmacSigner uses HMAC SHA256 for signing payloads.
type HmacSigner struct {
	ApiKey    string
	SecretKey []byte
}

func NewHmacSignerFromConfig(config *Config) (*HmacSigner, error) {
	apiKey, secretKey := "", ""
	apiKey, secretKey = config.KeyManagerConfig.APIKey, config.KeyManagerConfig.HMACKey

	return NewHmacSigner(apiKey, secretKey), nil
}

func NewHmacSigner(apiKey string, secretKey string) *HmacSigner {
	return &HmacSigner{
		ApiKey:    apiKey,
		SecretKey: []byte(secretKey),
	}
}

// Sign signs provided payload and returns encoded string sum.
func (hs *HmacSigner) Sign(payload []byte) string {
	mac := hmac.New(sha256.New, hs.SecretKey)
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

func (hs *HmacSigner) Verify(payload []byte, hash string) bool {
	mac := hmac.New(sha256.New, hs.SecretKey)
	mac.Write(payload)
	res := hex.EncodeToString(mac.Sum(nil))
	return hash == res
}
