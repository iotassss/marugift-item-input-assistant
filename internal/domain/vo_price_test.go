package domain_test

import (
	"testing"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestPrice_NewPrice(t *testing.T) {
	_, err := domain.NewPrice(-1)
	assert.Error(t, err)

	_, err = domain.NewPrice(0)
	assert.NoError(t, err)

	_, err = domain.NewPrice(99999)
	assert.NoError(t, err)

	_, err = domain.NewPrice(100000)
	assert.Error(t, err)
}

func TestPrice_Value(t *testing.T) {
	price, err := domain.NewPrice(0)
	assert.NoError(t, err)
	assert.Equal(t, 0, price.Value())

	price, err = domain.NewPrice(99999)
	assert.NoError(t, err)
	assert.Equal(t, 99999, price.Value())
}
