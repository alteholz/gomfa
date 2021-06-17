package gomfa

import (
	"testing"
)

func BenchmarkLd(b *testing.B) {
	var bm, em, dlim float64
	var p [3]float64
	var q [3]float64
	var e [3]float64
	var p1 [3]float64

	bm = 0.00028574
	p[0] = -0.763276255
	p[1] = -0.608633767
	p[2] = -0.216735543
	q[0] = -0.763276255
	q[1] = -0.608633767
	q[2] = -0.216735543
	e[0] = 0.76700421
	e[1] = 0.605629598
	e[2] = 0.211937094
	em = 8.91276983
	dlim = 3e-10

	for i := 0; i < 100000000; i++ {
		Ld(bm, &p, &q, &e, em, dlim, &p1)
	}
}

func TestLd(t *testing.T) {
	var bm, em, dlim float64
	var p [3]float64
	var q [3]float64
	var e [3]float64
	var p1 [3]float64

	bm = 0.00028574
	p[0] = -0.763276255
	p[1] = -0.608633767
	p[2] = -0.216735543
	q[0] = -0.763276255
	q[1] = -0.608633767
	q[2] = -0.216735543
	e[0] = 0.76700421
	e[1] = 0.605629598
	e[2] = 0.211937094
	em = 8.91276983
	dlim = 3e-10

	Ld(bm, &p, &q, &e, em, dlim, &p1)

	if !CheckFloat(p1[0], -0.7632762548968159627, 1e-12) {
		t.Errorf("p1[0] != -0.7632762548968159627\n")
	}
	if !CheckFloat(p1[1], -0.6086337670823762701, 1e-12) {
		t.Errorf("p1[1] != -0.6086337670823762701\n")
	}
	if !CheckFloat(p1[2], -0.2167355431320546947, 1e-12) {
		t.Errorf("p1[2] != -0.2167355431320546947\n")
	}
}
