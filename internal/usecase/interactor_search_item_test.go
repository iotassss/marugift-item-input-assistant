package usecase_test

import (
	"testing"
	"time"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"github.com/iotassss/marugift-item-input-assistant/internal/repository/memory"
	"github.com/iotassss/marugift-item-input-assistant/internal/unittest"
	"github.com/iotassss/marugift-item-input-assistant/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSearchItemInteractor_Execute(t *testing.T) {
	// repository
	itemRepo := memory.NewItemRepository()

	// test data
	item := unittest.Must(domain.NewItem(
		unittest.Must(domain.NewID(1)),
		unittest.Must(domain.NewSKU("1234-567")),
		unittest.Must(domain.NewName("商品名")),
		unittest.Must(domain.NewSymbol("SYM")),
		unittest.Must(domain.NewPrice(1000)),
		unittest.Must(domain.NewPurpose(1)),
		unittest.Must(domain.NewDeliverySchedule(1)),
		unittest.Must(domain.NewWrappingOption(1)),
		unittest.Must(domain.NewSeasonLabel("2025_summer")),
		time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 5, 31, 0, 0, 0, 0, time.UTC),
		true,
		false,
	))

	// save test data
	unittest.Must(itemRepo.Save(item))

	// input test data
	inputData := usecase.SearchItemInputData{
		SKU: "1234-567",
	}

	// execute
	interactor := usecase.NewSearchItemInteractor()
	presenter := &presenter{}
	err := interactor.Execute(inputData, presenter, itemRepo)

	assert.NoError(t, err)
}

type presenter struct{}

func (p *presenter) Present(outputData usecase.SearchItemOutputData) error {
	if outputData.Name != "商品名" {
		return assert.AnError
	}
	if outputData.SKU != "1234-567" {
		return assert.AnError
	}
	if outputData.Symbol != "SYM" {
		return assert.AnError
	}
	if outputData.Price != 1000 {
		return assert.AnError
	}

	return nil
}

func (p *presenter) PresentError(err error) error {
	return err
}
