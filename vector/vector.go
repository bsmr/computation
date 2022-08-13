package vector

import (
	"github.com/bsmr/computation"
	"github.com/bsmr/computation/internal/common"
)

type Vector[T computation.Numeric] struct {
	data *common.Container[T]
}

func New[T computation.Numeric](rank int, values ...T) (*Vector[T], error) {
	data, err := common.New(rank, 1, values...)

	return &Vector[T]{
		data: data,
	}, err
}

func (v *Vector[T]) At(index int) (T, error) {
	return v.data.At(index, 1)
}

func (v *Vector[T]) SetAt(index int, newValue T) (T, error) {
	return v.data.SetAt(index, 1, newValue)
}

func (v *Vector[T]) Rank() int {
	m, _ := v.data.Rank()
	return m
}
