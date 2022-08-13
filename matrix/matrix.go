package matrix

import (
	"github.com/bsmr/computation"
	"github.com/bsmr/computation/internal/container"
)

// Matrix handles numerical vectors.
type Matrix[T computation.Numeric] struct {
	data *container.Container[T]
}

// New() creates a new MxN matrix.
func New[T computation.Numeric](m, n int, values ...T) (*Matrix[T], error) {
	data, err := container.New(m, n, values...)
	if err != nil {
		return nil, err
	}
	return &Matrix[T]{
		data: data,
	}, nil
}

// String() returns the matrix as an string.
func (m *Matrix[T]) String() string {
	return m.data.String()
}
