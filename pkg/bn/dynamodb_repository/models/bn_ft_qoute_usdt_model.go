package dynamodbrepository

import (
	"strconv"
)

type BnFtQouteUSDT struct {
	Symbol string `dynamodb:"symbol" dynamodbav:"symbol"` // primary key
	// CurrentLeverage int    `dynamodb:"current_leverage" dynamodbav:"current_leverage"`
	CountingLong  int `dynamodb:"counting_long" dynamodbav:"counting_long"`
	CountingShort int `dynamodb:"counting_short" dynamodbav:"counting_short"`
}

type counting int

func (c counting) Int() int {
	return int(c)
}

func (c counting) String() string {
	return strconv.Itoa(c.Int())
}

func NewBnFtQouteUSDT() *BnFtQouteUSDT {
	return &BnFtQouteUSDT{}
}

func (b *BnFtQouteUSDT) GetNextCountingLong() counting {
	return counting(b.CountingLong + 1)
}

func (b *BnFtQouteUSDT) GetNextCountingShort() counting {
	return counting(b.CountingShort + 1)
}

func (b *BnFtQouteUSDT) SetCountingLong(counting int) {
	b.CountingLong = counting
}

func (b *BnFtQouteUSDT) SetCountingShort(counting int) {
	b.CountingShort = counting
}

func (b *BnFtQouteUSDT) GetSymbol() string {
	return b.Symbol
}

func (b *BnFtQouteUSDT) GetCountingLong() int {
	return b.CountingLong
}

func (b *BnFtQouteUSDT) GetCountingShort() int {
	return b.CountingShort
}

func (b *BnFtQouteUSDT) SetSymbol(symbol string) {
	b.Symbol = symbol
}

func (b *BnFtQouteUSDT) IsFound() bool {
	return b.Symbol != ""
}
