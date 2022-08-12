package vector

import (
	"fmt"

	"github.com/bsmr/computation"
)

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

func Add[T computation.Numeric](vs ...*Vector[T]) (*Vector[T], error) {
	return operation(func(a, b T) T { return a + b }, vs...)
}

func Mul[T computation.Numeric](vs ...*Vector[T]) (*Vector[T], error) {
	return operation(func(a, b T) T { return a * b }, vs...)
}

func operation[T computation.Numeric](op func(a, b T) T, vs ...*Vector[T]) (*Vector[T], error) {
	var vr *Vector[T]
	var err error

	for index, v := range vs {
		switch index {
		case 0:
			vr, err = New(v.rank, v.values...)
			if err != nil {
				return nil, err
			}
		default:
			for i := 0; i < vr.rank; i++ {
				vr.values[i] = op(vr.values[i], v.values[i])
			}
		}
	}

	return vr, nil
}
