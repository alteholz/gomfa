package gomfa

import (
	"testing"
)

func BenchmarkS2xpv(b *testing.B) {
	var pv [2][3]float64
	var spv [2][3]float64
	var s1, s2 float64 = 2.0, 3.0

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.5
	pv[1][1] = 2.3
	pv[1][2] = -0.4

	for i := 0; i < 10000000; i++ {
		S2xpv(s1, s2, &pv, &spv)
	}
}

func TestS2xpv(t *testing.T) {
	var pv [2][3]float64
	var spv [2][3]float64
	var s1, s2 float64 = 2.0, 3.0

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.5
	pv[1][1] = 2.3
	pv[1][2] = -0.4

	S2xpv(s1, s2, &pv, &spv)

	if spv[0][0] != 0.6 {
		t.Errorf("spv[0][0] != 0.6\n")
	}
	if spv[0][1] != 2.4 {
		t.Errorf("spv[0][1] != 2.4\n")
	}
	if spv[0][2] != -5.0 {
		t.Errorf("spv[0][2] != -5.0\n")
	}
	if !CheckFloat(spv[1][0], 1.5, 1e-12) {
		t.Errorf("spv[1][0] != 1.5 -> %v\n", spv[1][0])
	}
	if !CheckFloat(spv[1][1], 6.9, 1e-12) {
		t.Errorf("spv[1][1] != 6.9 -> %v\n", spv[1][1])
	}
	if !CheckFloat(spv[1][2], -1.2, 1e-12) {
		t.Errorf("spv[1][2] != -1.2 -> %v\n", spv[1][2])
	}
}
