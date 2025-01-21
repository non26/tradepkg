package dynamodbfuture

type BnFtHistory struct {
	ClientId     string `dynamodbav:"client_id" dynamodb:"client_id"`
	Symbol       string `dynamodbav:"symbol" dynamodb:"symbol"`
	PositionSide string `dynamodbav:"position_side" dynamodb:"position_side"`
	CreatedAt    string `dynamodbav:"created_at" dynamodb:"created_at"`
}

func NewBinanceFutureHistory() *BnFtHistory {
	return &BnFtHistory{}
}

func (b *BnFtHistory) IsFound() bool {
	return b.Symbol != ""
}
