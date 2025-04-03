package utils_test

import (
	"testing"

	"github.com/non26/tradepkg/pkg/bn/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetQueryStringBinanceSignature(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		query_string := "symbol=BTCUSDT&side=BUY&type=LIMIT&quantity=1&price=9000&timeInForce=GTC&recvWindow=5000&timestamp=1499827319559"
		signature := "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71"
		result := utils.GetQueryStringBinanceSignature(query_string, signature)
		assert.Equal(t, result, "symbol=BTCUSDT&side=BUY&type=LIMIT&quantity=1&price=9000&timeInForce=GTC&recvWindow=5000&timestamp=1499827319559&signature=c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71")
	})
}

type TestModel_1 struct {
	Symbol string `json:"symbol"`
	Price  int64  `json:"price"`
}

type TestModel_2 struct {
	Symbol  string   `json:"symbol"`
	Symbol2 []string `json:"symbol2"`
	Price   int64    `json:"price"`
}

type TestModel_3 struct {
	Symbol  string      `json:"symbol"`
	Symbol2 TestModel_1 `json:"symbol2"`
	Price   int64       `json:"price"`
}

type TestModel_4 struct {
	Symbol  string        `json:"symbol"`
	Symbol2 []TestModel_1 `json:"symbol2"`
	Price   int64         `json:"price"`
}

func TestCreateQueryStringFrom(t *testing.T) {

	t.Run("success: plain object", func(t *testing.T) {
		model := TestModel_1{
			Symbol: "BTCUSDT",
			Price:  9000,
		}
		expected := "price=9000&symbol=BTCUSDT"
		actual := utils.CreateQueryStringFrom(&model)
		assert.Equal(t, expected, actual)
	})

	t.Run("success: array of objects", func(t *testing.T) {
		model := TestModel_2{
			Symbol:  "BTCUSDT",
			Symbol2: []string{"BTCUSDT", "ETHUSDT"},
			Price:   9000,
		}
		expected := "price=9000&symbol=BTCUSDT&symbol2=%5BBTCUSDT+ETHUSDT%5D"
		actual := utils.CreateQueryStringFrom(&model)
		assert.Equal(t, expected, actual)
	})

	t.Run("success: nested object", func(t *testing.T) {
		model := TestModel_3{
			Symbol:  "BTCUSDT",
			Symbol2: TestModel_1{Symbol: "BTCUSDT", Price: 9000},
			Price:   9000,
		}
		// expected := "price=9000&symbol=BTCUSDT&symbol2={\"symbol\":\"BTCUSDT\",\"price\":9000}"
		expected := "price=9000&symbol=BTCUSDT&symbol2=%7BBTCUSDT+9000%7D"
		actual := utils.CreateQueryStringFrom(&model)
		assert.Equal(t, expected, actual)
	})

	t.Run("success: array of nested objects", func(t *testing.T) {
		model := TestModel_4{
			Symbol: "BTCUSDT",
			Symbol2: []TestModel_1{
				{Symbol: "BTCUSDT", Price: 9000},
				{Symbol: "ETHUSDT", Price: 10000}},
			Price: 9000,
		}
		expected := "price=9000&symbol=BTCUSDT&symbol2=%5B%7BBTCUSDT+9000%7D+%7BETHUSDT+10000%7D%5D"
		actual := utils.CreateQueryStringFrom(&model)
		assert.Equal(t, expected, actual)
	})
}
