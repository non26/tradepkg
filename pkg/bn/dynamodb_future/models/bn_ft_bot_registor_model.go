package dynamodbfuture

type BnFtBotRegistor struct {
	BotID        string `dynamodbav:"bot_id" dynamodb:"bot_id"`
	BotOrderID   string `dynamodbav:"bot_order_id" dynamodb:"bot_order_id"`
	IsActive     bool   `dynamodbav:"is_active" dynamodb:"is_active"`
	Symbol       string `dynamodbav:"symbol" dynamodb:"symbol"`
	PositionSide string `dynamodbav:"position_side" dynamodb:"position_side"`
	AmountQ      string `dynamodbav:"amount_q" dynamodb:"amount_q"`
	AccountId    string `dynamodbav:"account_id" dynamodb:"account_id"`
	Setting      string `dynamodbav:"setting" dynamodb:"setting"`
}

func NewBnFtBotRegistor() *BnFtBotRegistor {
	return &BnFtBotRegistor{}
}

func (b *BnFtBotRegistor) IsFound() bool {
	return b.BotID != "" && b.BotOrderID != ""
}

func (b *BnFtBotRegistor) SetSetting(setting []byte) {
	b.Setting = string(setting)
}
