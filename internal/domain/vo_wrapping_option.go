package domain

import "fmt"

type WrappingOption struct {
	value int
}

func NewWrappingOption(value int) (WrappingOption, error) {
	if value < 1 || value > 5 {
		return WrappingOption{}, fmt.Errorf("value must be between 1 and 5")
	}

	return WrappingOption{value: value}, nil
}

func (w WrappingOption) Value() int { return w.value }
