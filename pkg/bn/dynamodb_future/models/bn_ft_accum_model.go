package dynamodbfuture

type BnFtAccumulation struct {
	AccumulationID string `dynamodbav:"accumulation_id" dynamodb:"accumulation_id"`
	Symbol         string `dynamodbav:"symbol" dynamodb:"symbol"`
	Side           string `dynamodbav:"side" dynamodb:"side"`
	PositionSide   string `dynamodbav:"position_side" dynamodb:"position_side"`
	TotalAmountB   string `dynamodbav:"total_amount_b" dynamodb:"total_amount_b"`
	CurrentAmountB string `dynamodbav:"current_amount_b" dynamodb:"current_amount_b"`
}

func NewBnFtAccumulation() *BnFtAccumulation {
	return &BnFtAccumulation{}
}

func NewBnFtAccumulationWith(accumulation *BnFtAccumulation) *BnFtAccumulation {
	return &BnFtAccumulation{
		AccumulationID: accumulation.AccumulationID,
		Symbol:         accumulation.Symbol,
		Side:           accumulation.Side,
		PositionSide:   accumulation.PositionSide,
		TotalAmountB:   accumulation.TotalAmountB,
		CurrentAmountB: accumulation.CurrentAmountB,
	}
}

func (b *BnFtAccumulation) GetAccumulateID() string {
	return b.AccumulationID
}

func (b *BnFtAccumulation) SetAccumulateID(accumulateID string) {
	b.AccumulationID = accumulateID
}

func (b *BnFtAccumulation) GetSymbol() string {
	return b.Symbol
}

func (b *BnFtAccumulation) SetSymbol(symbol string) {
	b.Symbol = symbol
}

func (b *BnFtAccumulation) GetSide() string {
	return b.Side
}

func (b *BnFtAccumulation) SetSide(side string) {
	b.Side = side
}

func (b *BnFtAccumulation) GetPositionSide() string {
	return b.PositionSide
}

func (b *BnFtAccumulation) SetPositionSide(positionSide string) {
	b.PositionSide = positionSide
}

func (b *BnFtAccumulation) GetTotalAmountQ() string {
	return b.TotalAmountB
}

func (b *BnFtAccumulation) SetTotalAmountQ(totalAmountQ string) {
	b.TotalAmountB = totalAmountQ
}

func (b *BnFtAccumulation) GetCurrentAmountQ() string {
	return b.CurrentAmountB
}

func (b *BnFtAccumulation) SetCurrentAmountQ(currentAmountQ string) {
	b.CurrentAmountB = currentAmountQ
}
