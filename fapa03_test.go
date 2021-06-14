package gomfa

import (
	"testing"
)

func BenchmarkFapa03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Fapa03(0.80)
	}
}

func TestFapa03(t *testing.T) {
	var fapa float64

	fapa = Fapa03(0.80)

	if !CheckFloat(fapa, 0.1950884762240000000e-1, 1e-12) {
		t.Errorf("fapa !=  0.1950884762240000000e-1 -> %v\n", fapa)
	}
}
