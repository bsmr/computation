package orientation

import "testing"

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
