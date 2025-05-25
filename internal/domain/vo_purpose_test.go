package domain_test

import (
	"testing"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestPurpose_NewPurpose(t *testing.T) {
	_, err := domain.NewPurpose(0)
	assert.Error(t, err)

	_, err = domain.NewPurpose(1)
	assert.NoError(t, err)

	_, err = domain.NewPurpose(9)
	assert.NoError(t, err)

	_, err = domain.NewPurpose(10)
	assert.Error(t, err)
}

func TestPurpose_Value(t *testing.T) {
	purpose, err := domain.NewPurpose(5)
	assert.NoError(t, err)
	assert.Equal(t, 5, purpose.Value())

	purpose, err = domain.NewPurpose(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, purpose.Value())
}
