package gomfa

import (
	"testing"
)

func BenchmarkZpv(b *testing.B) {
	var pv [2][3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	for i := 0; i < 10000000; i++ {
		Zpv(&pv)
	}
}

func TestZpv(t *testing.T) {
	var pv [2][3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	Zpv(&pv)

	if pv[0][0] != 0.0 {
		t.Errorf("pv[0][0] != 0.0\n")
	}
	if pv[0][1] != 0.0 {
		t.Errorf("pv[0][1] != 0.0\n")
	}
	if pv[0][2] != 0.0 {
		t.Errorf("pv[0][2] != 0.0\n")
	}
	if pv[1][0] != 0.0 {
		t.Errorf("pv[0][0] != 0.0\n")
	}
	if pv[1][1] != 0.0 {
		t.Errorf("pv[0][1] != 0.0\n")
	}
	if pv[1][2] != 0.0 {
		t.Errorf("pv[0][2] != 0.0\n")
	}
}
