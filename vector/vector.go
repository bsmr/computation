package vector

import (
	"github.com/bsmr/computation"
	"github.com/bsmr/computation/internal/common"
	"github.com/bsmr/computation/internal/orientation"
)

type Vector[T computation.Numeric] struct {
	data   *common.Container[T]
	layout orientation.Orientation
}

func New[T computation.Numeric](rank int, values ...T) (*Vector[T], error) {
	data, err := common.New(rank, 1, values...)

	return &Vector[T]{
		data:   data,
		layout: orientation.Check(data),
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

func (v *Vector[T]) String() string {
	return v.data.String()
}

func Addition[T computation.Numeric](a, b *Vector[T]) (*Vector[T], error) {
	c, err := common.Add(a.data, b.data)
	if err != nil {
		return nil, err
	}
	return &Vector[T]{
		data:   c,
		layout: orientation.Check(c),
	}, nil
}

func Transposition[T computation.Numeric](a *Vector[T]) (*Vector[T], error) {
	c, err := common.Transposition(a.data)
	if err != nil {
		return nil, err
	}
	return &Vector[T]{
		data:   c,
		layout: orientation.Check(c),
	}, nil
}
