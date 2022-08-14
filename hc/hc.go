package hc

import (
	"github.com/bsmr/computation"
	"github.com/bsmr/computation/matrix"
)

// Homogeneours Coordinates

/*
        /                    \
        | a11  a12  a13  a14 |
        |                    |
        | a21  a22  a23  a24 |
	A = |                    |
        | a31  a32  a33  a34 |
        |                    |
        |  0    0    0    1  |
        \                    /
*/

func NewNull[T computation.Numeric]() (*matrix.Matrix[T], error) {
	return matrix.New[T](4, 4)
}

func NewIdentity[T computation.Numeric]() (*matrix.Matrix[T], error) {
	values := []T{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	return matrix.New(4, 4, values...)
}

func NewTranslation[T computation.Numeric](x, y, z T) (*matrix.Matrix[T], error) {
	values := []T{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
	return matrix.New(4, 4, values...)
}
