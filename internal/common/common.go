package common

import (
	"fmt"

	"github.com/bsmr/computation"
)

type Container[T computation.Numeric] interface {
	Copy() (Container[T], error)
	Rank() int
	At(int) (T, error)
	SetAt(int, T) (T, error)
	Values() []T
}

func Add[T computation.Numeric](vs ...Container[T]) (Container[T], error) {
	return Operation(func(a, b T) T { return a + b }, vs...)
}

func Mul[T computation.Numeric](vs ...Container[T]) (Container[T], error) {
	return Operation(func(a, b T) T { return a * b }, vs...)
}

func Operation[T computation.Numeric](op func(a, b T) T, containers ...Container[T]) (Container[T], error) {
	var result Container[T]
	var err error

	for index, v := range containers {
		if v == nil {
			return nil, fmt.Errorf("vector %d is nil", index)
		}

		switch index {
		case 0:
			result, err = v.Copy()
			if err != nil {
				return nil, err
			}
		default:
			for i := 0; i < result.Rank(); i++ {
				v1, err := result.At(i)
				if err != nil {
					return nil, err
				}
				v2, err := v.At(i)
				if err != nil {
					return nil, err
				}
				result.SetAt(i, op(v1, v2))
			}
		}
	}

	return result, nil
}
