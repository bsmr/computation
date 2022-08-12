package vector

import "fmt"

type Vector[T any] struct {
	rank   int
	values []T
}

func New[T any](rank int, values ...T) (*Vector[T], error) {
	lv := len(values)
	if rank < lv {
		return nil, fmt.Errorf("too many values %d for rank %d", lv, rank)
	}

	vector := &Vector[T]{
		rank: rank,
	}

	copy(values, vector.values)

	return vector, nil
}

func (v *Vector[T]) At(index int) (T, error) {
	var r T

	if index < 0 {
		return r, fmt.Errorf("out of bounds: index %d is less than 0", index)
	}
	if index > (v.rank - 1) {
		return r, fmt.Errorf("out of bounds: index %d is greater than %d", index, v.rank)
	}

	return v.values[index], nil
}
