package vector

import (
	"fmt"

	"github.com/bsmr/computation"
	"github.com/bsmr/computation/internal/common"
)

type Vector[T computation.Numeric] struct {
	rank   int
	values []T
}

const (
	minRank = 1
)

func New[T computation.Numeric](rank int, values ...T) (*Vector[T], error) {
	if rank < minRank {
		return nil, fmt.Errorf("rank must be %d or higher", minRank)
	}

	lv := len(values)
	if rank < lv {
		return nil, fmt.Errorf("too many values %d for rank %d", lv, rank)
	}

	vector := &Vector[T]{
		rank:   rank,
		values: make([]T, rank),
	}

	copy(vector.values, values)

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

func (v *Vector[T]) SetAt(index int, newValue T) (oldValue T, err error) {
	v.values[index], oldValue = newValue, v.values[index]
	return
}

func (v *Vector[T]) Copy() (common.Container[T], error) {
	return New(v.rank, v.values...)
}

func (v *Vector[T]) Rank() int {
	return v.rank
}

func (v *Vector[T]) Values() []T {
	return v.values
}
