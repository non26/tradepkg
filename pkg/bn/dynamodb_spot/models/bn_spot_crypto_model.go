package dynamodbspot

import (
	"strconv"
)

type BnSpotCrypto struct {
	Symbol   string `dynamodb:"symbol" dynamodbav:"symbol"` // primary key
	Counting int64  `dynamodb:"counting" dynamodbav:"counting"`
}

type counting int64

func (c counting) Int() int64 {
	return int64(c)
}

func (c counting) String() string {
	return strconv.FormatInt(c.Int(), 10)
}

func NewBnSpotCrypto() *BnSpotCrypto {
	return &BnSpotCrypto{}
}

func (b *BnSpotCrypto) GetNextCounting() counting {
	return counting(b.Counting + 1)
}

func (b *BnSpotCrypto) SetCounting(counting int64) {
	b.Counting = counting
}

func (b *BnSpotCrypto) GetSymbol() string {
	return b.Symbol
}

func (b *BnSpotCrypto) GetCounting() int64 {
	return b.Counting
}

func (b *BnSpotCrypto) SetSymbol(symbol string) {
	b.Symbol = symbol
}

func (b *BnSpotCrypto) IsFound() bool {
	return b.Symbol != ""
}
