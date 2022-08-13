package orientation

import (
	"testing"

	"github.com/bsmr/computation/internal/container"
)

func TestString(t *testing.T) {
	for _, v := range []struct {
		o Orientation
		s string
	}{
		{o: Unknown, s: textUnknown},
		{o: MxOne, s: textMxOne},
		{o: OnexN, s: textOnexN},
		{o: OnexOne, s: textOnexOne},
		{o: MxN, s: textMxN},
	} {
		if s := v.o.String(); s != v.s {
			t.Errorf("%s.String() is %q, expected %q", v.o, s, v.s)
		}
	}
}

func TestParse(t *testing.T) {
	for _, v := range []struct {
		o Orientation
		s string
	}{
		{o: Unknown, s: textUnknown},
		{o: MxOne, s: textMxOne},
		{o: OnexN, s: textOnexN},
		{o: OnexOne, s: textOnexOne},
		{o: MxN, s: textMxN},
		{o: Unknown, s: "⨯⨯⨯"},
		{o: Unknown, s: "xXx"},
		{o: Unknown, s: "Semmelknödel"},
	} {
		if o := Parse(v.s); o != v.o {
			t.Errorf("Parse(%q) is %s, expected %s", v.o, o, v.s)
		}
	}
}

func TestCheck(t *testing.T) {
	for _, v := range []struct {
		o Orientation
		c *container.Container[int]
	}{
		{
			o: OnexOne,
			c: func() *container.Container[int] { c, _ := container.New[int](1, 1); return c }(),
		},
		{
			o: MxOne,
			c: func() *container.Container[int] { c, _ := container.New[int](4, 1); return c }(),
		},
		{
			o: OnexN,
			c: func() *container.Container[int] { c, _ := container.New[int](1, 4); return c }(),
		},
		{
			o: MxN,
			c: func() *container.Container[int] { c, _ := container.New[int](3, 4); return c }(),
		},
		{
			o: Unknown,
			c: func() *container.Container[int] { c, _ := container.New[int](0, 0); return c }(),
		},
		{
			o: Unknown,
			c: func() *container.Container[int] { c, _ := container.New[int](5, 0); return c }(),
		},
		{
			o: Unknown,
			c: func() *container.Container[int] { c, _ := container.New[int](0, 5); return c }(),
		},
	} {
		if o := Check(v.c); o != v.o {
			t.Errorf("Check(%s) is %s, expected %s", v.c, o, v.o)
		}
	}

}
