package hc

import (
	"testing"

	"github.com/bsmr/computation/dxr"
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

var (
	theta30 = dxr.NewDedrees(30)
)

func TestNewRotationZ(t *testing.T) {
	mhs, err := NewRotationZ(theta30)
	if err != nil {
		t.Fatalf("NewRotationZ() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestNewRotationX(t *testing.T) {
	mhs, err := NewRotationX(theta30)
	if err != nil {
		t.Fatalf("NewRotationX() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}

func TestNewRotationY(t *testing.T) {
	mhs, err := NewRotationY(theta30)
	if err != nil {
		t.Fatalf("NewRotationY() failed with: %s", err)
	}

	t.Logf("\nmhs: %s", mhs)
}
