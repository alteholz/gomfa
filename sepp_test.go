package gomfa

import (
	"testing"
)

func BenchmarkSepp(b *testing.B) {
	var av [3]float64
	var bv [3]float64

	av[0] = 1.0
	av[1] = 0.1
	av[2] = 0.2

	bv[0] = -3.0
	bv[1] = 1e-3
	bv[2] = 0.2

	for i := 0; i < 100000000; i++ {
		_ = Sepp(&av, &bv)
	}
}

func TestSepp(t *testing.T) {
	var av [3]float64
	var bv [3]float64
	var s float64

	av[0] = 1.0
	av[1] = 0.1
	av[2] = 0.2

	bv[0] = -3.0
	bv[1] = 1e-3
	bv[2] = 0.2

	s = Sepp(&av, &bv)

	if !CheckFloat(s, 2.860391919024660768, 1e-12) {
		t.Errorf("s != 2.860391919024660768\n")
	}
}
