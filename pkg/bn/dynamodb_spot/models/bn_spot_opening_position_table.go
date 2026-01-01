package dynamodbspot

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	bntime "github.com/non26/tradepkg/pkg/bn/bn_time"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnSpotOpeningPositionTable struct {
	TableName string `table:"bn_spot_opening_position"`
	*BnSpotOpeningPosition
}

func NewBinanceSpotOpeningPositionTable() *BnSpotOpeningPositionTable {
	return &BnSpotOpeningPositionTable{
		BnSpotOpeningPosition: NewBinanceSpotOpeningPosition(),
	}
}

func NewBinanceSpotOpeningPositionTableWith(openingPosition *BnSpotOpeningPosition) *BnSpotOpeningPositionTable {
	return &BnSpotOpeningPositionTable{
		BnSpotOpeningPosition: openingPosition,
	}
}

func (b *BnSpotOpeningPositionTable) GetData() *BnSpotOpeningPosition {
	return b.BnSpotOpeningPosition
}

func (b *BnSpotOpeningPositionTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnSpotOpeningPositionTable) SetCreatedAt() {
	b.CreatedAt = bntime.GetDBTime()
}

func (b *BnSpotOpeningPositionTable) GetSymbolTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v, t
}

func (b *BnSpotOpeningPositionTable) GetPositionSideTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v, t
}

func (b *BnSpotOpeningPositionTable) GetSideTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Side", "dynamodb")
	return v, t
}

func (b *BnSpotOpeningPositionTable) GetCreatedAtTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CreatedAt", "dynamodb")
	return v, t
}

func (b *BnSpotOpeningPositionTable) GetOrderTypeTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "OrderType", "dynamodb")
	return v, t
}

func (b *BnSpotOpeningPositionTable) GetClientIdTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ClientId", "dynamodb")
	return v, t
}

func (b *BnSpotOpeningPositionTable) GetAmountBTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountB", "dynamodb")
	return v, t
}
func (b *BnSpotOpeningPositionTable) GetKeyBySymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol": &types.AttributeValueMemberS{Value: b.Symbol},
	}
}

func (b *BnSpotOpeningPositionTable) GetKeyByClientID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientId},
	}
}

// func (b *BnSpotOpeningPositionTable) GetKeyByPositionSide() map[string]types.AttributeValue {
// 	return map[string]types.AttributeValue{
// 		"position_side": &types.AttributeValueMemberS{Value: b.PositionSide},
// 	}
// }

// func (b *BnSpotOpeningPositionTable) GetKeyByPositionSideAndSymbol() map[string]types.AttributeValue {
// 	return map[string]types.AttributeValue{
// 		"position_side": &types.AttributeValueMemberS{Value: b.PositionSide},
// 		"symbol":        &types.AttributeValueMemberS{Value: b.Symbol},
// 	}
// }

func (b *BnSpotOpeningPositionTable) Transform() {
	b.Symbol = strings.ToUpper(b.Symbol)
	// b.PositionSide = strings.ToUpper(b.PositionSide)
	// b.Side = strings.ToUpper(b.Side)
	// b.OrderType = strings.ToUpper(b.OrderType)
}
