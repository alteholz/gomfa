package gomfa

import (
	"testing"
)

func BenchmarkSxp(b *testing.B) {
	var p [3]float64
	var sp [3]float64
	var s float64 = 2.0

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	for i := 0; i < 100000000; i++ {
		Sxp(s, &p, &sp)
	}
}

func TestSxp(t *testing.T) {
	var p [3]float64
	var sp [3]float64
	var s float64 = 2.0

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	Sxp(s, &p, &sp)

	if sp[0] != 0.6 {
		t.Errorf("sp[0] != 0.6\n")
	}
	if sp[1] != 2.4 {
		t.Errorf("sp[1] != 2.4\n")
	}
	if sp[2] != -5.0 {
		t.Errorf("sp[2] != -5.0\n")
	}
}
