package gomfa

import (
	"testing"
)

func BenchmarkSeps(b *testing.B) {
	var al, ap, bl, bp float64

	al = 1.0
	ap = 0.1

	bl = 0.2
	bp = -3.0

	for i := 0; i < 100000000; i++ {
		_ = Seps(al, ap, bl, bp)
	}
}

func TestSeps(t *testing.T) {
	var al, ap, bl, bp, s float64

	al = 1.0
	ap = 0.1

	bl = 0.2
	bp = -3.0

	s = Seps(al, ap, bl, bp)

	if !CheckFloat(s, 2.346722016996998842, 1e-14) {
		t.Errorf("s != 2.346722016996998842\n")
	}
}
