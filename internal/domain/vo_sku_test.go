package domain_test

import (
	"testing"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestSKU_NewSKU(t *testing.T) {
	_, err := domain.NewSKU("1234-567")
	assert.NoError(t, err)
	_, err = domain.NewSKU("123-4567") // '-'が4桁目でない
	assert.Error(t, err)
	_, err = domain.NewSKU("abcde-fgh") // 数字でない
	assert.Error(t, err)
	_, err = domain.NewSKU("1234_567") // 5桁目が'-'でない
	assert.Error(t, err)
	_, err = domain.NewSKU("") // 空文字
	assert.Error(t, err)
}

func TestSKU_Value(t *testing.T) {
	sku, err := domain.NewSKU("1234-567")
	assert.NoError(t, err)
	assert.Equal(t, "1234-567", sku.Value())
}

func TestSKU_Equal(t *testing.T) {
	sku1, err := domain.NewSKU("1234-567")
	assert.NoError(t, err)
	sku2, err := domain.NewSKU("1234-567")
	assert.NoError(t, err)
	sku3, err := domain.NewSKU("1234-568")
	assert.NoError(t, err)

	assert.True(t, sku1.Equal(sku2))
	assert.False(t, sku1.Equal(sku3))
}
