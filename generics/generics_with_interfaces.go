package generics

import (
	"errors"
)

type vals interface {
	int | float64 | string
}

func Max[T vals](values []T) (T, error) {

	if len(values) == 0 {
		var zero T
		return zero, errors.New("Max of empty slice")
	}

	max := values[0]

	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max, nil
}
