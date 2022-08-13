package orientation

import (
	"strings"

	"github.com/bsmr/computation"
	"github.com/bsmr/computation/internal/container"
)

type Orientation int

const (
	Unknown = iota
	MxOne
	OnexN
	OnexOne
	MxN
)

const (
	textUnknown = "?⨯?"
	textMxOne   = "m⨯1"
	textOnexN   = "1⨯n"
	textOnexOne = "1⨯1"
	textMxN     = "m⨯n"
)

func (o Orientation) String() string {
	switch o {
	case MxOne:
		return textMxOne
	case OnexN:
		return textOnexN
	case MxN:
		return textMxN
	case OnexOne:
		return textOnexOne
	default:
		return textUnknown
	}
}

func Parse(s string) Orientation {
	switch strings.ToLower(s) {
	case textMxOne:
		return MxOne
	case textOnexN:
		return OnexN
	case textMxN:
		return MxN
	case textOnexOne:
		return OnexOne
	default:
		return Unknown
	}
}

func Check[T computation.Numeric](c *container.Container[T]) Orientation {
	m, n := c.Rank()
	switch {
	case m == 1 && n == 1:
		return OnexOne
	case m == 1 && n > 1:
		return OnexN
	case m > 1 && n == 1:
		return MxOne
	case m > 1 && n > 1:
		return MxN
	default:
		return Unknown
	}
}
