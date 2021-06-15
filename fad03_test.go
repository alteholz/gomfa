package gomfa

import (
	"testing"
)

func BenchmarkFad03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Fad03(0.80)
	}
}

func TestFad03(t *testing.T) {
	var fad float64

	fad = Fad03(0.80)

	if !CheckFloat(fad, 1.946709205396925672, 1e-12) {
		t.Errorf("fad !=  1.946709205396925672 -> %v\n", fad)
	}
}
