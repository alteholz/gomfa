package gomfa

import (
	"testing"
)

func BenchmarkTcgtt(b *testing.B) {
	var t1, t2 float64

	for i := 0; i < 10000000; i++ {
		_ = Tcgtt(2453750.5, 0.892862531, &t1, &t2)
	}
}

func TestTcgtt(t *testing.T) {
	var t1, t2 float64
	var j int

	j = Tcgtt(2453750.5, 0.892862531, &t1, &t2)

	if j != 0 {
		t.Errorf("j != 0\n")
	}
	if !CheckFloat(t1, 2453750.5, 1e-6) {
		t.Errorf("t1 !=  2453750.5 -> %v\n", t1)
	}
	if !CheckFloat(t2, 0.8928551387488816828, 1e-12) {
		t.Errorf("t2 !=  0.8928551387488816828 -> %v\n", t2)
	}
}
