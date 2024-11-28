package sign

type IBinanceSignature[T any] interface {
	Sign(m *T, except_fields ...string) (string, error)
	GetQueryStringBinanceSignature(m *T, except_fields ...string) (string, error)
}

// type IBinanceSignature[T any] interface {
// 	GetSignature() string
// 	GetWSSignature() string
// 	GetEncodePayload() string
// 	GetURLValue() url.Values
// 	GetQueryStringBinanceSignature() string
// 	CreateQueryStringFromPayload(m *T) *binanceSignature[T]
// 	SetExcludeFields(fields []string)
// 	SetExcludeField(field string)
// }

// type binanceSignature[T any] struct {
// 	apiKey         string
// 	secretKey      string
// 	bnsign         string
// 	exclude_fields []string
// 	encodepayload  string
// 	urlValue       url.Values
// }

// func NewBinanceNonWsSignature[T any](apiKey, secretKey string) IBinanceSignature[T] {
// 	return &binanceSignature[T]{
// 		apiKey:         apiKey,
// 		secretKey:      secretKey,
// 		exclude_fields: []string{},
// 	}
// }

// func (b *binanceSignature[T]) CreateQueryStringFromPayload(m *T) *binanceSignature[T] {
// 	st := reflect.TypeOf(m).Elem()
// 	v := reflect.ValueOf(m).Elem()
// 	q := url.Values{}
// 	for i := 0; i < st.NumField(); i++ {
// 		field := st.Field(i).Tag.Get("json")
// 		if !b.isExcludeField(field) {
// 			value := v.FieldByIndex([]int{i})
// 			if value.Kind() == reflect.String {
// 				q.Add(field, value.String())
// 				continue
// 			}
// 			if value.Kind() == reflect.Int64 {
// 				q.Add(field, fmt.Sprintf("%d", value.Int()))
// 				continue
// 			}
// 		}
// 	}
// 	b.urlValue = q
// 	b.encodepayload = q.Encode()
// 	return b
// }

// func (b *binanceSignature[T]) isExcludeField(field string) bool {
// 	return slices.Contains(b.exclude_fields, field)
// }

// func (b *binanceSignature[T]) SetExcludeFields(fields []string) {
// 	b.exclude_fields = fields
// }

// func (b *binanceSignature[T]) SetExcludeField(field string) {
// 	b.exclude_fields = append(b.exclude_fields, field)
// }

// func (b *binanceSignature[T]) GetQueryStringBinanceSignature() string {
// 	encodeData := fmt.Sprintf("%v&signature=%v", b.encodepayload, b.bnsign)
// 	return encodeData
// }

// func (b *binanceSignature[T]) GetSignature() string {
// 	payload := []byte(b.encodepayload)
// 	mac := hmac.New(sha256.New, []byte(b.secretKey))
// 	mac.Write(payload)
// 	b.bnsign = hex.EncodeToString(mac.Sum(nil))
// 	return b.bnsign
// }

// func (b *binanceSignature[T]) GetWSSignature() string {
// 	privateKey := ed25519.PrivateKey(b.secretKey)
// 	signature := ed25519.Sign(privateKey, []byte(b.encodepayload))
// 	// base64Signature := base64.StdEncoding.EncodeToString(signature)
// 	return string(signature)
// }

// func (b *binanceSignature[T]) GetEncodePayload() string {
// 	return b.encodepayload
// }
// func (b *binanceSignature[T]) GetURLValue() url.Values {
// 	return b.urlValue
// }

// func Ed25519(secretKey string, data string) (*string, error) {
// 	block, _ := pem.Decode([]byte(secretKey))
// 	if block == nil {
// 		return nil, fmt.Errorf("Ed25519 pem.Decode failed, invalid pem format secretKey")
// 	}
// 	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
// 	if err != nil {
// 		return nil, fmt.Errorf("Ed25519 call ParsePKCS8PrivateKey failed, error=%v", err.Error())
// 	}
// 	ed25519PrivateKey, ok := privateKey.(ed25519.PrivateKey)
// 	if !ok {
// 		return nil, fmt.Errorf("Ed25519 convert PrivateKey failed")
// 	}
// 	pk := ed25519.PrivateKey(ed25519PrivateKey)
// 	signature := ed25519.Sign(pk, []byte(data))
// 	encodedSignature := base64.StdEncoding.EncodeToString(signature)
// 	return &encodedSignature, nil
// }
