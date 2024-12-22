package dynamodbrepository

type BnFtBotOnRun struct {
	BotID             string `dynamodbav:"bot_id" dynamodb:"bot_id"`
	BotOrderID        string `dynamodbav:"bot_order_id" dynamodb:"bot_order_id"`
	Symbol            string `dynamodbav:"symbol" dynamodb:"symbol"`
	PositionSide      string `dynamodbav:"position_side" dynamodb:"position_side"`
	PositionCondition string `dynamodbav:"position_condition" dynamodb:"position_condition"`
}

func NewBinanceFutureBotOnRun() *BnFtBotOnRun {
	return &BnFtBotOnRun{}
}

func (b *BnFtBotOnRun) IsFound() bool {
	return b.BotID != "" && b.BotOrderID != ""
}
