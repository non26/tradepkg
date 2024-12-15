package dynamodbrepository

import (
	"reflect"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BinanceFutureQouteUSDT struct {
	Symbol string `dynamodb:"symbol" dynamodbav:"symbol"` // primary key
	// CurrentLeverage int    `dynamodb:"current_leverage" dynamodbav:"current_leverage"`
	CountingLong  int `dynamodb:"counting_long" dynamodbav:"counting_long"`
	CountingShort int `dynamodb:"counting_short" dynamodbav:"counting_short"`
}

type counting int

func (c counting) Int() int {
	return int(c)
}

func (c counting) String() string {
	return strconv.Itoa(c.Int())
}

func newBinanceFutureQouteUSDT() *BinanceFutureQouteUSDT {
	return &BinanceFutureQouteUSDT{}
}

func (b *BinanceFutureQouteUSDT) GetNextCountingLong() counting {
	return counting(b.CountingLong + 1)
}

func (b *BinanceFutureQouteUSDT) GetNextCountingShort() counting {
	return counting(b.CountingShort + 1)
}

func (b *BinanceFutureQouteUSDT) SetCountingLong(counting int) {
	b.CountingLong = counting
}

func (b *BinanceFutureQouteUSDT) SetCountingShort(counting int) {
	b.CountingShort = counting
}

func (b *BinanceFutureQouteUSDT) GetSymbol() string {
	return b.Symbol
}

func (b *BinanceFutureQouteUSDT) GetCountingLong() int {
	return b.CountingLong
}

func (b *BinanceFutureQouteUSDT) GetCountingShort() int {
	return b.CountingShort
}

func (b *BinanceFutureQouteUSDT) SetSymbol(symbol string) {
	b.Symbol = symbol
}

func (b *BinanceFutureQouteUSDT) IsFound() bool {
	return b.Symbol != ""
}

// func (b *BinanceFutureQouteUSDT) GetCurrentLeverage() int {
// 	return b.CurrentLeverage
// }

// func (b *BinanceFutureQouteUSDT) SetCurrentLeverage(leverage int) {
// 	b.CurrentLeverage = leverage
// }

type BinanceFutureQouteUSTDTable struct {
	TableName string `table:"bn_future_qoute_usdt"`
	*BinanceFutureQouteUSDT
}

func (b *BinanceFutureQouteUSTDTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func NewBinanceFutureQouteUSTDTable() *BinanceFutureQouteUSTDTable {
	return &BinanceFutureQouteUSTDTable{
		BinanceFutureQouteUSDT: newBinanceFutureQouteUSDT(),
	}
}

func (b *BinanceFutureQouteUSTDTable) GetKeyBySymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol": &types.AttributeValueMemberS{Value: b.Symbol},
	}
}

func (b *BinanceFutureQouteUSTDTable) GetSymbolTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v
}

func (b *BinanceFutureQouteUSTDTable) GetCountingLongTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CountingLong", "dynamodb")
	return v
}

func (b *BinanceFutureQouteUSTDTable) GetCountingShortTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CountingShort", "dynamodb")
	return v
}

func (b *BinanceFutureQouteUSTDTable) GetCurrentLeverageTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CurrentLeverage", "dynamodb")
	return v
}
