package dynamodbspot

type BnSpotOpeningPosition struct {
	Symbol string `dynamodbav:"symbol" dynamodb:"symbol"` // primary key
	// PositionSide string `dynamodbav:"position_side" dynamodb:"position_side"` // second index
	ClientId string `dynamodbav:"client_id" dynamodb:"client_id"`
	// OrderType    string `dynamodbav:"order_type" dynamodb:"order_type"`
	// Side         string `dynamodbav:"side" dynamodb:"side"`
	AmountB   string `dynamodbav:"amount_b" dynamodb:"amount_b"`
	CreatedAt string `dynamodbav:"created_at" dynamodb:"created_at"`
}

func NewBinanceSpotOpeningPosition() *BnSpotOpeningPosition {
	return &BnSpotOpeningPosition{}
}

func (b *BnSpotOpeningPosition) IsFound() bool {
	return b.Symbol != ""
}
