package domain

import "fmt"

var labels = map[string]struct{}{
	"2025_summer": {},
	"2025_winter": {},
}

type SeasonLabel struct {
	value string
}

func NewSeasonLabel(label string) (SeasonLabel, error) {
	if _, ok := labels[label]; ok {
		return SeasonLabel{value: label}, nil
	}

	return SeasonLabel{}, fmt.Errorf("invalid season label: %s", label)
}

func (s SeasonLabel) Value() string { return s.value }
