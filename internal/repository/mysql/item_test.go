package mysql_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	"github.com/iotassss/marugift-item-input-assistant/internal/repository/mysql"
	"github.com/iotassss/marugift-item-input-assistant/internal/unittest"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func newTestItem() *domain.Item {
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
	fmt.Println("Created test item:", item)
	return item
}

func toItemModel(item *domain.Item) *mysql.ItemModel {
	fmt.Println("Converting item to model:", item)
	return &mysql.ItemModel{
		Model:                  gorm.Model{ID: uint(item.ID().Value())},
		Name:                   item.Name().Value(),
		SKU:                    item.SKU().Value(),
		Symbol:                 item.Symbol().Value(),
		Price:                  item.Price().Value(),
		Purpose:                item.Purpose().Value(),
		DeliverySchedule:       item.DeliverySchedule().Value(),
		WrappingOption:         item.WrappingOption().Value(),
		SeasonLabel:            item.SeasonLabel().Value(),
		ReceptionPeriodStartAt: item.ReceptionPeriodStartAt(),
		ReceptionPeriodEndAt:   item.ReceptionPeriodEndAt(),
		IsMarutoku:             item.IsMarutoku(),
		IsCenterShipment:       item.IsCenterShipment(),
	}
}

func TestItem_FindBySKU(t *testing.T) {
	db := setupTestDB(t)
	repo := mysql.NewItemRepository(db)
	item := newTestItem()

	if err := db.Save(toItemModel(item)).Error; err != nil {
		t.Fatalf("failed to save item: %v", err)
	}

	found, err := repo.FindBySKU(item.SKU())
	require.NoError(t, err)
	require.Equal(t, item.Name(), found.Name())
}
