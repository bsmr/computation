package common

import (
	"testing"
)

func TestString(t *testing.T) {
	c1, _ := New(3, 2, 1, 2, 3, 4, 5, 6)
	c2, _ := New(3, 2, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0)
	c3, _ := New(3, 2, 1+1i, 2+2i, 3+3i, 4+4i, 5+5i, 6+6i)

	c4, _ := New(3, 1, 1, 2, 3)
	c5, _ := New(1, 3, 1, 2, 3)

	for _, c := range []any{
		c1,
		c2,
		c3,
		c4,
		c5,
	} {
		t.Logf("%s", c)
	}
}

func TestTransposition(t *testing.T) {
	m1, _ := New(3, 1, 1, 2, 3)
	m2, err := Transposition(m1)
	if err != nil {
		t.Fatalf("Transposition() failed with: %s", err)
	}
	t.Logf("m1: %s", m1)
	t.Logf("m2: %s", m2)
}

func TestAddition(t *testing.T) {
	m1, _ := New(2, 3, 1, 2, 3, 4, 5, 6)
	m2, _ := New(2, 3, 7, 8, 9, 10, 11, 12)
	mr, _ := New(2, 3, 8, 10, 12, 14, 16, 18)

	t.Logf("m1: %s", m1)
	t.Logf("m2: %s", m2)

	m3, err := Addition(m1, m2)
	if err != nil {
		t.Fatalf("Addition() failed with: %s", err)
	}
	t.Logf("m3: %s", m3)

	if !m3.Equal(mr) {
		t.Fatal("m3 and mr are not equal")
	}
}

func TestScalarMatrixMultiplication(t *testing.T) {
	m1, _ := New(2, 3, 1, 2, 3, 4, 5, 6)
	s := 2
	mr, _ := New(2, 3, 2, 4, 6, 8, 10, 12)

	t.Logf("m1: %s", m1)
	t.Logf("s : %d", s)

	m2, err := ScalarMatrixMultiplication(s, m1)
	if err != nil {
		t.Fatalf("ScalarMatrixMultiplication() failed with: %s", err)
	}
	t.Logf("m2: %s", m2)

	if !m2.Equal(mr) {
		t.Fatal("m3 and mr are not equal")
	}
}

func TestMatrixMatrixMultiplication(t *testing.T) {
	m1, _ := New(3, 1, 1, 2, 3)
	m2, _ := New(1, 2, 4, 5)
	mr, _ := New(3, 2, 4, 5, 8, 10, 12, 15)

	t.Logf("m1: %s", m1)
	t.Logf("m2: %s", m2)

	m3, err := MatrixMatrixMultiplication(m1, m2)
	if err != nil {
		t.Fatalf("MatrixMatrixMultiplication() failed with: %s", err)
	}
	t.Logf("m3: %s", m3)

	if !m3.Equal(mr) {
		t.Fatal("m3 and mr are not equal")
	}
}
