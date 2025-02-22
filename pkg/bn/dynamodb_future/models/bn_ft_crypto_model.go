package dynamodbfuture

import (
	"strconv"
)

type BnFtCrypto struct {
	Symbol        string `dynamodb:"symbol" dynamodbav:"symbol"` // primary key
	CountingLong  int64  `dynamodb:"counting_long" dynamodbav:"counting_long"`
	CountingShort int64  `dynamodb:"counting_short" dynamodbav:"counting_short"`
}

type counting int64

func (c counting) Int() int64 {
	return int64(c)
}

func (c counting) String() string {
	return strconv.Itoa(int(c.Int()))
}

func NewBnFtCrypto() *BnFtCrypto {
	return &BnFtCrypto{}
}

func (b *BnFtCrypto) GetNextCountingLong() counting {
	return counting(b.CountingLong + 1)
}

func (b *BnFtCrypto) GetNextCountingShort() counting {
	return counting(b.CountingShort + 1)
}

func (b *BnFtCrypto) SetCountingLong(counting int64) {
	b.CountingLong = counting
}

func (b *BnFtCrypto) SetCountingShort(counting int64) {
	b.CountingShort = counting
}

func (b *BnFtCrypto) GetSymbol() string {
	return b.Symbol
}

func (b *BnFtCrypto) GetCountingLong() int64 {
	return b.CountingLong
}

func (b *BnFtCrypto) GetCountingShort() int64 {
	return b.CountingShort
}

func (b *BnFtCrypto) SetSymbol(symbol string) {
	b.Symbol = symbol
}

func (b *BnFtCrypto) IsFound() bool {
	return b.Symbol != ""
}
