package bn

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type SignHMACSHA256[T any] struct {
	apiKey      string
	secretKey   string
	queryString string
}

func NewSignHMACSHA256[T any](apiKey, secretKey string) IBinanceSignature[T] {
	return &SignHMACSHA256[T]{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (s *SignHMACSHA256[T]) Sign(m *T, except_fields ...string) (string, error) {
	queryString := CreateQueryStringFrom(m, except_fields...)
	s.queryString = queryString
	payload := []byte(queryString)
	mac := hmac.New(sha256.New, []byte(s.secretKey))
	mac.Write(payload)
	bnsign := hex.EncodeToString(mac.Sum(nil))
	return bnsign, nil
}

func (s *SignHMACSHA256[T]) GetQueryStringBinanceSignature(m *T, except_fields ...string) (string, error) {
	signature, err := s.Sign(m, except_fields...)
	if err != nil {
		return "", err
	}
	return GetQueryStringBinanceSignature(s.queryString, signature), nil
}
