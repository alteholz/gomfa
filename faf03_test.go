package gomfa

import (
	"testing"
)

func BenchmarkFaf03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Faf03(0.80)
	}
}

func TestFaf03(t *testing.T) {
	var faf float64

	faf = Faf03(0.80)

	if !CheckFloat(faf, 0.2597711366745499518, 1e-12) {
		t.Errorf("faf !=  0.2597711366745499518 -> %v\n", faf)
	}
}
