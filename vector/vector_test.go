package vector

import (
	"testing"

	"github.com/bsmr/computation"
	"github.com/bsmr/computation/internal/common"
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
	r1, err := common.Add[int](v1, v2, v3)
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("r1 = %#v", r1)

	//r2, err := Mul(v1)
	//r2, err := Mul(v1, v2)
	r2, err := common.Mul[int](v1, v2, v3)
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("r2 = %#v", r2)
}

func TestSpecialCase(t *testing.T) {
	r1, err := common.Add[int]()
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("r1 = %#v", r1)

	r2, err := common.Mul[int]()
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

	if _, err := common.Add[int](bad[int]()); err == nil {
		t.Fatal("Add() did not generate an error")
	}

	for _, c := range [][]common.Container[int]{
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
		if _, err := common.Add(c...); err == nil {
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

func TestContainerCompatibility(t *testing.T) {
	v1, _ := New(3, 1, 2, 3)
	v2, _ := New(3, 4, 5, 6)

	cs := []common.Container[int]{
		v1,
		v2,
	}

	for _, c := range cs {
		t.Logf("Rank(): %d", c.Rank())
	}
}

func equal(as, bs []int) bool {
	if len(as) != len(bs) {
		return false
	}
	for i, v := range as {
		if v != bs[i] {
			return false
		}
	}

	return true
}

func TestValues(t *testing.T) {
	for _, vs := range [][]int{
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
	} {
		v, err := New(len(vs), vs...)
		if err != nil {
			t.Errorf("New(%d,%#v) failed with: %s", len(vs), vs, err)
			continue
		}
		if rs := v.Values(); !equal(rs, vs) {
			t.Errorf("Values() is %#v, expected %#v", rs, vs)
		}
	}
}
