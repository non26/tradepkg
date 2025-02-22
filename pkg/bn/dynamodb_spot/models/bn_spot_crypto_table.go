package dynamodbspot

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnSpotCryptoTable struct {
	TableName string `table:"bn_spot_crypto"`
	*BnSpotCrypto
}

func NewBinanceSpotCryptoTable() *BnSpotCryptoTable {
	return &BnSpotCryptoTable{
		BnSpotCrypto: NewBnSpotCrypto(),
	}
}

func NewBinanceSpotCryptoTableWith(qouteUSDT *BnSpotCrypto) *BnSpotCryptoTable {
	return &BnSpotCryptoTable{
		BnSpotCrypto: qouteUSDT,
	}
}

func NewBinanceSpotCryptoTableRecord(symbol string) *BnSpotCrypto {
	table := NewBnSpotCrypto()
	table.Symbol = symbol
	table.Counting = 1
	return table
}

func (b *BnSpotCryptoTable) GetData() *BnSpotCrypto {
	return b.BnSpotCrypto
}

func (b *BnSpotCryptoTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnSpotCryptoTable) GetKeyBySymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol": &types.AttributeValueMemberS{Value: b.Symbol},
	}
}

func (b *BnSpotCryptoTable) GetSymbolTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
	return v, t
}

func (b *BnSpotCryptoTable) GetCountingTableField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Counting", "dynamodb")
	return v, t
}

func (b *BnSpotCryptoTable) Transform() {
	b.Symbol = strings.ToUpper(b.Symbol)
}
