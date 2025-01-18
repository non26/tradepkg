package dynamodbrepository

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtAdvancedPositionTable struct {
	TableName string `table:"bn_future_advanced_position"`
	*BnFtAdvancedPositionModel
}

func NewBinanceFutureAdvancedPositionTable() *BnFtAdvancedPositionTable {
	return &BnFtAdvancedPositionTable{
		BnFtAdvancedPositionModel: NewBinanceFutureAdvancedPosition(),
	}
}

func NewBinanceFutureAdvancedPositionTableWith(data *BnFtAdvancedPositionModel) *BnFtAdvancedPositionTable {
	return &BnFtAdvancedPositionTable{
		BnFtAdvancedPositionModel: data,
	}
}

func (b *BnFtAdvancedPositionTable) GetData() *BnFtAdvancedPositionModel {
	return b.BnFtAdvancedPositionModel
}

func (b *BnFtAdvancedPositionTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtAdvancedPositionTable) SetCreatedAt() {
	b.CreatedAt = utils.GetDBTime()
}

func (b *BnFtAdvancedPositionTable) GetSymbolTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v
}

func (b *BnFtAdvancedPositionTable) GetClientIdTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ClientId", "dynamodb")
	return v
}

func (b *BnFtAdvancedPositionTable) GetSideTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Side", "dynamodb")
	return v
}

func (b *BnFtAdvancedPositionTable) GetPositionSideTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v
}

func (b *BnFtAdvancedPositionTable) GetAmountQTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountQ", "dynamodb")
	return v
}

func (b *BnFtAdvancedPositionTable) GetAmountBTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountB", "dynamodb")
	return v
}

func (b *BnFtAdvancedPositionTable) GetWatchingConfigTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "WatchingConfig", "dynamodb")
	return v
}

func (b *BnFtAdvancedPositionTable) GetCreatedAtTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CreatedAt", "dynamodb")
	return v
}

func (b *BnFtAdvancedPositionTable) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol":        &types.AttributeValueMemberS{Value: b.Symbol},
		"position_side": &types.AttributeValueMemberS{Value: b.PositionSide},
	}
}

func (b *BnFtAdvancedPositionTable) GetKeyByClientId() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientId},
	}
}

func (b *BnFtAdvancedPositionTable) Transform() {
	b.Symbol = strings.ToUpper(b.Symbol)
	b.PositionSide = strings.ToUpper(b.PositionSide)
	b.Side = strings.ToUpper(b.Side)
}
