package dynamodbfuture

type BnFtAdvancedPosition struct {
	ClientID     string `dynamodbav:"client_id" dynamodb:"client_id"`
	Symbol       string `dynamodbav:"symbol" dynamodb:"symbol"`
	PositionSide string `dynamodbav:"position_side" dynamodb:"position_side"`
	Side         string `dynamodbav:"side" dynamodb:"side"`
	AmountB      string `dynamodbav:"amount_b" dynamodb:"amount_b"`
}

func NewBnFtAdvancedPosition() *BnFtAdvancedPosition {
	return &BnFtAdvancedPosition{}
}

func (b *BnFtAdvancedPosition) IsFound() bool {
	return b.ClientID != ""
}
