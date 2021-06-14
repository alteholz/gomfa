package gomfa

import (
	"testing"
)

func BenchmarkPas(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Pas(1.0, 0.1, 0.2, -1.0)
	}
}

func TestPas(t *testing.T) {
	var theta float64

	theta = Pas(1.0, 0.1, 0.2, -1.0)

	if !CheckFloat(theta, -2.724544922932270424, 1e-12) {
		t.Errorf("pv[0][0] != -2.724544922932270424 -> %v\n", theta)
	}
}
