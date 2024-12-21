package dynamodbrepository

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtHistoryTable struct {
	TableName string `table:"bn_future_history"`
	*BnFtHistory
}

func NewBinanceFutureHistoryTable() *BnFtHistoryTable {
	return &BnFtHistoryTable{
		BnFtHistory: NewBinanceFutureHistory(),
	}
}

func (b *BnFtHistoryTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtHistoryTable) GetHistoryIdTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "HistoryId", "dynamodb")
	return v
}

func (b *BnFtHistoryTable) GetSymbolTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v
}

func (b *BnFtHistoryTable) GetPositionSideTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v
}

func (b *BnFtHistoryTable) GetKeyClientID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientId},
	}
}
