package dynamodbfuture

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtCryptoTable struct {
	TableName string `table:"bn_future_crypto"`
	*BnFtCrypto
}

func NewBinanceFutureCryptoTable() *BnFtCryptoTable {
	return &BnFtCryptoTable{
		BnFtCrypto: NewBnFtCrypto(),
	}
}

func NewBinanceFutureCryptoTableWith(qouteUSDT *BnFtCrypto) *BnFtCryptoTable {
	return &BnFtCryptoTable{
		BnFtCrypto: qouteUSDT,
	}
}

func NewBinanceFutureCryptoTableRecord(symbol string, position_side string) *BnFtCrypto {
	table := NewBnFtCrypto()
	table.Symbol = symbol
	if position_side == bnconstant.LONG {
		table.CountingLong = 1
		table.CountingShort = 0
	} else {
		table.CountingLong = 0
		table.CountingShort = 1
	}

	return table
}

func (b *BnFtCryptoTable) GetData() *BnFtCrypto {
	return b.BnFtCrypto
}

func (b *BnFtCryptoTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtCryptoTable) GetKeyBySymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol": &types.AttributeValueMemberS{Value: b.Symbol},
	}
}

func (b *BnFtCryptoTable) GetSymbolTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v, t
}

func (b *BnFtCryptoTable) GetCountingLongTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CountingLong", "dynamodb")
	return v, t
}

func (b *BnFtCryptoTable) GetCountingShortTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CountingShort", "dynamodb")
	return v, t
}

func (b *BnFtCryptoTable) GetCurrentLeverageTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CurrentLeverage", "dynamodb")
	return v, t
}

func (b *BnFtCryptoTable) Transform() {
	b.Symbol = strings.ToUpper(b.Symbol)
}
