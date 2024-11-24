package bn

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"reflect"
	"slices"
)

type IBinanceSignature[T any] interface {
	GetSignature(key []byte) string
	GetQueryStringBinanceSignature() string
	CreateQueryStringFromPayload(m *T) url.Values
}

type binanceSignature[T any] struct {
	apiKey         string
	secretKey      string
	bnsign         string
	exclude_fields []string
	encodepayload  string
}

func NewBinanceNonWsSignature[T any](apiKey, secretKey string) IBinanceSignature[T] {
	return &binanceSignature[T]{
		apiKey:         apiKey,
		secretKey:      secretKey,
		exclude_fields: []string{},
	}
}

func (b *binanceSignature[T]) CreateQueryStringFromPayload(m *T) url.Values {
	st := reflect.TypeOf(m).Elem()
	v := reflect.ValueOf(m).Elem()
	q := url.Values{}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i).Tag.Get("json")
		if !b.isExcludeField(field) {
			value := v.FieldByIndex([]int{i}).String()
			q.Add(field, value)
		}
	}
	b.encodepayload = q.Encode()
	return q
}

func (b *binanceSignature[T]) isExcludeField(field string) bool {
	return slices.Contains(b.exclude_fields, field)
}

func (b *binanceSignature[T]) SetExcludeFields(fields []string) {
	b.exclude_fields = fields
}

func (b *binanceSignature[T]) SetExcludeField(field string) {
	b.exclude_fields = append(b.exclude_fields, field)
}

func (b *binanceSignature[T]) GetQueryStringBinanceSignature() string {
	encodeData := fmt.Sprintf("%v&signature=%v", b.encodepayload, b.bnsign)
	return encodeData
}

func (b *binanceSignature[T]) GetSignature(key []byte) string {
	payload := []byte(b.encodepayload)
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)
	b.bnsign = hex.EncodeToString(mac.Sum(nil))
	return b.bnsign
}
