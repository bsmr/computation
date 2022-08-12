package vector

import "testing"

func TestMisc(t *testing.T) {
	t.Logf("%f", float32(2)+float32(3))
	t.Logf("%f", float32(2)*float32(3))
	t.Logf("%f", float64(2)+float64(3))
	t.Logf("%f", float64(2)*float64(3))
	t.Logf("%f", complex64(2)+complex64(3))
	t.Logf("%f", complex64(2)*complex64(3))
	t.Logf("%f", complex128(2)+complex128(3))
	t.Logf("%f", complex128(2)*complex128(3))
}

func TestOps(t *testing.T) {
	v1, err := New(4, 2, 3, 4, 5)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}
	t.Logf("v1 = %#v", v1)

	v2, err := New(4, 6, 7, 8, 9)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}
	t.Logf("v2 = %#v", v2)

	v3, err := Add(v1, v2)
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("v3 = %#v", v3)

	v4, err := Mul(v1, v2)
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("v4 = %#v", v4)
}
