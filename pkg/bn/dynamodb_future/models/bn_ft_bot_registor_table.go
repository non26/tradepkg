package dynamodbfuture

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtBotRegistorTable struct {
	TableName string `table:"bot_registor"`
	*BnFtBotRegistor
}

func NewBnFtBotRegistorTable(botRegistor *BnFtBotRegistor) *BnFtBotRegistorTable {
	return &BnFtBotRegistorTable{
		BnFtBotRegistor: botRegistor,
	}
}

func (b *BnFtBotRegistorTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtBotRegistorTable) GetKeyBotIDAndOrderID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"bot_id":       &types.AttributeValueMemberS{Value: b.BotID},
		"bot_order_id": &types.AttributeValueMemberS{Value: b.BotOrderID},
	}
}

func (b *BnFtBotRegistorTable) GetBotIDTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "BotID", "dynamodb")
	return v, t
}

func (b *BnFtBotRegistorTable) GetBotOrderIDTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "BotOrderID", "dynamodb")
	return v, t
}

func (b *BnFtBotRegistorTable) GetIsActiveTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "IsActive", "dynamodb")
	return v, t
}

func (b *BnFtBotRegistorTable) GetSymbolTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v, t
}

func (b *BnFtBotRegistorTable) GetPositionSideTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v, t
}

func (b *BnFtBotRegistorTable) GetAmountQtyTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountQ", "dynamodb")
	return v, t
}

func (b *BnFtBotRegistorTable) GetAccountIdTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AccountId", "dynamodb")
	return v, t
}

func (b *BnFtBotRegistorTable) GetSettingTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Setting", "dynamodb")
	return v, t
}
