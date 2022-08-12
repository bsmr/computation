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
