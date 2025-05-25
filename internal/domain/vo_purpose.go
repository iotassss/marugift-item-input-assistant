package domain

import "fmt"

type Purpose struct {
	value int
}

func NewPurpose(value int) (Purpose, error) {
	if value < 1 || value > 9 {
		return Purpose{}, fmt.Errorf("value must be between 1 and 9")
	}

	return Purpose{value: value}, nil
}

func (p Purpose) Value() int { return p.value }
