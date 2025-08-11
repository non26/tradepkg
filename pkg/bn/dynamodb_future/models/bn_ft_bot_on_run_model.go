package dynamodbfuture

type BnFtBotOnRun struct {
	BotID        string `dynamodbav:"bot_id" dynamodb:"bot_id"`
	BotOrderID   string `dynamodbav:"bot_order_id" dynamodb:"bot_order_id"`
	Symbol       string `dynamodbav:"symbol" dynamodb:"symbol"`
	PositionSide string `dynamodbav:"position_side" dynamodb:"position_side"`
	// PositionCondition string `dynamodbav:"position_condition" dynamodb:"position_condition"`
	AmountB   string `dynamodbav:"amount_b" dynamodb:"amount_b"`
	IsActive  bool   `dynamodbav:"is_active" dynamodb:"is_active"`
	AccountId string `dynamodbav:"account_id" dynamodb:"account_id"`
	Setting   string `dynamodbav:"setting" dynamodb:"setting"`
}

func NewBinanceFutureBotOnRun() *BnFtBotOnRun {
	return &BnFtBotOnRun{}
}

func (b *BnFtBotOnRun) IsFound() bool {
	return b.BotID != "" && b.BotOrderID != ""
}

func (b *BnFtBotOnRun) SetSetting(setting []byte) {
	b.Setting = string(setting)
}
