package dynamodbrepository

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtBotTable struct {
	TableName string `table:"bot"`
	*BnFtBot
}

func NewBnFtBotTable() *BnFtBotTable {
	return &BnFtBotTable{
		BnFtBot: NewBnFtBot(),
	}
}

func (b *BnFtBotTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtBotTable) GetKeyByBotID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"bot_id": &types.AttributeValueMemberS{Value: b.BotID},
	}
}
