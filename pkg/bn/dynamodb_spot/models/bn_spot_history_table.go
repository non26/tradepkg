package dynamodbspot

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	bntime "github.com/non26/tradepkg/pkg/bn/bn_time"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnSpotHistoryTable struct {
	TableName string `table:"bn_spot_history"`
	*BnSpotHistory
}

func NewBinanceSpotHistoryTable() *BnSpotHistoryTable {
	return &BnSpotHistoryTable{
		BnSpotHistory: NewBinanceSpotHistory(),
	}
}

func NewBinanceSpotHistoryTableWith(history *BnSpotHistory) *BnSpotHistoryTable {
	return &BnSpotHistoryTable{
		BnSpotHistory: history,
	}
}

func (b *BnSpotHistoryTable) GetData() *BnSpotHistory {
	return b.BnSpotHistory
}

func (b *BnSpotHistoryTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnSpotHistoryTable) GetHistoryIdTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "HistoryId", "dynamodb")
	return v, t
}

func (b *BnSpotHistoryTable) GetSymbolTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v, t
}

func (b *BnSpotHistoryTable) GetCreatedAtTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CreatedAt", "dynamodb")
	return v, t
}

func (b *BnSpotHistoryTable) GetKeyClientID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientId},
	}
}

func (b *BnSpotHistoryTable) Transform() {
	b.Symbol = strings.ToUpper(b.Symbol)
}

func (b *BnSpotHistoryTable) SetCreatedAt() {
	b.CreatedAt = bntime.GetDBTime()
}
