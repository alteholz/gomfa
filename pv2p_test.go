package gomfa

import (
	"testing"
)

func BenchmarkPv2p(b *testing.B) {
	var pv [2][3]float64
	var p [3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	for i := 0; i < 100000000; i++ {
		Pv2p(&pv, &p)
	}
}

func TestPv2p(t *testing.T) {
	var pv [2][3]float64
	var p [3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	Pv2p(&pv, &p)

	if p[0] != 0.3 {
		t.Errorf("p[0] != 0.3\n")
	}
	if p[1] != 1.2 {
		t.Errorf("p[1] != 1.2\n")
	}
	if p[2] != -2.5 {
		t.Errorf("p[2] != -2.5\n")
	}

}
