package domain_test

import (
	"testing"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestDeliverySchedule_NewDeliverySchedule(t *testing.T) {
	_, err := domain.NewDeliverySchedule(-1)
	assert.Error(t, err)

	_, err = domain.NewDeliverySchedule(0)
	assert.NoError(t, err)

	_, err = domain.NewDeliverySchedule(9)
	assert.NoError(t, err)

	_, err = domain.NewDeliverySchedule(10)
	assert.Error(t, err)
}

func TestDeliverySchedule_Value(t *testing.T) {
	deliverySchedule, err := domain.NewDeliverySchedule(5)
	assert.NoError(t, err)
	assert.Equal(t, 5, deliverySchedule.Value())

	deliverySchedule, err = domain.NewDeliverySchedule(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, deliverySchedule.Value())
}
