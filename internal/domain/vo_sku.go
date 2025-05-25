package domain

import "fmt"

type SKU struct {
	value string
}

func NewSKU(value string) (SKU, error) {
	// 8桁のSKUであることを確認する
	if len(value) != 8 {
		return SKU{}, fmt.Errorf("SKU must be 8 characters long")
	}
	// 先頭4桁が数字であることを確認する
	for i := 0; i < 4; i++ {
		if value[i] < '0' || value[i] > '9' {
			return SKU{}, fmt.Errorf("SKU must start with 4 digits")
		}
	}
	// 5桁目が'-'であることを確認する
	if value[4] != '-' {
		return SKU{}, fmt.Errorf("SKU must contain '-' at the 5th position")
	}
	// 末尾3桁が数字であることを確認する
	for i := 5; i < 7; i++ {
		if value[i] < '0' || value[i] > '9' {
			return SKU{}, fmt.Errorf("SKU must end with 3 digits")
		}
	}

	return SKU{value: value}, nil
}

func (s SKU) Value() string {
	return s.value
}

func (s SKU) Equal(other SKU) bool {
	return s.value == other.value
}
