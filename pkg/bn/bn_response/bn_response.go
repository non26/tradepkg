package bnresponse

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
)

type IBinanceServiceHttpResponse[R any] interface {
	SetResponse(res *http.Response)
	DecodeBinanceServiceResponse() error
	GetBinanceServiceResponse() *R
	// TODO function to get error code
}

type binanceServiceHttpResponse[R any] struct {
	res   *http.Response
	bnres *R
}

func NewBinanceServiceHttpResponse[R any]() *binanceServiceHttpResponse[R] {
	b := binanceServiceHttpResponse[R]{}
	return &b
}

func (b *binanceServiceHttpResponse[R]) SetResponse(res *http.Response) {
	b.res = res
}

func (b *binanceServiceHttpResponse[R]) DecodeBinanceServiceResponse() error {
	// if b.res.StatusCode != http.StatusOK {
	// 	return b.binanceErrorResponse()
	// }
	// return b.binanceSuccessResponse()
	return b.binanceResponse()
}

func (b *binanceServiceHttpResponse[R]) GetBinanceServiceResponse() *R {
	return b.bnres
}

// func (b *binanceServiceHttpResponse[R]) binanceErrorResponse() error {
// 	defer b.res.Body.Close()
// 	bnResponseError := new(models.ResponseBinanceFutureError)
// 	json.NewDecoder(b.res.Body).Decode(bnResponseError)
// 	if bnResponseError.Code == -2013 {
// 		b.bnres = new(R)
// 		return nil
// 	}
// 	return utils.NewBinanceFail(
// 		b.res.StatusCode,
// 		bnResponseError.Code,
// 		bnResponseError.Message,
// 	)
// }

// func (b *binanceServiceHttpResponse[R]) binanceSuccessResponse() error {
func (b *binanceServiceHttpResponse[R]) binanceResponse() error {
	defer b.res.Body.Close()
	b.bnres = new(R)
	switch reflect.TypeOf(*b.bnres).Kind() {
	case reflect.Slice:
		bnResponse := new(R)
		bodyBytes, err := io.ReadAll(b.res.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bodyBytes, bnResponse)
		if err != nil {
			return err
		}
		b.bnres = bnResponse
	default:
		bnResponse := new(R)
		err := json.NewDecoder(b.res.Body).Decode(bnResponse)
		if err != nil {
			return err
		}
		b.bnres = new(R)
	}
	return nil
}
