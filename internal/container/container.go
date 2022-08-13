package container

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bsmr/computation"
)

type Container[T computation.Numeric] struct {
	m, n int
	data []T
}

const (
	minRankM = 1
	minRankN = 1
)

// New() creates a new container with a dimension of (m,n).
// Additional values for initializing the container may be supplied.
func New[T computation.Numeric](m, n int, values ...T) (*Container[T], error) {
	if m < minRankM {
		return nil, fmt.Errorf("m is %d, it must be equal or greater to %d", m, minRankM)
	}
	if n < minRankN {
		return nil, fmt.Errorf("n is %d, it must be equal or greater to %d", n, minRankN)
	}
	lv := len(values)
	if lv > (m * n) {
		return nil, fmt.Errorf("the number of values (%d) exceeds %d*%d (%d) entries", lv, m, n, m*n)
	}
	v := &Container[T]{
		m:    m,
		n:    n,
		data: make([]T, m*n),
	}

	if nc := copy(v.data, values); nc != lv {
		return nil, fmt.Errorf("number of values copied is %d, expected %d", nc, lv)
	}

	return v, nil
}

// Rank() returns the dimensions (m,n) of the container.
func (c *Container[T]) Rank() (int, int) {
	return c.m, c.n
}

func between(min, pos, max int) bool {
	return ((min <= pos) && (pos < max))
}

func (c *Container[T]) position(i, j int) (int, error) {
	if !between(0, i, c.m) {
		return 0, fmt.Errorf("i is %d, it must be in range [0, %d[", i, c.m)
	}
	if !between(0, j, c.n) {
		return 0, fmt.Errorf("j is %d, it must be in range [0, %d[", j, c.n)
	}
	return (i * c.n) + j, nil
}

// At() returns the value at position (i,j).
// It may raise an error, if position(i,j) is out of bounds.
func (c *Container[T]) At(i, j int) (value T, err error) {
	at, err := c.position(i, j)
	if err != nil {
		return
	}
	return c.data[at], nil
}

// SetAt() sets a new value at postion (i,j), and returns its old value.
// It may raise an error, if position(i,j) is out of bounds.
func (c *Container[T]) SetAt(i, j int, newValue T) (oldValue T, err error) {
	at, err := c.position(i, j)
	if err != nil {
		return
	}
	c.data[at], oldValue = newValue, c.data[at]
	return
}

// Copy() returns a new copy of the current container.
func (c *Container[T]) Copy() (*Container[T], error) {
	return New(c.m, c.n, c.data...)
}

// String() retrurn the container formatted as a string.
func (c *Container[T]) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("\n<%dx%d>\n", c.m, c.n))

	for i := 0; i < c.m; i++ {
		sb.WriteString("[")
		for j := 0; j < c.n; j++ {
			if j != 0 {
				sb.WriteString("  ")
			}
			v, _ := c.At(i, j)
			sb.WriteString(fmt.Sprintf("<(%d,%d):%v>", i+1, j+1, v))
		}
		sb.WriteString("]\n")
	}

	return sb.String()
}

func (c *Container[T]) Equal(x *Container[T]) bool {
	if c.m != x.m {
		return false
	}
	if c.n != x.n {
		return false
	}
	if !reflect.DeepEqual(c.data, x.data) {
		return false
	}
	return true
}

func (c *Container[T]) Column(j int) (*Container[T], error) {
	m, n := c.Rank()

	if !between(0, j, n) {
		return nil, fmt.Errorf("j is %v, but should be between [0,%v[", j, n)
	}

	r, err := New[T](m, 1)
	if err != nil {
		return nil, err
	}

	for i := 0; i < m; i++ {
		v, err := c.At(i, j)
		if err != nil {
			return nil, err
		}
		_, err = r.SetAt(i, 0, v)
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

func (c *Container[T]) Row(i int) (*Container[T], error) {
	m, n := c.Rank()

	if !between(0, i, m) {
		return nil, fmt.Errorf("i is %v, but should be between [0,%v[", i, m)
	}

	r, err := New[T](1, n)
	if err != nil {
		return nil, err
	}

	for j := 0; j < n; j++ {
		v, err := c.At(i, j)
		if err != nil {
			return nil, err
		}
		_, err = r.SetAt(0, j, v)
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

func (c *Container[T]) Values() ([]T, error) {
	vs := []T{}
	m, n := c.Rank()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v, err := c.At(i, j)
			if err != nil {
				return nil, err
			}
			vs = append(vs, v)
		}
	}

	return vs, nil
}

func Transposition[T computation.Numeric](a *Container[T]) (*Container[T], error) {
	c, err := New[T](a.n, a.m)
	if err != nil {
		return nil, err
	}
	for i := 0; i < a.m; i++ {
		for j := 0; j < a.n; j++ {
			v, err := a.At(i, j)
			if err != nil {
				return nil, err
			}
			_, err = c.SetAt(j, i, v)
			if err != nil {
				return nil, err
			}
		}
	}

	return c, nil
}

func Addition[T computation.Numeric](a, b *Container[T]) (*Container[T], error) {
	if a.m != b.m {
		return nil, fmt.Errorf("m does not match: %d vs %d", a.m, b.m)
	}
	if a.n != b.n {
		return nil, fmt.Errorf("n does not match: %d vs %d", a.n, b.n)
	}
	c, err := New[T](a.m, a.n)
	if err != nil {
		return nil, err
	}

	for i := 0; i < a.m; i++ {
		for j := 0; j < a.n; j++ {
			va, err := a.At(i, j)
			if err != nil {
				return nil, err
			}
			vb, err := b.At(i, j)
			if err != nil {
				return nil, err
			}
			_, err = c.SetAt(i, j, va+vb)
			if err != nil {
				return nil, err
			}
		}
	}

	return c, nil
}

func ScalarMatrixMultiplication[T computation.Numeric](s T, a *Container[T]) (*Container[T], error) {
	c, err := New[T](a.m, a.n)
	if err != nil {
		return nil, err
	}

	for i := 0; i < a.m; i++ {
		for j := 0; j < a.n; j++ {
			va, err := a.At(i, j)
			if err != nil {
				return nil, err
			}
			_, err = c.SetAt(i, j, s*va)
			if err != nil {
				return nil, err
			}
		}
	}

	return c, nil
}

func MatrixMatrixMultiplication[T computation.Numeric](a, b *Container[T]) (*Container[T], error) {
	if a.n != b.m {
		return nil, fmt.Errorf("ranks a.n:%d and  b.m:%d do not match", a.n, b.m)
	}

	c, err := New[T](a.m, b.n)
	if err != nil {
		return nil, err
	}

	for i := 0; i < a.m; i++ {
		for j := 0; j < b.n; j++ {

			var sum T

			for k := 0; k < a.n; k++ {
				va, err := a.At(i, k)
				if err != nil {
					return nil, err
				}
				vb, err := b.At(k, j)
				if err != nil {
					return nil, err
				}
				sum += va * vb
			}

			_, err = c.SetAt(i, j, sum)
			if err != nil {
				return nil, err
			}
		}
	}

	return c, nil
}
