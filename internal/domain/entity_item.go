package domain

type Item struct {
	ID string `json:"id"`
}

func NewItem(id string) *Item {
	return &Item{
		ID: id,
	}
}
