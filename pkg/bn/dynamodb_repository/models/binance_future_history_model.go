package dynamodbrepository

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BinanceFutureHistory struct {
	ClientId     string `dynamodbav:"client_id" dynamodb:"client_id"`
	Symbol       string `dynamodbav:"symbol" dynamodb:"symbol"`
	PositionSide string `dynamodbav:"position_side" dynamodb:"position_side"`
	// HistoryId    string `dynamodbav:"history_id" dynamodb:"history_id"`
}

func NewBinanceFutureHistory() *BinanceFutureHistory {
	return &BinanceFutureHistory{}
}

func (b *BinanceFutureHistory) IsFound() bool {
	return b.Symbol != ""
}

func (b *BinanceFutureHistory) GetKeyClientID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientId},
	}
}

type BinanceFutureHistoryTable struct {
	TableName string `table:"bn_future_history"`
	*BinanceFutureHistory
}

func NewBinanceFutureHistoryTable() *BinanceFutureHistoryTable {
	return &BinanceFutureHistoryTable{
		BinanceFutureHistory: NewBinanceFutureHistory(),
	}
}

func (b *BinanceFutureHistoryTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BinanceFutureHistoryTable) GetHistoryIdTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "HistoryId", "dynamodb")
	return v
}

func (b *BinanceFutureHistoryTable) GetSymbolTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v
}

func (b *BinanceFutureHistoryTable) GetPositionSideTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v
}
