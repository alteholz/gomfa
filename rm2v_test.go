package gomfa

import (
	"testing"
)

func BenchmarkRm2v(b *testing.B) {
	var r [3][3]float64
	var w [3]float64

	r[0][0] = 0.00
	r[0][1] = -0.80
	r[0][2] = -0.60

	r[1][0] = 0.80
	r[1][1] = -0.36
	r[1][2] = 0.48

	r[2][0] = 0.60
	r[2][1] = 0.48
	r[2][2] = -0.64

	for i := 0; i < 100000000; i++ {
		Rm2v(&r, &w)
	}
}

func TestRm2v(t *testing.T) {
	var r [3][3]float64
	var w [3]float64

	r[0][0] = 0.00
	r[0][1] = -0.80
	r[0][2] = -0.60

	r[1][0] = 0.80
	r[1][1] = -0.36
	r[1][2] = 0.48

	r[2][0] = 0.60
	r[2][1] = 0.48
	r[2][2] = -0.64

	Rm2v(&r, &w)

	if !CheckFloat(w[0], 0.0, 1e-12) {
		t.Errorf("w[0] != 0.0\n")
	}
	if !CheckFloat(w[1], 1.413716694115406957, 1e-12) {
		t.Errorf("w[1] != 1.413716694115406957\n")
	}
	if !CheckFloat(w[2], -1.884955592153875943, 1e-12) {
		t.Errorf("w[2] != -1.884955592153875943\n")
	}
}
