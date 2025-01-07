package dynamodbrepository

import (
	"strconv"

	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtOpeningPosition struct {
	Symbol         string `dynamodbav:"symbol" dynamodb:"symbol"` // primary key
	ClientId       string `dynamodbav:"client_id" dynamodb:"client_id"`
	OrderType      string `dynamodbav:"order_type" dynamodb:"order_type"`
	Leverage       string `dynamodbav:"leverage" dynamodb:"leverage"`
	PositionSide   string `dynamodbav:"position_side" dynamodb:"position_side"`
	Side           string `dynamodbav:"side" dynamodb:"side"`
	AmountQ        string `dynamodbav:"amount_q" dynamodb:"amount_q"`
	AmountB        string `dynamodbav:"amount_b" dynamodb:"amount_b"`
	CreatedAt      string `dynamodbav:"created_at" dynamodb:"created_at"`
	WatchingConfig string `dynamodbav:"watching_config" dynamodb:"watching_config"`
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

func (b *BnFtOpeningPosition) SetCreatedAt() {
	b.CreatedAt = utils.GetDBTime()
}
