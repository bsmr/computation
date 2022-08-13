package vector

import (
	"fmt"
	"strings"

	"github.com/bsmr/computation"
	"github.com/bsmr/computation/internal/container"
	"github.com/bsmr/computation/internal/orientation"
)

type Vector[T computation.Numeric] struct {
	data   *container.Container[T]
	layout orientation.Orientation
}

func New[T computation.Numeric](values ...T) (*Vector[T], error) {
	rank := len(values)
	data, err := container.New(rank, 1, values...)

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
	var sb strings.Builder
	sb.WriteString(v.data.String())
	sb.WriteString(fmt.Sprintf("mode: %s", v.layout))
	return sb.String()
}

func Addition[T computation.Numeric](a, b *Vector[T]) (*Vector[T], error) {
	c, err := container.Addition(a.data, b.data)
	if err != nil {
		return nil, err
	}
	return &Vector[T]{
		data:   c,
		layout: orientation.Check(c),
	}, nil
}

func Transposition[T computation.Numeric](a *Vector[T]) (*Vector[T], error) {
	c, err := container.Transposition(a.data)
	if err != nil {
		return nil, err
	}
	return &Vector[T]{
		data:   c,
		layout: orientation.Check(c),
	}, nil
}
