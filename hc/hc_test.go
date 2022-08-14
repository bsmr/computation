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
