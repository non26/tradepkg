package dynamodbrepository

import (
	"strconv"
)

type BnFtOpeningPosition struct {
	Symbol             string `dynamodbav:"symbol" dynamodb:"symbol"` // primary key
	ClientId           string `dynamodbav:"client_id" dynamodb:"client_id"`
	OrderType          string `dynamodbav:"order_type" dynamodb:"order_type"`
	Leverage           string `dynamodbav:"leverage" dynamodb:"leverage"`
	PositionSide       string `dynamodbav:"position_side" dynamodb:"position_side"`
	Side               string `dynamodbav:"side" dynamodb:"side"`
	AmountQ            string `dynamodbav:"amount_q" dynamodb:"amount_q"`
	AmountB            string `dynamodbav:"amount_b" dynamodb:"amount_b"`
	BuyOrderCreatedAt  string `dynamodbav:"buy_created_at" dynamodb:"buy_created_at"`
	SellOrderCreatedAt string `dynamodbav:"sell_created_at" dynamodb:"sell_created_at"`
	WatchingConfig     string `dynamodbav:"watching_config" dynamodb:"watching_config"`
}

func NewBinanceFutureOpeningPosition() *BnFtOpeningPosition {
	return &BnFtOpeningPosition{
		// ExchangeId: 1,
	}
}

func (b *BnFtOpeningPosition) IsFound() bool {
	return b.Symbol != ""
}

func (b *BnFtOpeningPosition) AddMoreAmountQ(amountQ string) {
	amountQInt, _ := strconv.Atoi(amountQ)
	prevAmountQInt, _ := strconv.Atoi(b.AmountQ)
	b.AmountQ = strconv.Itoa(amountQInt + prevAmountQInt)
}

func (b *BnFtOpeningPosition) SetWatchingConfig(watchingConfig []byte) {
	b.WatchingConfig = string(watchingConfig)
}
