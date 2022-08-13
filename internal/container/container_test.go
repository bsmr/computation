package container

import (
	"reflect"
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

var (
	m3x4vs = []int{11, 12, 13, 14, 21, 22, 23, 24, 31, 32, 33, 34}
)

func matrix3x4() (*Container[int], error) {
	return New(3, 4, m3x4vs...)
}

func TestValues(t *testing.T) {
	m1, err := matrix3x4()
	if err != nil {
		t.Fatalf("matrix3x4() failed with: %s", err)
	}
	t.Logf("m1: %s", m1)

	vs, err := m1.Values()
	if err != nil {
		t.Fatalf("Values() failed with: %s", err)
	}
	t.Logf("vs: %#v", vs)

	if !reflect.DeepEqual(vs, m3x4vs) {
		t.Fatalf("Values() is %#v, expected %#v", vs, m3x4vs)
	}
}

func TestColumn(t *testing.T) {
	m1, err := matrix3x4()
	if err != nil {
		t.Fatalf("matrix3x4() failed with: %s", err)
	}
	t.Logf("m1: %s", m1)

	ce := []int{13, 23, 33}
	cr, err := New(3, 1, ce...)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}

	col := 2
	ci, err := m1.Column(col)
	if err != nil {
		t.Fatalf("Column(%d) failed with: %s", col+1, err)
	}
	t.Logf("Column(%d): %s", col+1, ci)

	if !ci.Equal(cr) {
		t.Fatalf("Column(%d) is %s, expected %s", col+1, ci, cr)
	}

	vr, err := ci.Values()
	if err != nil {
		t.Fatalf("ci.Values() failed with: %s", err)
	}

	if !reflect.DeepEqual(vr, ce) {
		t.Fatalf("Values() is %#v, expected %#v", vr, ce)
	}
}

func TestRow(t *testing.T) {
	m1, err := matrix3x4()
	if err != nil {
		t.Fatalf("matrix3x4() failed with: %s", err)
	}
	t.Logf("m1: %s", m1)

	re := []int{21, 22, 23, 24}
	rr, err := New(1, 4, re...)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}

	row := 1
	ri, err := m1.Row(row)
	if err != nil {
		t.Fatalf("Row(%d) failed with: %s", row+1, err)
	}
	t.Logf("Row(%d): %s", row+1, ri)

	if !ri.Equal(rr) {
		t.Fatalf("Row(%d) is %s, expected %s", row+1, ri, rr)
	}

	xr, err := ri.Values()
	if err != nil {
		t.Fatalf("ri.Values() failed with: %s", err)
	}

	if !reflect.DeepEqual(xr, re) {
		t.Fatalf("Values() is %#v, expected %#v", xr, re)
	}
}
