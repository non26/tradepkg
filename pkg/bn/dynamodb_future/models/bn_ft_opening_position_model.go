package dynamodbfuture

import (
	"strconv"
)

type BnFtOpeningPosition struct {
	Symbol         string `dynamodbav:"symbol" dynamodb:"symbol"`               // primary key
	PositionSide   string `dynamodbav:"position_side" dynamodb:"position_side"` // second index
	ClientId       string `dynamodbav:"client_id" dynamodb:"client_id"`
	OrderType      string `dynamodbav:"order_type" dynamodb:"order_type"`
	Leverage       string `dynamodbav:"leverage" dynamodb:"leverage"`
	Side           string `dynamodbav:"side" dynamodb:"side"`
	AmountB        string `dynamodbav:"amount_b" dynamodb:"amount_b"`
	CreatedAt      string `dynamodbav:"created_at" dynamodb:"created_at"`
	WatchingConfig string `dynamodbav:"watching_config" dynamodb:"watching_config"`
}

func NewBinanceFutureOpeningPosition() *BnFtOpeningPosition {
	return &BnFtOpeningPosition{}
}

func (b *BnFtOpeningPosition) IsFound() bool {
	return b.Symbol != ""
}

func (b *BnFtOpeningPosition) AddMoreAmountB(amountB string) {
	amountQInt, _ := strconv.Atoi(amountB)
	prevAmountQInt, _ := strconv.Atoi(b.AmountB)
	b.AmountB = strconv.Itoa(amountQInt + prevAmountQInt)
}

func (b *BnFtOpeningPosition) SetWatchingConfig(watchingConfig []byte) {
	b.WatchingConfig = string(watchingConfig)
}
