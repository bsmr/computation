package dxr

import "math"

// Degree and Radians converter

type DxR struct {
	degrees float64
	radians float64
}

func New() *DxR {
	return &DxR{}
}

func NewDedrees(degress float64) *DxR {
	dxr := New()
	dxr.SetDegress(degress)
	return dxr
}

func NewRadians(radians float64) *DxR {
	dxr := New()
	dxr.SetRadians(radians)
	return dxr
}

func (dxr *DxR) Degress() float64 {
	return dxr.degrees
}

func (dxr *DxR) Radians() float64 {
	return dxr.radians
}

func (dxr *DxR) SetDegress(degrees float64) {
	dxr.degrees = degrees
	dxr.radians = degrees * math.Pi / 180.0
}

func (dxr *DxR) SetRadians(radians float64) {
	dxr.radians = radians
	dxr.degrees = radians * 180.0 / math.Pi
}
