package dynamodbrepository

type BnFtBotOnRun struct {
	BotID             string `json:"bot_id"`
	BotOrderID        string `json:"bot_order_id"`
	Symbol            string `json:"symbol"`
	PositionSide      string `json:"position_side"`
	PositionCondition string `json:"position_condition"`
}

func NewBinanceFutureBotOnRun() *BnFtBotOnRun {
	return &BnFtBotOnRun{}
}

func (b *BnFtBotOnRun) IsFound() bool {
	return b.BotID != "" && b.BotOrderID != ""
}
