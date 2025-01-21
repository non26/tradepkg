package dynamodbfuture

type BnFtAccumulation struct {
	AccumulateID   string `dynamodbav:"accumulate_id" dynamodb:"accumulate_id"`
	Symbol         string `dynamodbav:"symbol" dynamodb:"symbol"`
	Side           string `dynamodbav:"side" dynamodb:"side"`
	PositionSide   string `dynamodbav:"position_side" dynamodb:"position_side"`
	TotalAmountQ   string `dynamodbav:"total_amount_q" dynamodb:"total_amount_q"`
	CurrentAmountQ string `dynamodbav:"current_amount_q" dynamodb:"current_amount_q"`
}

func NewBnFtAccumulation() *BnFtAccumulation {
	return &BnFtAccumulation{}
}

func NewBnFtAccumulationWith(accumulation *BnFtAccumulation) *BnFtAccumulation {
	return &BnFtAccumulation{
		AccumulateID:   accumulation.AccumulateID,
		Symbol:         accumulation.Symbol,
		Side:           accumulation.Side,
		PositionSide:   accumulation.PositionSide,
		TotalAmountQ:   accumulation.TotalAmountQ,
		CurrentAmountQ: accumulation.CurrentAmountQ,
	}
}

func (b *BnFtAccumulation) GetAccumulateID() string {
	return b.AccumulateID
}

func (b *BnFtAccumulation) SetAccumulateID(accumulateID string) {
	b.AccumulateID = accumulateID
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
	return b.TotalAmountQ
}

func (b *BnFtAccumulation) SetTotalAmountQ(totalAmountQ string) {
	b.TotalAmountQ = totalAmountQ
}

func (b *BnFtAccumulation) GetCurrentAmountQ() string {
	return b.CurrentAmountQ
}

func (b *BnFtAccumulation) SetCurrentAmountQ(currentAmountQ string) {
	b.CurrentAmountQ = currentAmountQ
}
