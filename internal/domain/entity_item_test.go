package domain_test

import (
	"testing"
	"time"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
	unittest "github.com/iotassss/marugift-item-input-assistant/internal/unittest"
)

func newTestItem() *domain.Item {
	item, _ := domain.NewItem(
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
		true,
	)
	return item
}

func TestItem_NewItem(t *testing.T) {
	_, err := domain.NewItem(
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
	)
	if err != nil {
		t.Fatalf("failed to create item: %v", err)
	}

	_, err = domain.NewItem(
		domain.NewDraftID(),
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
	)
	if err == nil {
		t.Fatal("expected error for zero ID, got nil")
	}
}

func TestItem_NewDraftItem(t *testing.T) {
	item, err := domain.NewDraftItem(
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
	)
	if err != nil {
		t.Fatalf("failed to create draft item: %v", err)
	}

	if item.ID().Value() != 0 {
		t.Fatalf("expected draft item ID to be zero, got %d", item.ID().Value())
	}
}

func TestItem_ID(t *testing.T) {
	item := newTestItem()
	if got := item.ID().Value(); got != 1 {
		t.Errorf("ID() = %v, want %v", got, 1)
	}
}

func TestItem_Name(t *testing.T) {
	item := newTestItem()
	if got := item.Name().String(); got != "商品名" {
		t.Errorf("Name() = %v, want %v", got, "商品名")
	}
}

func TestItem_SKU(t *testing.T) {
	item := newTestItem()
	if got := item.SKU().Value(); got != "1234-567" {
		t.Errorf("SKU() = %v, want %v", got, "1234-567")
	}
}

func TestItem_Symbol(t *testing.T) {
	item := newTestItem()
	if got := item.Symbol().Value(); got != "SYM" {
		t.Errorf("Symbol() = %v, want %v", got, "SYM")
	}
}

func TestItem_Price(t *testing.T) {
	item := newTestItem()
	if got := item.Price().Value(); got != 1000 {
		t.Errorf("Price() = %v, want %v", got, 1000)
	}
}

func TestItem_Purpose(t *testing.T) {
	item := newTestItem()
	if got := item.Purpose().Value(); got != 1 {
		t.Errorf("Purpose() = %v, want %v", got, 1)
	}
}

func TestItem_DeliverySchedule(t *testing.T) {
	item := newTestItem()
	if got := item.DeliverySchedule().Value(); got != 1 {
		t.Errorf("DeliverySchedule() = %v, want %v", got, 1)
	}
}

func TestItem_WrappingOption(t *testing.T) {
	item := newTestItem()
	if got := item.WrappingOption().Value(); got != 1 {
		t.Errorf("WrappingOption() = %v, want %v", got, 1)
	}
}
func TestItem_SeasonLabel(t *testing.T) {
	item := newTestItem()
	if got := item.SeasonLabel().Value(); got != "2025_summer" {
		t.Errorf("SeasonLabel() = %v, want %v", got, "2025_summer")
	}
}

func TestItem_ReceptionPeriodStartAt(t *testing.T) {
	item := newTestItem()
	expectedStart := time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC)
	if got := item.ReceptionPeriodStartAt(); !got.Equal(expectedStart) {
		t.Errorf("ReceptionPeriodStartAt() = %v, want %v", got, expectedStart)
	}
}

func TestItem_ReceptionPeriodEndAt(t *testing.T) {
	item := newTestItem()
	expectedEnd := time.Date(2025, 5, 31, 0, 0, 0, 0, time.UTC)
	if got := item.ReceptionPeriodEndAt(); !got.Equal(expectedEnd) {
		t.Errorf("ReceptionPeriodEndAt() = %v, want %v", got, expectedEnd)
	}
}

func TestItem_IsMarutoku(t *testing.T) {
	item := newTestItem()
	if got := item.IsMarutoku(); !got {
		t.Errorf("IsMarutoku() = %v, want %v", got, true)
	}
}

func TestItem_IsCenterShipment(t *testing.T) {
	item := newTestItem()
	if got := item.IsCenterShipment(); !got {
		t.Errorf("IsCenterShipment() = %v, want %v", got, true)
	}
}

func TestItem_SetID(t *testing.T) {
	item := unittest.Must(domain.NewDraftItem(
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
	err := item.SetID(unittest.Must(domain.NewID(1)))
	if err != nil {
		t.Fatalf("SetID failed: %v", err)
	}
	if item.ID().Value() != 1 {
		t.Fatalf("expected ID to be 1, got %d", item.ID().Value())
	}

	err = item.SetID(unittest.Must(domain.NewID(2)))
	if err == nil {
		t.Fatal("expected error when setting ID again, got nil")
	}
}
