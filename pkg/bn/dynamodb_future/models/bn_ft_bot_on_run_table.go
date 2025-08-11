package dynamodbfuture

import (
	"reflect"
	"strings"

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

func NewBinanceFutureBotOnRunTableWith(botOnRun *BnFtBotOnRun) *BnFtBotOnRunTable {
	return &BnFtBotOnRunTable{
		BnFtBotOnRun: botOnRun,
	}
}

func (b *BnFtBotOnRunTable) GetData() *BnFtBotOnRun {
	return b.BnFtBotOnRun
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

func (b *BnFtBotOnRunTable) GetBotIDTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "BotID", "dynamodb")
	return v, t
}

func (b *BnFtBotOnRunTable) GetBotOrderIDTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "BotOrderID", "dynamodb")
	return v, t
}

func (b *BnFtBotOnRunTable) GetSymbolTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v, t
}

func (b *BnFtBotOnRunTable) GetPositionSideTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v, t
}

// func (b *BnFtBotOnRunTable) GetPositionConditionTableField() (string, reflect.Type) {
// 	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionCondition", "dynamodb")
// 	return v, t
// }

func (b *BnFtBotOnRunTable) GetAmountQtyTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountB", "dynamodb")
	return v, t
}

func (b *BnFtBotOnRunTable) GetIsActiveTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "IsActive", "dynamodb")
	return v, t
}

func (b *BnFtBotOnRunTable) GetAccountIdTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AccountId", "dynamodb")
	return v, t
}

func (b *BnFtBotOnRunTable) GetSettingTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Setting", "dynamodb")
	return v, t
}

func (b *BnFtBotOnRunTable) Transform() {
	b.Symbol = strings.ToUpper(b.Symbol)
	b.PositionSide = strings.ToUpper(b.PositionSide)
}
