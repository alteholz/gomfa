package gomfa

import (
	"testing"
)

func BenchmarkPvm(b *testing.B) {
	var pv [2][3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.45
	pv[1][1] = -0.25
	pv[1][2] = 1.1

	for i := 0; i < 1000000; i++ {
		_, _ = Pvm(&pv)
	}
}

func TestPvm(t *testing.T) {
	var pv [2][3]float64
	var r, s float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.45
	pv[1][1] = -0.25
	pv[1][2] = 1.1

	r, s = Pvm(&pv)

	if r != 2.789265136196270604 {
		t.Errorf("r != 2.789265136196270604 -> %v\n", r)
	}
	if s != 1.214495780149111922 {
		t.Errorf("s != 1.214495780149111922 -> %v\n", s)
	}

}

func BenchmarkPvmORG(b *testing.B) {
	var pv [2][3]float64
	var r, s float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.45
	pv[1][1] = -0.25
	pv[1][2] = 1.1

	for i := 0; i < 1000000; i++ {
		PvmORG(&pv, &r, &s)
	}
}

func TestPvmORG(t *testing.T) {
	var pv [2][3]float64
	var r, s float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.45
	pv[1][1] = -0.25
	pv[1][2] = 1.1

	PvmORG(&pv, &r, &s)

	if r != 2.789265136196270604 {
		t.Errorf("r != 2.789265136196270604 -> %v\n", r)
	}
	if s != 1.214495780149111922 {
		t.Errorf("s != 1.214495780149111922 -> %v\n", s)
	}

}
