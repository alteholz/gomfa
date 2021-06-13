package gomfa

import (
	"testing"
)

func BenchmarkSxpv(b *testing.B) {
	var pv [2][3]float64
	var spv [2][3]float64
	var s float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.5
	pv[1][1] = 3.2
	pv[1][2] = -0.7

	for i := 0; i < 10000000; i++ {
		Sxpv(s, &pv, &spv)
	}
}

func TestSxpv(t *testing.T) {
	var pv [2][3]float64
	var spv [2][3]float64
	var s float64 = 2.0

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.5
	pv[1][1] = 3.2
	pv[1][2] = -0.7

	Sxpv(s, &pv, &spv)

	if spv[0][0] != 0.6 {
		t.Errorf("spv[0][0] != 0.6\n")
	}
	if spv[0][1] != 2.4 {
		t.Errorf("spv[0][1] != 2.4\n")
	}
	if spv[0][2] != -5.0 {
		t.Errorf("spv[0][2] != -5.0\n")
	}
	if spv[1][0] != 1.0 {
		t.Errorf("spv[1][0] != 1.0\n")
	}
	if spv[1][1] != 6.4 {
		t.Errorf("spv[1][1] != 6.4\n")
	}
	if spv[1][2] != -1.4 {
		t.Errorf("spv[1][2] != -1.4\n")
	}
}
