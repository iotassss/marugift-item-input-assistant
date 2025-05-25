package domain

import "fmt"

type DeliverySchedule struct {
	value int
}

func NewDeliverySchedule(period int) (DeliverySchedule, error) {
	if period < 0 || period > 9 {
		return DeliverySchedule{}, fmt.Errorf("period must be between 0 and 9")
	}

	return DeliverySchedule{value: period}, nil
}

func (ds DeliverySchedule) Value() int { return ds.value }
