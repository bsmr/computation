package hc

import (
	"testing"

	"github.com/bsmr/computation/internal/container"
)

func TestHC(t *testing.T) {
	mhcs := []int{
		11, 12, 13, 14,
		21, 22, 23, 24,
		31, 34, 33, 34,
		41, 42, 43, 44,
	}

	mhs, err := container.New(4, 4, mhcs...)
	if err != nil {
		t.Fatalf("container.New() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestNewNull(t *testing.T) {
	mhs, err := NewNull[float64]()
	if err != nil {
		t.Fatalf("NewNull() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestNewIdentity(t *testing.T) {
	mhs, err := NewIdentity[float64]()
	if err != nil {
		t.Fatalf("NewIdentity() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestNewTranslation(t *testing.T) {
	mhs, err := NewTranslation[float64](2, 3, 4)
	if err != nil {
		t.Fatalf("NewTranslation() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestNewScaling(t *testing.T) {
	mhs, err := NewScaling[float64](2, 3, 4)
	if err != nil {
		t.Fatalf("NewScaling() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestNewRotationZ(t *testing.T) {
	theta := D2R(30)
	mhs, err := NewRotationZ(theta)
	if err != nil {
		t.Fatalf("NewRotationZ() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestNewRotationX(t *testing.T) {
	theta := D2R(30)
	mhs, err := NewRotationX(theta)
	if err != nil {
		t.Fatalf("NewRotationX() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestNewRotationY(t *testing.T) {
	theta := D2R(30)
	mhs, err := NewRotationY(theta)
	if err != nil {
		t.Fatalf("NewRotationY() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestConversions(t *testing.T) {
	d0 := 30.0
	r1 := D2R(d0)
	d2 := R2D(r1)

	t.Logf("d0: %v", d0)
	t.Logf("r1: %v", r1)
	t.Logf("d2: %v", d2)
}
