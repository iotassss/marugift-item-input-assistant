package mysql

import (
	"errors"
	"time"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"gorm.io/gorm"
)

type ItemModel struct {
	gorm.Model
	Name                   string    `gorm:"column:name;not null;size:100"`
	SKU                    string    `gorm:"column:sku;unique;not null;size:8"`
	Symbol                 string    `gorm:"column:symbol;not null;size:20"`
	Price                  int       `gorm:"column:price;not null"`
	Purpose                int       `gorm:"column:purpose;not null"`
	DeliverySchedule       int       `gorm:"column:delivery_schedule;not null"`
	WrappingOption         int       `gorm:"column:wrapping_option;not null"`
	SeasonLabel            string    `gorm:"column:season_label;not null;size:50"`
	ReceptionPeriodStartAt time.Time `gorm:"column:reception_period_start_at;not null"`
	ReceptionPeriodEndAt   time.Time `gorm:"column:reception_period_end_at;not null"`
	IsMarutoku             bool      `gorm:"column:is_marutoku;not null"`
	IsCenterShipment       bool      `gorm:"column:is_center_shipment;not null"`
}

func (ItemModel) TableName() string {
	return "items"
}

func toItemDomain(model ItemModel) (*domain.Item, error) {
	id, err := domain.NewID(int(model.ID))
	if err != nil {
		return nil, err
	}
	sku, err := domain.NewSKU(model.SKU)
	if err != nil {
		return nil, err
	}
	name, err := domain.NewName(model.Name)
	if err != nil {
		return nil, err
	}
	symbol, err := domain.NewSymbol(model.Symbol)
	if err != nil {
		return nil, err
	}
	price, err := domain.NewPrice(model.Price)
	if err != nil {
		return nil, err
	}
	purpose, err := domain.NewPurpose(model.Purpose)
	if err != nil {
		return nil, err
	}
	deliverySchedule, err := domain.NewDeliverySchedule(model.DeliverySchedule)
	if err != nil {
		return nil, err
	}
	wrappingOption, err := domain.NewWrappingOption(model.WrappingOption)
	if err != nil {
		return nil, err
	}
	seasonLabel, err := domain.NewSeasonLabel(model.SeasonLabel)
	if err != nil {
		return nil, err
	}
	receptionPeriodStartAt := model.ReceptionPeriodStartAt
	receptionPeriodEndAt := model.ReceptionPeriodEndAt

	item, err := domain.NewItem(
		id,
		sku,
		name,
		symbol,
		price,
		purpose,
		deliverySchedule,
		wrappingOption,
		seasonLabel,
		receptionPeriodStartAt,
		receptionPeriodEndAt,
		true,
		true,
	)
	if err != nil {
		return nil, err
	}

	return item, nil
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) FindBySKU(sku domain.SKU) (*domain.Item, error) {
	var model ItemModel
	if err := r.db.First(&model, "sku = ?", sku.Value()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrEntityNotFound
		}
		return nil, err
	}
	return toItemDomain(model)
}

func (r *ItemRepository) Save(item *domain.Item) (*domain.Item, error) {
	model := ItemModel{
		Model:  gorm.Model{ID: uint(item.ID().Value())},
		Name:   item.Name().Value(),
		SKU:    item.SKU().Value(),
		Symbol: item.Symbol().Value(),
		Price:  item.Price().Value(),
	}
	if err := r.db.Save(&model).Error; err != nil {
		return nil, err
	}
	id, err := domain.NewID(int(model.ID))
	if err != nil {
		return nil, err
	}
	if err = item.SetID(id); err != nil {
		return nil, err
	}
	return item, nil
}
