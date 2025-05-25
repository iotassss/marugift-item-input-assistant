package domain

import "fmt"

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	if len(value) == 0 {
		return Name{}, fmt.Errorf("name cannot be empty")
	}
	if len(value) > 200 {
		return Name{}, fmt.Errorf("name cannot exceed 200 characters")
	}

	return Name{value: value}, nil
}

func (n Name) Value() string  { return n.value }
func (n Name) String() string { return n.value }
