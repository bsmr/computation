package vector

import "testing"

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

	v3, err := New(4, 10, 11, 12, 13)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}
	t.Logf("v3 = %#v", v3)

	//r1, err := Add(v1)
	//r1, err := Add(v1, v2)
	r1, err := Add(v1, v2, v3)
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("r1 = %#v", r1)

	//r2, err := Mul(v1)
	//r2, err := Mul(v1, v2)
	r2, err := Mul(v1, v2, v3)
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("r2 = %#v", r2)
}

func TestSpecialCase(t *testing.T) {
	r1, err := Add[int]()
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("r1 = %#v", r1)

	r2, err := Mul[int]()
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("r2 = %#v", r2)
}
