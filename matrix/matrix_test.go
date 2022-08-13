package matrix

import "testing"

func TestBasic(t *testing.T) {
	m, err := New[int](3, 4)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}
	t.Logf("m: %s", m)
}
