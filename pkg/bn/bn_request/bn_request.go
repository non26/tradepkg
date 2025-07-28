package binancerequest

import (
	"net/http"

	"github.com/non26/tradepkg/pkg/bn/sign"
)

type IBinanceServiceHttpRequest[T any] interface {
	NewBinanceHttpRequest(url string) error
	CreateRequestSignUrl(request *T, secretKey string) (string, error)
	RequestPostMethod()
	RequestGetMethod()
	AddBnDefaultFutureHeader(apiKey string)
	AddBnDefaultSpotHeader(apiKey string)
	AddHeader(key, value string)
	GetBinanceRequest() *http.Request
	SetQueryString(queryString string)
}

type IBnFutureServiceRequest interface {
	PrepareRequest() // validate field before send to binance ,for example, convert to upper/lower case
	GetData() interface{}
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
	bnsign := sign.NewSignHMACSHA256[T]("", secretKey)
	bnsign.Sign(request)
	signature, err := bnsign.GetQueryStringBinanceSignature(request)
	if err != nil {
		return "", err
	}
	return signature, nil
}

func (b *binanceServiceHttpRequest[T]) RequestPostMethod() {
	b.req.Method = http.MethodPost
}

func (b *binanceServiceHttpRequest[T]) RequestGetMethod() {
	b.req.Method = http.MethodGet
}

func (b *binanceServiceHttpRequest[T]) AddBnDefaultFutureHeader(apiKey string) {
	b.req.Header.Add("X-MBX-APIKEY", apiKey)
	b.req.Header.Add("CONTENT-TYPE", "application/x-www-form-urlencoded")
}

func (b *binanceServiceHttpRequest[T]) AddBnDefaultSpotHeader(apiKey string) {
	b.req.Header.Add("X-MBX-APIKEY", apiKey)
	b.req.Header.Add("CONTENT-TYPE", "application/x-www-form-urlencoded")
}

func (b *binanceServiceHttpRequest[T]) AddHeader(key, value string) {
	b.req.Header.Add(key, value)
}

func (b *binanceServiceHttpRequest[T]) GetBinanceRequest() *http.Request {
	return b.req
}

func (b *binanceServiceHttpRequest[T]) SetQueryString(queryString string) {
	if queryString != "" {
		b.req.URL.RawQuery = queryString
	}
}
