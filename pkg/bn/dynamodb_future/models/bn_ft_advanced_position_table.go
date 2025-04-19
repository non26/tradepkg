package dynamodbfuture

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtAdvancedPositionTable struct {
	TableName string `table:"bn_future_advanced_position"`
	*BnFtAdvancedPosition
}

func NewBnFtAdvancedPositionTable() *BnFtAdvancedPositionTable {
	return &BnFtAdvancedPositionTable{
		BnFtAdvancedPosition: NewBnFtAdvancedPosition(),
	}
}

func NewBnFtAdvancedPositionTableWith(advancedPosition *BnFtAdvancedPosition) *BnFtAdvancedPositionTable {
	return &BnFtAdvancedPositionTable{
		BnFtAdvancedPosition: advancedPosition,
	}
}

func (b *BnFtAdvancedPositionTable) GetData() *BnFtAdvancedPosition {
	return b.BnFtAdvancedPosition
}

func (b *BnFtAdvancedPositionTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtAdvancedPositionTable) GetClientIDTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ClientID", "dynamodb")
	return v, t
}

func (b *BnFtAdvancedPositionTable) GetSymbolTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v, t
}

func (b *BnFtAdvancedPositionTable) GetPositionSideTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v, t
}

func (b *BnFtAdvancedPositionTable) GetSideTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Side", "dynamodb")
	return v, t
}

func (b *BnFtAdvancedPositionTable) GetAmountBTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountB", "dynamodb")
	return v, t
}

func (b *BnFtAdvancedPositionTable) Transform() {
	b.Symbol = strings.ToUpper(b.Symbol)
	b.PositionSide = strings.ToUpper(b.PositionSide)
	b.Side = strings.ToUpper(b.Side)
}

func (b *BnFtAdvancedPositionTable) GetKeyClientID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientID},
	}
}
