package vector

import (
	"reflect"
	"testing"

	"github.com/bsmr/computation/internal/container"
	"github.com/bsmr/computation/internal/orientation"
)

func TestOps(t *testing.T) {
	v1, err := New(2, 3, 4, 5)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}
	t.Logf("v1 = %s", v1)

	v2, err := New(6, 7, 8, 9)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}
	t.Logf("v2 = %s", v2)

	v3, err := New(10, 11, 12, 13)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}
	t.Logf("v3 = %s", v3)

	//r1, err := Add(v1)
	r1, err := Addition(v1, v2)
	//r1, err := common.Add[int](v1, v2, v3)
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("r1 = %s", r1)

	//r2, err := Mul(v1)
	m2, _ := Transposition(v2)
	t.Logf("m2 = %s", m2)
	r2, err := container.MatrixMatrixMultiplication(v1.data, m2.data)
	//r2, err := common.Mul[int](v1, v2, v3)
	if err != nil {
		t.Fatalf("Add() failed with: %s", err)
	}
	t.Logf("r2 = %s", r2)
}

func TestAddition(t *testing.T) {
	v1, err := New(1, 2, 3, 4)
	if err != nil {
		t.Fatalf("v1.New() failed with: %s", err)
	}
	v2, err := New(6, 7, 8, 9)
	if err != nil {
		t.Fatalf("v2.New() failed with: %s", err)
	}
	rs := []int{7, 9, 11, 13}
	vr, err := New(rs...)
	if err != nil {
		t.Fatalf("vr.New() failed with: %s", err)
	}

	t.Logf("v1 = %s", v1)
	t.Logf("v2 = %s", v2)

	v3, err := Addition(v1, v2)
	if err != nil {
		t.Fatalf("Addition() failed with: %s", err)
	}
	t.Logf("v3 = %s", v3)

	if !v3.Equal(vr) {
		t.Fatalf("v3 is %s, expected %s", v3, vr)
	}

	vs, err := v3.Values()
	if err != nil {
		t.Fatalf("Values() failed with: %s", err)
	}
	if !reflect.DeepEqual(vs, rs) {
		t.Fatalf("Values() is %#v, expected %#v", vs, rs)
	}
}

func TestTransposition(t *testing.T) {
	vals := []int{1, 2, 3, 4}
	v1, err := New(vals...)
	if err != nil {
		t.Fatalf("v1.New() failed with: %s", err)
	}
	data, _ := container.New(1, 4, vals...)
	vr := &Vector[int]{
		data:   data,
		layout: orientation.OnexN,
	}

	t.Logf("v1 = %s", v1)

	v3, err := Transposition(v1)
	if err != nil {
		t.Fatalf("Transposition() failed with: %s", err)
	}
	t.Logf("v3 = %s", v3)

	if !v3.EqualReal(vr) {
		t.Fatalf("v3 is %s, expected %s", v3, vr)
	}
	if v3.EqualReal(v1) {
		t.Fatalf("v3 and v3 should not be equal")
	}

	vs, err := v3.Values()
	if err != nil {
		t.Fatalf("Values() failed with: %s", err)
	}
	if !reflect.DeepEqual(vs, vals) {
		t.Fatalf("Values() is %#v, expected %#v", vs, vals)
	}
}

func TestScalarMultiplication(t *testing.T) {
	v1, err := New(1, 2, 3, 4)
	if err != nil {
		t.Fatalf("v1.New() failed with: %s", err)
	}
	s := 3
	vr, err := New(3, 6, 9, 12)
	if err != nil {
		t.Fatalf("vr.New() failed with: %s", err)
	}

	t.Logf("v1 = %s", v1)
	t.Logf("s  = %v", s)

	v3, err := ScalarMultiplication(s, v1)
	if err != nil {
		t.Fatalf("ScalarMultiplication() failed with: %s", err)
	}
	t.Logf("v3 = %s", v3)

	if !v3.Equal(vr) {
		t.Fatalf("v3 is %s, expected %s", v3, vr)
	}
}

func TestEqual(t *testing.T) {
	v1, err := New(1, 2, 3)
	if err != nil {
		t.Fatalf("v1.New() failed with: %s", err)
	}
	v2, err := New(1, 2, 4)
	if err != nil {
		t.Fatalf("v2.New() failed with: %s", err)
	}
	v3, err := New(1, 2, 3, 5)
	if err != nil {
		t.Fatalf("v3.New() failed with: %s", err)
	}

	t.Logf("v1 = %s", v1)
	t.Logf("v2 = %s", v2)
	t.Logf("v3 = %s", v3)

	if v1.Equal(v2) {
		t.Errorf("v1 and v2 are not equal")
	}
	if v1.Equal(v3) {
		t.Errorf("v1 and v3 are not equal")
	}
}
