package domain

import (
	"fmt"
	"time"
)

type Item struct {
	id                     ID
	sku                    SKU
	name                   Name
	symbol                 Symbol
	price                  Price
	purpose                Purpose
	deliverySchedule       DeliverySchedule
	wrappingOption         WrappingOption
	seasonLabel            SeasonLabel
	receptionPeriodStartAt time.Time
	receptionPeriodEndAt   time.Time
	isMarutoku             bool
	isCenterShipment       bool
}

func NewItem(
	id ID,
	sku SKU,
	name Name,
	symbol Symbol,
	price Price,
	purpose Purpose,
	deliverySchedule DeliverySchedule,
	wrappingOption WrappingOption,
	seasonLabel SeasonLabel,
	receptionPeriodStart time.Time,
	receptionPeriodEnd time.Time,
	isMarutoku bool,
	isCenterShipment bool,
) (*Item, error) {
	if id.Value() < 1 {
		return &Item{}, fmt.Errorf("ID must be greater than or equal to 1")
	}
	if receptionPeriodStart.After(receptionPeriodEnd) {
		return &Item{}, fmt.Errorf("invalid reception period: start must be before end")
	}

	return &Item{
		id:                     id,
		sku:                    sku,
		name:                   name,
		symbol:                 symbol,
		price:                  price,
		purpose:                purpose,
		deliverySchedule:       deliverySchedule,
		wrappingOption:         wrappingOption,
		seasonLabel:            seasonLabel,
		receptionPeriodStartAt: receptionPeriodStart,
		receptionPeriodEndAt:   receptionPeriodEnd,
		isMarutoku:             isMarutoku,
		isCenterShipment:       isCenterShipment,
	}, nil
}

func NewDraftItem(
	sku SKU,
	name Name,
	symbol Symbol,
	price Price,
	purpose Purpose,
	deliverySchedule DeliverySchedule,
	wrappingOption WrappingOption,
	seasonLabel SeasonLabel,
	receptionPeriodStart time.Time,
	receptionPeriodEnd time.Time,
	isMarutoku bool,
	isCenterShipment bool,
) (*Item, error) {
	if receptionPeriodStart.After(receptionPeriodEnd) {
		return &Item{}, fmt.Errorf("invalid reception period: start must be before end")
	}

	return &Item{
		id:                     NewDraftID(),
		sku:                    sku,
		name:                   name,
		symbol:                 symbol,
		price:                  price,
		purpose:                purpose,
		deliverySchedule:       deliverySchedule,
		wrappingOption:         wrappingOption,
		seasonLabel:            seasonLabel,
		receptionPeriodStartAt: receptionPeriodStart,
		receptionPeriodEndAt:   receptionPeriodEnd,
		isMarutoku:             isMarutoku,
		isCenterShipment:       isCenterShipment,
	}, nil
}

func (i *Item) ID() ID                             { return i.id }
func (i *Item) Name() Name                         { return i.name }
func (i *Item) SKU() SKU                           { return i.sku }
func (i *Item) Symbol() Symbol                     { return i.symbol }
func (i *Item) Price() Price                       { return i.price }
func (i *Item) Purpose() Purpose                   { return i.purpose }
func (i *Item) DeliverySchedule() DeliverySchedule { return i.deliverySchedule }
func (i *Item) WrappingOption() WrappingOption     { return i.wrappingOption }
func (i *Item) SeasonLabel() SeasonLabel           { return i.seasonLabel }
func (i *Item) ReceptionPeriodStartAt() time.Time  { return i.receptionPeriodStartAt }
func (i *Item) ReceptionPeriodEndAt() time.Time    { return i.receptionPeriodEndAt }
func (i *Item) IsMarutoku() bool                   { return i.isMarutoku }
func (i *Item) IsCenterShipment() bool             { return i.isCenterShipment }

func (i *Item) SetID(id ID) error {
	if i.id.Value() != 0 {
		return fmt.Errorf("ID is already set: %v", i.id)
	}
	i.id = id

	return nil
}
