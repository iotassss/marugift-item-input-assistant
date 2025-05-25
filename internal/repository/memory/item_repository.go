package memory

import (
	"errors"
	"sync"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
)

type ItemRepository struct {
	data map[domain.ID]*domain.Item
	mu   sync.RWMutex
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		data: make(map[domain.ID]*domain.Item),
	}
}

func (r *ItemRepository) FindBySKU(sku domain.SKU) (*domain.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, item := range r.data {
		if item.SKU().Equal(sku) {
			return item, nil
		}
	}

	return nil, errors.New("item not found")
}

func (r *ItemRepository) Save(item *domain.Item) (*domain.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if item == nil {
		return nil, errors.New("item cannot be nil")
	}

	if item.ID().Value() == 0 {
		newID := domain.NewDraftID()
		item.SetID(newID)
	}
	r.data[item.ID()] = item

	return item, nil
}
