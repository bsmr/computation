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
	switch v.layout {
	case orientation.MxOne:
		return v.data.At(index, 0)
	case orientation.OnexN:
		return v.data.At(0, index)
	default:
		var dummy T
		return dummy, fmt.Errorf("cannot handle layout %s", v.layout)
	}
}

func (v *Vector[T]) SetAt(index int, newValue T) (T, error) {
	switch v.layout {
	case orientation.MxOne:
		return v.data.SetAt(index, 0, newValue)
	case orientation.OnexN:
		return v.data.SetAt(0, index, newValue)
	default:
		var dummy T
		return dummy, fmt.Errorf("cannot handle layout %s", v.layout)
	}
}

func (v *Vector[T]) Rank() int {
	m, n := v.data.Rank()
	switch v.layout {
	case orientation.MxOne:
		return m
	case orientation.OnexN:
		return n
	default:
		return -1
	}
}

func (v *Vector[T]) String() string {
	var sb strings.Builder
	sb.WriteString(v.data.String())
	sb.WriteString(fmt.Sprintf("mode: %s", v.layout))
	return sb.String()
}

func (v *Vector[T]) Equal(a *Vector[T]) bool {
	if v.Rank() != a.Rank() {
		return false
	}
	for i := 0; i < v.Rank(); i++ {
		v1, _ := v.At(i)
		v2, _ := a.At(i)
		//if !reflect.DeepEqual(v1, v2) {
		if v1 != v2 {
			return false
		}
	}
	return true
}

func (v *Vector[T]) EqualReal(a *Vector[T]) bool {
	if v.layout != a.layout {
		return false
	}
	return v.Equal(a)
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

func ScalarMultiplication[T computation.Numeric](s T, a *Vector[T]) (*Vector[T], error) {
	c, err := container.ScalarMatrixMultiplication(s, a.data)
	if err != nil {
		return nil, err
	}
	return &Vector[T]{
		data:   c,
		layout: orientation.Check(c),
	}, nil
}
