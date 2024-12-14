package dynamodbrepository

import (
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BinanceFutureOpeningPosition struct {
	Symbol             string `dynamodbav:"symbol" dynamodb:"symbol"` // primary key
	ClientId           string `dynamodbav:"client_id" dynamodb:"client_id"`
	OrderType          string `dynamodbav:"order_type" dynamodb:"order_type"`
	Leverage           string `dynamodbav:"leverage" dynamodb:"leverage"`
	PositionSide       string `dynamodbav:"position_side" dynamodb:"position_side"`
	Side               string `dynamodbav:"side" dynamodb:"side"`
	AmountQ            string `dynamodbav:"amount_q" dynamodb:"amount_q"`
	AmountB            string `dynamodbav:"amount_b" dynamodb:"amount_b"`
	BuyOrderCreatedAt  string `dynamodbav:"buy_created_at" dynamodb:"buy_created_at"`
	SellOrderCreatedAt string `dynamodbav:"sell_created_at" dynamodb:"sell_created_at"`
	WatchingConfig     string `dynamodbav:"watching_config" dynamodb:"watching_config"`
}

func (b *BinanceFutureOpeningPosition) IsEmpty() bool {
	return b.Symbol == ""
}

func (b *BinanceFutureOpeningPosition) GetKeyBySymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol": &types.AttributeValueMemberS{Value: b.Symbol},
	}
}

func (b *BinanceFutureOpeningPosition) GetKeyByClientID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientId},
	}
}

func (b *BinanceFutureOpeningPosition) GetKeyByPositionSide() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"position_side": &types.AttributeValueMemberS{Value: b.PositionSide},
	}
}

func (b *BinanceFutureOpeningPosition) GetKeyByPositionSideAndSymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"position_side": &types.AttributeValueMemberS{Value: b.PositionSide},
		"symbol":        &types.AttributeValueMemberS{Value: b.Symbol},
	}
}

func NewBinanceFutureOpeningPosition() *BinanceFutureOpeningPosition {
	return &BinanceFutureOpeningPosition{
		// ExchangeId: 1,
	}
}

type BinanceFutureOpeningPositionTable struct {
	TableName string `table:"bn_future_opening_position"`
	*BinanceFutureOpeningPosition
}

func NewBinanceFutureOpeningPositionTable() *BinanceFutureOpeningPositionTable {
	return &BinanceFutureOpeningPositionTable{
		BinanceFutureOpeningPosition: NewBinanceFutureOpeningPosition(),
	}
}

func (b *BinanceFutureOpeningPositionTable) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BinanceFutureOpeningPositionTable) SetBuyCreatedAt() {
	b.BuyOrderCreatedAt = b.time()
}

func (b *BinanceFutureOpeningPositionTable) SetSellCreatedAt() {
	b.SellOrderCreatedAt = b.time()
}

func (b *BinanceFutureOpeningPositionTable) time() string {
	return time.Now().Format(time.RFC3339)
}

func (b *BinanceFutureOpeningPositionTable) GetLeverageTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Leverage", "dynamodb")
	return v
}

func (b *BinanceFutureOpeningPositionTable) GetOrderTypeTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "OrderType", "dynamodb")
	return v
}

func (b *BinanceFutureOpeningPositionTable) GetClientIdTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ClientId", "dynamodb")
	return v
}

func (b *BinanceFutureOpeningPositionTable) GetAmountQTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountQ", "dynamodb")
	return v
}

func (b *BinanceFutureOpeningPositionTable) GetAmountBTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountB", "dynamodb")
	return v
}

func (b *BinanceFutureOpeningPositionTable) GetWatchingConfigTableField() string {
	v, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "WatchingConfig", "dynamodb")
	return v
}
