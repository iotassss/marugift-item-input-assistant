package domain_test

import (
	"testing"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestSymbol_NewSymbol(t *testing.T) {
	// 任意の文字列で生成できることをテスト
	cases := []string{"A", "B-123", "", "シンボル"}
	for _, c := range cases {
		_, err := domain.NewSymbol(c)
		assert.NoError(t, err)
	}
}

func TestSymbol_Value(t *testing.T) {
	// 任意の文字列で生成したSymbolのValueが正しいことをテスト
	cases := []string{"A", "B-123", "", "シンボル"}
	for _, c := range cases {
		symbol, err := domain.NewSymbol(c)
		assert.NoError(t, err)
		assert.Equal(t, c, symbol.Value())
	}
}
