package gomfa

import (
	"testing"
)

func BenchmarkUt1tai(b *testing.B) {
	var a1, a2 float64

	for i := 0; i < 10000000; i++ {
		_ = Ut1tai(2453750.5, 0.892104561, -32.6659, &a1, &a2)
	}
}

func TestUt1tai(t *testing.T) {
	var a1, a2 float64
	var j int

	j = Ut1tai(2453750.5, 0.892104561, -32.6659, &a1, &a2)

	if j != 0 {
		t.Errorf("j != 0\n")
	}
	if !CheckFloat(a1, 2453750.5, 1e-6) {
		t.Errorf("a1 !=  2453750.5 -> %v\n", a1)
	}
	if !CheckFloat(a2, 0.8924826385462962963, 1e-12) {
		t.Errorf("a2 !=  0.8924826385462962963 -> %v\n", a2)
	}
}
