package dynamodbfuture

type BnFtAdvancedPositionModel struct {
	Symbol         string `dynamodbav:"symbol" dynamodb:"symbol"`       // primary key
	ClientId       string `dynamodbav:"client_id" dynamodb:"client_id"` // secondary key
	Side           string `dynamodbav:"side" dynamodb:"side"`
	PositionSide   string `dynamodbav:"position_side" dynamodb:"position_side"`
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
		AmountB:        data.AmountB,
		WatchingConfig: data.WatchingConfig,
		CreatedAt:      data.CreatedAt,
	}
}

func (b *BnFtAdvancedPositionModel) IsFound() bool {
	return b.Symbol != ""
}

func (b *BnFtAdvancedPositionModel) GetSymbol() string {
	return b.Symbol
}

func (b *BnFtAdvancedPositionModel) SetSymbol(symbol string) {
	b.Symbol = symbol
}

func (b *BnFtAdvancedPositionModel) GetClientId() string {
	return b.ClientId
}

func (b *BnFtAdvancedPositionModel) SetClientId(clientId string) {
	b.ClientId = clientId
}

func (b *BnFtAdvancedPositionModel) GetSide() string {
	return b.Side
}

func (b *BnFtAdvancedPositionModel) SetSide(side string) {
	b.Side = side
}

func (b *BnFtAdvancedPositionModel) GetPositionSide() string {
	return b.PositionSide
}

func (b *BnFtAdvancedPositionModel) SetPositionSide(positionSide string) {
	b.PositionSide = positionSide
}

func (b *BnFtAdvancedPositionModel) GetAmountB() string {
	return b.AmountB
}

func (b *BnFtAdvancedPositionModel) SetAmountB(amountB string) {
	b.AmountB = amountB
}

func (b *BnFtAdvancedPositionModel) GetWatchingConfig() string {
	return b.WatchingConfig
}

func (b *BnFtAdvancedPositionModel) SetWatchingConfig(watchingConfig string) {
	b.WatchingConfig = watchingConfig
}

func (b *BnFtAdvancedPositionModel) GetCreatedAt() string {
	return b.CreatedAt
}

func (b *BnFtAdvancedPositionModel) SetCreatedAt(createdAt string) {
	b.CreatedAt = createdAt
}
