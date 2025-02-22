package dynamodbspot

type BnSpotHistory struct {
	ClientId  string `dynamodbav:"client_id" dynamodb:"client_id"` // primary key
	Symbol    string `dynamodbav:"symbol" dynamodb:"symbol"`
	CreatedAt string `dynamodbav:"created_at" dynamodb:"created_at"`
}

func NewBinanceSpotHistory() *BnSpotHistory {
	return &BnSpotHistory{}
}

func (b *BnSpotHistory) IsFound() bool {
	return b.ClientId != ""
}
