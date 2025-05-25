package domain

type ItemRepository interface {
	FindBySKU(sku SKU) (*Item, error)
	Save(item *Item) (*Item, error)
}
