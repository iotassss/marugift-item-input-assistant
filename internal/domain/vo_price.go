package domain

import "fmt"

type Price struct {
	value int
}

func NewPrice(value int) (Price, error) {
	if value < 0 {
		return Price{}, fmt.Errorf("price must be greater than or equal to 0")
	}
	// 表示枠の上限からはみ出さないようにするため5桁まで
	if value > 99999 {
		return Price{}, fmt.Errorf("price must be less than or equal to 99999")
	}

	return Price{value: value}, nil
}

func (p Price) Value() int {
	return p.value
}
