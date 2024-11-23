package bn

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"reflect"
)

type IBinanceSignature[T any] interface {
	Sign(payload []byte, key []byte) string
	CreateBinanceSignature(data *url.Values) string
	GetQueryStringFromStructType(m *T) url.Values
}

type binanceSignature[T any] struct {
	apiKey    string
	secretKey string
}

func NewBinanceNonWsSignature[T any](apiKey, secretKey string) IBinanceSignature[T] {
	return &binanceSignature[T]{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (b *binanceSignature[T]) GetQueryStringFromStructType(m *T) url.Values {
	st := reflect.TypeOf(m).Elem()
	v := reflect.ValueOf(m).Elem()
	q := url.Values{}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i).Tag.Get("json")
		value := v.FieldByIndex([]int{i}).String()
		q.Add(field, value)
	}
	return q
}

func (b *binanceSignature[T]) CreateBinanceSignature(data *url.Values) string {
	payload := data.Encode()
	encodeString := b.Sign([]byte(payload), []byte(b.secretKey))
	encodeData := fmt.Sprintf("%v&signature=%v", data.Encode(), encodeString)
	return encodeData
}

func (b *binanceSignature[T]) Sign(payload []byte, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}
