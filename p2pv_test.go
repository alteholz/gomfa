package gomfa

import (
	"testing"
)

func BenchmarkP2pv(b *testing.B) {
	var pv [2][3]float64
	var p [3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	for i := 0; i < 1000000; i++ {
		P2pv(&p, &pv)
	}
}

func TestP2pv(t *testing.T) {
	var pv [2][3]float64
	var p [3]float64

	p[0] = 0.25
	p[1] = 1.2
	p[2] = 3.0

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	P2pv(&p, &pv)

	if pv[0][0] != 0.25 {
		t.Errorf("pv[0][0] != 0.25\n")
	}
	if pv[0][1] != 1.2 {
		t.Errorf("pv[0][1] != 1.2\n")
	}
	if pv[0][2] != 3.0 {
		t.Errorf("pv[0][2] != 3.0\n")
	}
	if pv[1][0] != 0.0 {
		t.Errorf("pv[1][0] != 0.0\n")
	}
	if pv[1][1] != 0.0 {
		t.Errorf("pv[1][1] != 0.0\n")
	}
	if pv[1][2] != 0.0 {
		t.Errorf("pv[1][2] != 0.0\n")
	}

}
