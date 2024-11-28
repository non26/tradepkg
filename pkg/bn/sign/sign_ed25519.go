package sign

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/non26/tradepkg/pkg/bn/utils"
)

type SignEd25519[T any] struct {
	private_key string
	queryString string
}

func NewSignEd25519[T any](privateKey string) IBinanceSignature[T] {
	return &SignEd25519[T]{
		private_key: privateKey,
	}
}

func (s *SignEd25519[T]) Sign(m *T, except_fields ...string) (string, error) {
	queryString := utils.CreateQueryStringFrom(m, except_fields...)
	s.queryString = queryString
	//private example
	//  "-----BEGIN PRIVATE KEY-----\nMC4CAQAwBQYDK2VwBCIEIKwzSGBlYbCE5cvr7ggQDU//tsiurzHbZSj7MM5ai8Aa\n-----END PRIVATE KEY-----"
	block, rest := pem.Decode([]byte(s.private_key))
	_ = rest
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", fmt.Errorf("failed to decode PEM block containing the private key")
	}

	// Parse the RSA private key
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	pk, ok := privateKey.(ed25519.PrivateKey)
	_ = pk
	if !ok {
		return "", fmt.Errorf("failed to convert private key to *rsa.PrivateKey")
	}

	// Sign the hashed payload using the private key
	signature := ed25519.Sign(pk, []byte(queryString))
	// Encode the signature in base64
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	return signatureBase64, nil
}

func (s *SignEd25519[T]) GetQueryStringBinanceSignature(m *T, except_fields ...string) (string, error) {
	signatureBase64, err := s.Sign(m, except_fields...)
	if err != nil {
		return "", err
	}
	return utils.GetQueryStringBinanceSignature(s.queryString, signatureBase64), nil
}
