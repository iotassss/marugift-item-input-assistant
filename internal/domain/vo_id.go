package domain

import "fmt"

type ID struct {
	value int
}

func NewID(value int) (ID, error) {
	if value <= 0 {
		return ID{}, fmt.Errorf("ID must be greater than 0")
	}
	return ID{value: value}, nil
}

func NewDraftID() ID {
	return ID{value: 0}
}

func (id ID) Value() int { return id.value }
