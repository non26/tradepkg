package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinanceFail_success(t *testing.T) {
	bn_fail := NewBinanceFail(400, 1000, "test")
	assert.Equal(t, bn_fail.Error(), "BnHttp:400,BnCode:1000,BnMsg:test")
}
