package usecase

import "github.com/iotassss/marugift-item-input-assistant/internal/domain"

type SearchItemInteractor struct{}

func NewSearchItemInteractor() *SearchItemInteractor {
	return &SearchItemInteractor{}
}

func (uc *SearchItemInteractor) Execute(
	input SearchItemInputData,
	presenter SearchItemPresenter,
	itemRepo domain.ItemRepository,
) error {
	sku, err := domain.NewSKU(input.SKU)
	if err != nil {
		return presenter.PresentError(err)
	}

	item, err := itemRepo.FindBySKU(sku)
	if err != nil {
		return presenter.PresentError(err)
	}

	outputData := SearchItemOutputData{
		Name:   item.Name().Value(),
		SKU:    item.SKU().Value(),
		Symbol: item.Symbol().Value(),
		Price:  item.Price().Value(),
	}
	return presenter.Present(outputData)
}
