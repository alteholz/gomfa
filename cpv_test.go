package gomfa

import (
	"testing"
)

func BenchmarkCpv(b *testing.B) {
	var pv [2][3]float64
	var c [2][3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	for i := 0; i < 10000000; i++ {
		Cpv(&pv, &c)
	}
}

func TestCpv(t *testing.T) {
	var pv [2][3]float64
	var c [2][3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	Cpv(&pv, &c)

	if c[0][0] != 0.3 {
		t.Errorf("c[0][0] != 0.3\n")
	}
	if c[0][1] != 1.2 {
		t.Errorf("c[0][1] != 1.2\n")
	}
	if c[0][2] != -2.5 {
		t.Errorf("c[0][2] != -2.5\n")
	}
	if c[1][0] != -0.5 {
		t.Errorf("c[1][0] != -0.5\n")
	}
	if c[1][1] != 3.1 {
		t.Errorf("c[1][1] != 3.1\n")
	}
	if c[1][2] != 0.9 {
		t.Errorf("c[1][2] != 0.9\n")
	}
}
