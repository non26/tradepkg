package dynamodbrepository

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtQouteUSTDTable struct {
	TableName string `table:"bn_future_qoute_usdt"`
	*BnFtQouteUSDT
}

func NewBinanceFutureQouteUSTDTable() *BnFtQouteUSTDTable {
	return &BnFtQouteUSTDTable{
		BnFtQouteUSDT: newBnFtQouteUSDT(),
	}
}

func (b *BnFtQouteUSTDTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtQouteUSTDTable) GetKeyBySymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol": &types.AttributeValueMemberS{Value: b.Symbol},
	}
}

func (b *BnFtQouteUSTDTable) GetSymbolTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v
}

func (b *BnFtQouteUSTDTable) GetCountingLongTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CountingLong", "dynamodb")
	return v
}

func (b *BnFtQouteUSTDTable) GetCountingShortTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CountingShort", "dynamodb")
	return v
}

func (b *BnFtQouteUSTDTable) GetCurrentLeverageTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CurrentLeverage", "dynamodb")
	return v
}
