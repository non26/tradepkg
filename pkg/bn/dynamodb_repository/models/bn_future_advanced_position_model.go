package dynamodbrepository

type BnFtAdvancedPositionModel struct {
	Symbol         string `dynamodbav:"symbol" dynamodb:"symbol"`       // primary key
	ClientId       string `dynamodbav:"client_id" dynamodb:"client_id"` // secondary key
	Side           string `dynamodbav:"side" dynamodb:"side"`
	PositionSide   string `dynamodbav:"position_side" dynamodb:"position_side"`
	AmountQ        string `dynamodbav:"amount_q" dynamodb:"amount_q"`
	AmountB        string `dynamodbav:"amount_b" dynamodb:"amount_b"`
	WatchingConfig string `dynamodbav:"watching_config" dynamodb:"watching_config"`
	CreatedAt      string `dynamodbav:"created_at" dynamodb:"created_at"`
}

func NewBinanceFutureAdvancedPosition() *BnFtAdvancedPositionModel {
	return &BnFtAdvancedPositionModel{}
}

func NewBinanceFutureAdvancedPositionWith(data BnFtAdvancedPositionModel) *BnFtAdvancedPositionModel {
	return &BnFtAdvancedPositionModel{
		Symbol:         data.Symbol,
		ClientId:       data.ClientId,
		Side:           data.Side,
		PositionSide:   data.PositionSide,
		AmountQ:        data.AmountQ,
		AmountB:        data.AmountB,
		WatchingConfig: data.WatchingConfig,
		CreatedAt:      data.CreatedAt,
	}
}

func (b *BnFtAdvancedPositionModel) IsFound() bool {
	return b.Symbol != ""
}
