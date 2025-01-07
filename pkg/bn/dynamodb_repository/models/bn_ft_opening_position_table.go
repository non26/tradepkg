package dynamodbrepository

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtOpeningPositionTable struct {
	TableName string `table:"bn_future_opening_position"`
	*BnFtOpeningPosition
}

func NewBinanceFutureOpeningPositionTable() *BnFtOpeningPositionTable {
	return &BnFtOpeningPositionTable{
		BnFtOpeningPosition: NewBinanceFutureOpeningPosition(),
	}
}

func (b *BnFtOpeningPositionTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtOpeningPositionTable) SetCreatedAt() {
	b.CreatedAt = utils.GetDBTime()
}

func (b *BnFtOpeningPositionTable) GetLeverageTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Leverage", "dynamodb")
	return v
}

func (b *BnFtOpeningPositionTable) GetOrderTypeTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "OrderType", "dynamodb")
	return v
}

func (b *BnFtOpeningPositionTable) GetClientIdTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ClientId", "dynamodb")
	return v
}

func (b *BnFtOpeningPositionTable) GetAmountQTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountQ", "dynamodb")
	return v
}

func (b *BnFtOpeningPositionTable) GetAmountBTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountB", "dynamodb")
	return v
}

func (b *BnFtOpeningPositionTable) GetWatchingConfigTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "WatchingConfig", "dynamodb")
	return v
}

func (b *BnFtOpeningPositionTable) GetKeyBySymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol": &types.AttributeValueMemberS{Value: b.Symbol},
	}
}

func (b *BnFtOpeningPositionTable) GetKeyByClientID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientId},
	}
}

func (b *BnFtOpeningPositionTable) GetKeyByPositionSide() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"position_side": &types.AttributeValueMemberS{Value: b.PositionSide},
	}
}

func (b *BnFtOpeningPositionTable) GetKeyByPositionSideAndSymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"position_side": &types.AttributeValueMemberS{Value: b.PositionSide},
		"symbol":        &types.AttributeValueMemberS{Value: b.Symbol},
	}
}
