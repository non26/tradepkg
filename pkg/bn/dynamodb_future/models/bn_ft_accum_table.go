package dynamodbfuture

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtAccumulationTable struct {
	TableName string `table:"bn_future_accumulation"`
	*BnFtAccumulation
}

func NewBnFtAccumulationTable() *BnFtAccumulationTable {
	return &BnFtAccumulationTable{}
}

func NewBnFtAccumulationTableWith(accumulation *BnFtAccumulation) *BnFtAccumulationTable {
	return &BnFtAccumulationTable{
		BnFtAccumulation: accumulation,
	}
}

func (b *BnFtAccumulationTable) GetItem() *BnFtAccumulation {
	return b.BnFtAccumulation
}

func (b *BnFtAccumulationTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtAccumulationTable) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"accumulate_id": &types.AttributeValueMemberS{Value: b.AccumulationID},
	}
}

func (b *BnFtAccumulationTable) GetAccumulateIDTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AccumulationID", "dynamodb")
	return v
}

func (b *BnFtAccumulationTable) GetSymbolTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v
}

func (b *BnFtAccumulationTable) GetSideTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Side", "dynamodb")
	return v
}

func (b *BnFtAccumulationTable) GetPositionSideTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v
}

func (b *BnFtAccumulationTable) GetTotalAmountQTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "TotalAmountQ", "dynamodb")
	return v
}

func (b *BnFtAccumulationTable) Transform() {
	b.Symbol = strings.ToUpper(b.Symbol)
	b.Side = strings.ToUpper(b.Side)
	b.PositionSide = strings.ToUpper(b.PositionSide)
}
