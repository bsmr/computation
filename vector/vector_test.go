package vector

import (
	"testing"

	"github.com/bsmr/computation"
)

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

func bad[T computation.Numeric]() *Vector[T] {
	v := &Vector[T]{
		rank:   1,
		values: []T{1, 2, 3},
	}
	return v
}

func TestErrors(t *testing.T) {
	if _, err := New(0, 1); err == nil {
		t.Fatal("New(0, 1) did not return an error")
	}

	if _, err := Add(bad[int]()); err == nil {
		t.Fatal("Add() did not generate an error")
	}

	for _, c := range [][]*Vector[int]{
		{nil},
		{nil, nil},
		{
			&Vector[int]{
				rank:   1,
				values: []int{1, 2},
			},
			nil,
		},
		{
			nil,
			&Vector[int]{
				rank:   1,
				values: []int{3, 4},
			},
		},
		{
			&Vector[int]{
				rank:   1,
				values: []int{1, 2},
			},
			nil,
			&Vector[int]{
				rank:   1,
				values: []int{3, 4},
			},
		},
	} {
		if _, err := Add(c...); err == nil {
			t.Fatalf("Add(%#v) did not generate an error", c)
		}
	}
}

func TestAt(t *testing.T) {
	v0, _ := New(3, 1, 2, 3)

	for i, c := range v0.values {
		v, err := v0.At(i)
		if err != nil {
			t.Errorf("%#v.At(%d) failed with: %s", v0, i, err)
		}
		if v != c {
			t.Errorf("%#v.At(%d) is %d, expected %d", v0, i, v, c)
		}
	}

	for _, c := range []int{
		-1,
		3,
		4,
	} {
		if _, err := v0.At(c); err == nil {
			t.Fatalf("%#v.At(%d) did not return an errror", v0, c)
		}
	}
}
