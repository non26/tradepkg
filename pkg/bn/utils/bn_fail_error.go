package utils

import (
	"fmt"
)

type BinanceFail struct {
	HttpStatus int
	BnCode     int
	BnMessage  string
}

// httpStatus = http status code from binance response
// serviceNameCode = service name code from binance response
// serviceNameMessage = service name message from binance response
func NewBinanceFail(httpStatus int, bnCode int, bnMessage string) *BinanceFail {
	return &BinanceFail{
		HttpStatus: httpStatus,
		BnCode:     bnCode,
		BnMessage:  bnMessage,
	}
}

func (b *BinanceFail) Error() string {
	return fmt.Sprintf("BnHttp:%v,BnCode:%v,BnMsg:%v", b.HttpStatus, b.BnCode, b.BnMessage)
}
