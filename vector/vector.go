package vector

import (
	"fmt"
	"strings"

	"github.com/bsmr/computation"
	"github.com/bsmr/computation/internal/container"
	"github.com/bsmr/computation/internal/orientation"
)

// Vector handles numerical vectors.
type Vector[T computation.Numeric] struct {
	data   *container.Container[T]
	layout orientation.Orientation
}

// New() creates a new value. The rank is specified by the number of values passed in.
func New[T computation.Numeric](values ...T) (*Vector[T], error) {
	rank := len(values)
	data, err := container.New(rank, 1, values...)

	return &Vector[T]{
		data:   data,
		layout: orientation.Check(data),
	}, err
}

// NewValue() creates a new vector of rank with all values set to zero.
func NewZero[T computation.Numeric](rank int) (*Vector[T], error) {
	slice := make([]T, rank)
	return New(slice...)
}

// NewValue() creates a new vector of rank with all values set to value.
func NewValue[T computation.Numeric](rank int, value T) (*Vector[T], error) {
	slice := make([]T, rank)
	for p := 0; p < rank; p++ {
		slice[p] = value
	}
	return New(slice...)
}

// At() returns the value at index.
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

// SetAt() sets newValue at index, and returns the old value.
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

// Rank() returns the rank of V.
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

// String() returns the components as a string.
func (v *Vector[T]) String() string {
	var sb strings.Builder
	sb.WriteString(v.data.String())
	sb.WriteString(fmt.Sprintf("mode: %s", v.layout))
	return sb.String()
}

// Equal() checks if V and A are equal.
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

// EqualReal() is a little bit strickter than Equal().
func (v *Vector[T]) EqualReal(a *Vector[T]) bool {
	if v.layout != a.layout {
		return false
	}
	return v.Equal(a)
}

// Values() returns all the elements as a slice of T.
func (v *Vector[T]) Values() ([]T, error) {
	switch v.layout {
	case orientation.MxOne:
		c, err := v.data.Column(0)
		if err != nil {
			return nil, err
		}
		return c.Values()
	case orientation.OnexN:
		r, err := v.data.Row(0)
		if err != nil {
			return nil, err
		}
		return r.Values()
	default:
		return nil, fmt.Errorf("Values() not supported for %s", v.layout)
	}
}

// Addition() returns A+B.
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

// Transposition() returns A^T.
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

// ScalarMultiplication() returns s*A.
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
