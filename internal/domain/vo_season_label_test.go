package domain_test

import (
	"testing"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestSeasonLabel_NewSeasonLabel(t *testing.T) {
	_, err := domain.NewSeasonLabel("2025_summer")
	assert.NoError(t, err)

	_, err = domain.NewSeasonLabel("2025_winter")
	assert.NoError(t, err)

	_, err = domain.NewSeasonLabel("2024_spring")
	assert.Error(t, err)

	_, err = domain.NewSeasonLabel("")
	assert.Error(t, err)
}

func TestSeasonLabel_Value(t *testing.T) {
	seasonLabel, err := domain.NewSeasonLabel("2025_summer")
	assert.NoError(t, err)
	assert.Equal(t, "2025_summer", seasonLabel.Value())

	seasonLabel, err = domain.NewSeasonLabel("2025_winter")
	assert.NoError(t, err)
	assert.Equal(t, "2025_winter", seasonLabel.Value())
}
