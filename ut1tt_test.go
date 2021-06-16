package gomfa

import (
	"testing"
)

func BenchmarkUt1tt(b *testing.B) {
	var t1, t2 float64

	for i := 0; i < 10000000; i++ {
		_ = Ut1tt(2453750.5, 0.892104561, 64.8499, &t1, &t2)
	}
}

func TestUt1tt(t *testing.T) {
	var t1, t2 float64
	var j int

	j = Ut1tt(2453750.5, 0.892104561, 64.8499, &t1, &t2)

	if j != 0 {
		t.Errorf("j != 0\n")
	}
	if !CheckFloat(t1, 2453750.5, 1e-6) {
		t.Errorf("t1 !=  2453750.5 -> %v\n", t1)
	}
	if !CheckFloat(t2, 0.8928551385462962963, 1e-12) {
		t.Errorf("t2 !=  0.8928551385462962963 -> %v\n", t2)
	}
}
