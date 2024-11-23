package bn

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"reflect"
)

// type IBinanceSignature interface {
// 	Sign() error
// }

type BinanceSignature struct {
	apiKey    string
	secretKey string
}

func NewBinanceSignature(apiKey, secretKey string) *BinanceSignature {
	return &BinanceSignature{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func GetQueryStringFromStructType[T any](m *T) url.Values {
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

func CreateBinanceSignature(data *url.Values, binanceSecretKey string) string {
	payload := data.Encode()
	encodeString := Sign([]byte(payload), []byte(binanceSecretKey))
	encodeData := fmt.Sprintf("%v&signature=%v", data.Encode(), encodeString)
	return encodeData
}

func Sign(payload []byte, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}
