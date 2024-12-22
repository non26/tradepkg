package dynamodbrepository

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtBotOnRunTable struct {
	TableName string `table:"bot_on_run"`
	*BnFtBotOnRun
}

func NewBinanceFutureBotOnRunTable(botOnRun *BnFtBotOnRun) *BnFtBotOnRunTable {
	return &BnFtBotOnRunTable{
		BnFtBotOnRun: botOnRun,
	}
}

func (b *BnFtBotOnRunTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtBotOnRunTable) GetKeyBotIDAndOrderID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"bot_id":       &types.AttributeValueMemberS{Value: b.BotID},
		"bot_order_id": &types.AttributeValueMemberS{Value: b.BotOrderID},
	}
}

func (b *BnFtBotOnRunTable) GetKeyBotID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"bot_id": &types.AttributeValueMemberS{Value: b.BotID},
	}
}

func (b *BnFtBotOnRunTable) GetSymbolTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v
}

func (b *BnFtBotOnRunTable) GetPositionSideTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v
}

func (b *BnFtBotOnRunTable) GetPositionConditionTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionCondition", "dynamodb")
	return v
}
