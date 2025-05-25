package usecase

import "github.com/iotassss/marugift-item-input-assistant/internal/domain"

type SearchItemInputData struct {
	SKU string
}

type SearchItemOutputData struct {
	Name   string
	SKU    string
	Symbol string
	Price  int
}

type SearchItemPresenter interface {
	Present(outputData SearchItemOutputData) error
	PresentError(err error) error
}

type SearchItemUsecase interface {
	Execute(
		input SearchItemInputData,
		presenter SearchItemPresenter,
		itemRepo domain.ItemRepository,
	) error
}
