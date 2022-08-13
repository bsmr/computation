package vector

import (
	"testing"

	"github.com/bsmr/computation/internal/container"
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
