package domain_test

import (
	"testing"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestWrappingOption_NewWrappingOption(t *testing.T) {
	_, err := domain.NewWrappingOption(0)
	assert.Error(t, err)

	_, err = domain.NewWrappingOption(1)
	assert.NoError(t, err)

	_, err = domain.NewWrappingOption(5)
	assert.NoError(t, err)

	_, err = domain.NewWrappingOption(6)
	assert.Error(t, err)
}
