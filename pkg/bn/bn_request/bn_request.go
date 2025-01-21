package binancerequest

import (
	"io"
	"net/http"
	"strings"

	"github.com/non26/tradepkg/pkg/bn/sign"
)

type IBinanceServiceHttpRequest[T any] interface {
	NewBinanceHttpRequest(url string) error
	CreateRequestSignUrl(request *T, secretKey string) (string, error)
	RequestPostMethod(signature string)
	RequestGetMethod(signature string)
	AddHeader(apiKey string)
	GetBinanceRequest() *http.Request
}

/*
T is type of request
*/
type binanceServiceHttpRequest[T any] struct {
	req *http.Request
}

func NewBinanceServiceHttpRequest[T any]() IBinanceServiceHttpRequest[T] {
	return &binanceServiceHttpRequest[T]{}
}

func (b *binanceServiceHttpRequest[T]) NewBinanceHttpRequest(
	url string,
) error {
	req, err := http.NewRequest(
		"", url, nil,
	)
	if err != nil {
		return err
	}
	b.req = req
	return nil
}

func (b *binanceServiceHttpRequest[T]) CreateRequestSignUrl(request *T, secretKey string) (string, error) {
	// data := GetQueryStringFromStructType(request)
	// sig := CreateBinanceSignature(&data, secretKey)
	bnsign := sign.NewSignHMACSHA256[T]("", secretKey)
	signature, err := bnsign.GetQueryStringBinanceSignature(request)
	if err != nil {
		return "", err
	}
	return signature, nil
}

func (b *binanceServiceHttpRequest[T]) RequestPostMethod(signature string) {
	b.req.Method = http.MethodPost
	if signature != "" {
		var body io.Reader = strings.NewReader(signature)
		rc, ok := body.(io.ReadCloser)
		if !ok {
			rc = io.NopCloser(body)
		}
		b.req.Body = rc
	}
}

func (b *binanceServiceHttpRequest[T]) RequestGetMethod(signature string) {
	b.req.Method = http.MethodGet
	if signature != "" {
		b.req.URL.RawQuery = signature
	}
}

func (b *binanceServiceHttpRequest[T]) AddHeader(apiKey string) {
	b.req.Header.Add("X-MBX-APIKEY", apiKey)
	b.req.Header.Add("CONTENT-TYPE", "application/x-www-form-urlencoded")
}

func (b *binanceServiceHttpRequest[T]) GetBinanceRequest() *http.Request {
	return b.req
}
