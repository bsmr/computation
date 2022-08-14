package dxr

import (
	"math"
	"testing"
)

const (
	threshhold = 0.0000000000001
)

func isEqual(a, b float64) bool {
	return inRange(a, b, threshhold)
}

func inRange(a, b, c float64) bool {
	return math.Abs(a-b) < c
}

func TestDxR(t *testing.T) {
	var d, r float64

	for _, c := range []*DxR{
		New(),
		NewDedrees(0.0),
		NewRadians(0.0),
	} {
		for _, v := range []struct {
			d float64
			r float64
		}{
			{d: 0.0, r: 0.0 * math.Pi},
			{d: 30.0, r: (30.0 / 180.0) * math.Pi},
			{d: 45.0, r: 0.25 * math.Pi},
			{d: 60.0, r: (60.0 / 180.0) * math.Pi},
			{d: 90.0, r: 0.5 * math.Pi},
			{d: 120.0, r: (120.0 / 180.0) * math.Pi},
			{d: 180.0, r: 1.0 * math.Pi},
			{d: 270.0, r: 1.5 * math.Pi},
			{d: 360.0, r: 2.0 * math.Pi},
		} {
			//t.Logf("d: %v, r: %v", v.d, v.r)

			c.SetDegress(v.d)
			r = c.Radians()
			if !isEqual(r, v.r) {
				t.Errorf("%v and %v are out of threshold %v", r, v.r, threshhold)
			}
			//t.Logf("v.d: %v, v.r: %v - r: %v - %t", v.d, v.r, r, v.r == r)

			c.SetRadians(v.r)
			d = c.Degress()
			if !isEqual(d, v.d) {
				t.Errorf("%v and %v are out of threshold %v", d, v.d, threshhold)
			}
			//t.Logf("v.d: %v, v.r: %v - d: %v - %t", v.d, v.r, d, v.d == d)
		}
	}
}
