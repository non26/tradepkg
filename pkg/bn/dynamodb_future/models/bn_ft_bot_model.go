package dynamodbfuture

type BnFtBot struct {
	BotID   string `dynamodbav:"bot_id" dynamodb:"bot_id"`
	BotName string `dynamodbav:"bot_name" dynamodb:"bot_name"`
}

func NewBnFtBot() *BnFtBot {
	return &BnFtBot{}
}

func (b *BnFtBot) IsFound() bool {
	return b.BotID != ""
}
